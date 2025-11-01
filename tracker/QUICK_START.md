# Kaunta Tracker - Quick Start

## 5-Minute Setup

### 1. Add Script to Your Site

```html
<!DOCTYPE html>
<html>
<head>
  <script
    defer
    data-website-id="YOUR-UUID-HERE"
    src="https://census.yourdomain.com/k.js">
  </script>
</head>
<body>
  <!-- Your content -->
</body>
</html>
```

Replace:
- `YOUR-UUID-HERE` with your website UUID from Kaunta
- `https://census.yourdomain.com` with your Kaunta instance URL

### 2. Track Custom Events

```javascript
kaunta.track('button_click');
kaunta.track('signup', { plan: 'pro', source: 'homepage' });
```

### 3. Done

Kaunta automatically tracks:
- Pageviews (including SPA navigation)
- Scroll depth
- Engagement time
- Outbound link clicks
- Custom events

---

## Common Use Cases

### E-commerce

```javascript
kaunta.track('product_view', { product_id: 'SKU-001', price: 29.99 });
kaunta.track('add_to_cart', { product_id: 'SKU-001', quantity: 2 });
kaunta.track('purchase', { order_id: 'ORD-123', total: 59.98 });
```

### SaaS Product

```javascript
kaunta.track('feature_used', { feature: 'export_pdf', plan: 'pro' });
kaunta.track('signup', { plan: 'free' });
kaunta.track('upgrade', { from: 'free', to: 'pro' });
```

### Content Site

```javascript
kaunta.track('article_view', { article_id: 'post-123', category: 'tutorial' });
kaunta.track('newsletter_signup', { list: 'weekly' });
kaunta.track('social_share', { platform: 'twitter' });
```

---

## Configuration

| Attribute | Default | Description |
|-----------|---------|-------------|
| `data-website-id` | required | Your website UUID |
| `data-api-url` | script location | API endpoint |
| `data-auto-track` | true | Auto-track pageviews |
| `data-track-outbound` | true | Auto-track outbound links |
| `data-respect-dnt` | true | Respect Do Not Track |
| `data-exclude-hash` | false | Remove URL hash |
| `data-domains` | all | Comma-separated whitelist |

**Example - Disable outbound tracking:**

```html
<script
  data-website-id="uuid"
  data-track-outbound="false"
  src="https://census.yourdomain.com/k.js">
</script>
```

**Example - Multi-domain:**

```html
<script
  data-website-id="uuid"
  data-domains="example.com, app.example.com"
  src="https://census.yourdomain.com/k.js">
</script>
```

---

## Testing

### 1. Check if Loaded

Open browser console (F12):
```javascript
console.log(window.kaunta);
// Should output: {track: ƒ, trackPageview: ƒ}
```

### 2. Track Test Event

```javascript
kaunta.track('test_event', { source: 'console' });
```

### 3. Monitor Network

1. Open DevTools (F12)
2. Network tab
3. Filter by "send"
4. Trigger events
5. Inspect POST requests

Expected: POST `/api/send` with status 200 or 202

---

## Troubleshooting

### Script Not Tracking?

Check if loaded:
```javascript
console.log(window.kaunta);
```

Check if website ID is set:
```html
<script data-website-id="..." ...>
```

Check if DNT is enabled:
```javascript
console.log(navigator.doNotTrack);
```

If DNT is '1' or 'yes', set `data-respect-dnt="false"` for testing.

### Events Not Appearing?

1. Open DevTools → Network
2. Look for POST to `/api/send`
3. Check response status
4. Try incognito mode (ad blockers may block requests)

---

## Data Format

Every event includes:
- website UUID
- hostname
- URL
- page title
- referrer
- screen resolution
- language
- scroll depth
- engagement time

Custom events also include:
- event name
- custom properties

---

## Privacy

- No cookies
- No localStorage
- Respects Do Not Track
- No personal data collection
- GDPR compliant

---

## Size

- Source: 11 KB
- Minified: 3.2 KB
- Gzipped: 1.6 KB

---

See `README.md` for full documentation.
