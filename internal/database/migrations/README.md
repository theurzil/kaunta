# Kaunta Database Migrations

## Quick Start

### Fresh Installation

```bash
psql $DATABASE_URL -f migrations/000_initial_schema.sql
```

### Migrating from Umami

The migration is **idempotent** - safe to run on existing Umami databases:

```bash
psql $DATABASE_URL -f migrations/000_initial_schema.sql
```

It will:
- Skip existing tables/indexes (IF NOT EXISTS)
- Add Kaunta enhancements (scroll_depth, engagement_time, props)
- Preserve all existing Umami data

## Schema

### Base Tables (Umami-compatible)

- **website** - Tracked websites
- **session** - User sessions with device/location info
- **website_event** - Pageview and custom events

### Kaunta Enhancements

Added to `website_event`:
- `scroll_depth` - Scroll percentage (0-100)
- `engagement_time` - Time on page (milliseconds)
- `props` - Custom event properties (JSONB)

## Manual Migration

```bash
# Remote server
psql "postgresql://user:pass@host:5432/db" -f migrations/000_initial_schema.sql

# Local Docker
docker exec -i postgres_container psql -U umami -d umamidb < migrations/000_initial_schema.sql
```

## Verification

```sql
-- Check tables exist
\dt

-- Check Kaunta columns
\d website_event

-- Should see: scroll_depth, engagement_time, props
```

## Rollback

No rollback needed - migrations only ADD, never DROP.

To remove Kaunta enhancements:

```sql
ALTER TABLE website_event DROP COLUMN IF EXISTS scroll_depth;
ALTER TABLE website_event DROP COLUMN IF EXISTS engagement_time;
ALTER TABLE website_event DROP COLUMN IF EXISTS props;
```
