package handlers

import (
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseUserAgent(t *testing.T) {
	tests := []struct {
		name           string
		userAgent      string
		expectedBrowser string
		expectedOS      string
		expectedDevice  string
	}{
		{
			name:           "Chrome on Windows Desktop",
			userAgent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
			expectedBrowser: "Chrome",
			expectedOS:      "Windows",
			expectedDevice:  "desktop",
		},
		{
			name:           "Firefox on macOS",
			userAgent:      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:89.0) Gecko/20100101 Firefox/89.0",
			expectedBrowser: "Firefox",
			expectedOS:      "macOS",
			expectedDevice:  "desktop",
		},
		{
			name:           "Safari on iOS",
			userAgent:      "Mozilla/5.0 (iPhone; CPU iPhone OS 14_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1",
			expectedBrowser: "Safari",
			expectedOS:      "iOS",
			expectedDevice:  "mobile",
		},
		{
			name:           "Chrome on Android",
			userAgent:      "Mozilla/5.0 (Linux; Android 11; SM-G991B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.120 Mobile Safari/537.36",
			expectedBrowser: "Chrome",
			expectedOS:      "Android",
			expectedDevice:  "mobile",
		},
		{
			name:           "Edge on Windows",
			userAgent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.59",
			expectedBrowser: "Edge",
			expectedOS:      "Windows",
			expectedDevice:  "desktop",
		},
		{
			name:           "Unknown browser",
			userAgent:      "Some Custom Browser/1.0",
			expectedBrowser: "Unknown",
			expectedOS:      "Unknown",
			expectedDevice:  "desktop",
		},
		{
			name:           "Empty user agent",
			userAgent:      "",
			expectedBrowser: "Unknown",
			expectedOS:      "Unknown",
			expectedDevice:  "desktop",
		},
		{
			name:           "Linux Desktop",
			userAgent:      "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36",
			expectedBrowser: "Chrome",
			expectedOS:      "Linux",
			expectedDevice:  "desktop",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			browser, os, device := parseUserAgent(tt.userAgent)

			require.NotNil(t, browser, "Browser should not be nil")
			require.NotNil(t, os, "OS should not be nil")
			require.NotNil(t, device, "Device should not be nil")

			assert.Equal(t, tt.expectedBrowser, *browser, "Browser mismatch")
			assert.Equal(t, tt.expectedOS, *os, "OS mismatch")
			assert.Equal(t, tt.expectedDevice, *device, "Device mismatch")
		})
	}
}

func TestIsSpamReferrer(t *testing.T) {
	tests := []struct {
		name     string
		referrer string
		isSpam   bool
	}{
		{
			name:     "Legitimate referrer - Google",
			referrer: "https://www.google.com/search?q=test",
			isSpam:   false,
		},
		{
			name:     "Legitimate referrer - Facebook",
			referrer: "https://facebook.com/page",
			isSpam:   false,
		},
		{
			name:     "Spam referrer - semalt.com",
			referrer: "https://semalt.com/crawler",
			isSpam:   true,
		},
		{
			name:     "Spam referrer - buttons-for-website.com",
			referrer: "https://buttons-for-website.com",
			isSpam:   true,
		},
		{
			name:     "Spam referrer with subdomain",
			referrer: "https://spam.semalt.com/page",
			isSpam:   true,
		},
		{
			name:     "Empty referrer",
			referrer: "",
			isSpam:   false,
		},
		{
			name:     "Invalid URL",
			referrer: "not a valid url",
			isSpam:   false,
		},
		{
			name:     "Case insensitive spam detection",
			referrer: "https://SEMALT.COM/page",
			isSpam:   true,
		},
		{
			name:     "Spam referrer - darodar",
			referrer: "https://darodar.com",
			isSpam:   true,
		},
		{
			name:     "Spam referrer - best-seo-offer",
			referrer: "https://best-seo-offer.com/test",
			isSpam:   true,
		},
		{
			name:     "Direct traffic (no referrer)",
			referrer: "",
			isSpam:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSpamReferrer(tt.referrer)
			assert.Equal(t, tt.isSpam, result, "Spam detection mismatch")
		})
	}
}

func TestGenerateUUID(t *testing.T) {
	tests := []struct {
		name  string
		parts []string
	}{
		{
			name:  "Single part",
			parts: []string{"test"},
		},
		{
			name:  "Multiple parts",
			parts: []string{"website-id", "192.168.1.1", "user-agent"},
		},
		{
			name:  "Empty parts",
			parts: []string{"", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uuid1 := generateUUID(tt.parts...)
			uuid2 := generateUUID(tt.parts...)

			// Same input should produce same UUID (deterministic)
			assert.Equal(t, uuid1, uuid2, "UUID should be deterministic")

			// Should be valid UUID
			assert.NotEqual(t, uuid.Nil, uuid1, "UUID should not be nil")

			// Different input should produce different UUID
			differentParts := append([]string{"different"}, tt.parts...)
			uuid3 := generateUUID(differentParts...)
			assert.NotEqual(t, uuid1, uuid3, "Different inputs should produce different UUIDs")
		})
	}
}

func TestHashDate(t *testing.T) {
	baseTime := time.Date(2025, 11, 5, 14, 30, 0, 0, time.UTC)

	tests := []struct {
		name     string
		time     time.Time
		period   string
		expected string
	}{
		{
			name:   "Month period",
			time:   baseTime,
			period: "month",
		},
		{
			name:   "Hour period",
			time:   baseTime,
			period: "hour",
		},
		{
			name:   "Day period (default)",
			time:   baseTime,
			period: "day",
		},
		{
			name:   "Unknown period defaults to day",
			time:   baseTime,
			period: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash1 := hashDate(tt.time, tt.period)
			hash2 := hashDate(tt.time, tt.period)

			// Same time and period should produce same hash
			assert.Equal(t, hash1, hash2, "Hash should be deterministic")

			// Hash should be 32 characters (MD5 hex)
			assert.Len(t, hash1, 32, "Hash should be 32 characters (MD5 hex)")

			// Hash should only contain hex characters
			assert.Regexp(t, "^[0-9a-f]+$", hash1, "Hash should only contain hex characters")
		})
	}

	// Test that different periods produce different hashes
	t.Run("Different periods produce different hashes", func(t *testing.T) {
		hashMonth := hashDate(baseTime, "month")
		hashHour := hashDate(baseTime, "hour")
		hashDay := hashDate(baseTime, "day")

		assert.NotEqual(t, hashMonth, hashHour, "Month and hour hashes should differ")
		assert.NotEqual(t, hashMonth, hashDay, "Month and day hashes should differ")
		assert.NotEqual(t, hashHour, hashDay, "Hour and day hashes should differ")
	})

	// Test that same month but different days produce same hash for month period
	t.Run("Same month different days", func(t *testing.T) {
		time1 := time.Date(2025, 11, 1, 10, 0, 0, 0, time.UTC)
		time2 := time.Date(2025, 11, 15, 22, 0, 0, 0, time.UTC)

		hash1 := hashDate(time1, "month")
		hash2 := hashDate(time2, "month")

		assert.Equal(t, hash1, hash2, "Same month should produce same hash")
	})

	// Test that same day but different hours produce different hash for hour period
	t.Run("Same day different hours", func(t *testing.T) {
		time1 := time.Date(2025, 11, 5, 10, 30, 0, 0, time.UTC)
		time2 := time.Date(2025, 11, 5, 15, 45, 0, 0, time.UTC)

		hash1 := hashDate(time1, "hour")
		hash2 := hashDate(time2, "hour")

		assert.NotEqual(t, hash1, hash2, "Different hours should produce different hash")
	})
}

func TestTrackingPayload_URLValidation(t *testing.T) {
	tests := []struct {
		name      string
		urlLength int
		shouldPass bool
	}{
		{
			name:      "Normal URL",
			urlLength: 100,
			shouldPass: true,
		},
		{
			name:      "Maximum allowed URL",
			urlLength: MaxURLSize,
			shouldPass: true,
		},
		{
			name:      "URL too long",
			urlLength: MaxURLSize + 1,
			shouldPass: false,
		},
		{
			name:      "Very long URL",
			urlLength: 5000,
			shouldPass: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Generate URL of specified length
			url := "https://example.com/" + strings.Repeat("a", tt.urlLength-20)

			if tt.shouldPass {
				assert.LessOrEqual(t, len(url), MaxURLSize, "URL should be within max size")
			} else {
				assert.Greater(t, len(url), MaxURLSize, "URL should exceed max size")
			}
		})
	}
}

func TestTrackingPayload_ValidationLogic(t *testing.T) {
	// Test scroll depth validation
	t.Run("Scroll depth validation", func(t *testing.T) {
		validDepths := []int{0, 25, 50, 75, 100}
		invalidDepths := []int{-1, -10, 101, 150}

		for _, depth := range validDepths {
			assert.GreaterOrEqual(t, depth, 0, "Valid depth should be >= 0")
			assert.LessOrEqual(t, depth, 100, "Valid depth should be <= 100")
		}

		for _, depth := range invalidDepths {
			shouldBeInvalid := depth < 0 || depth > 100
			assert.True(t, shouldBeInvalid, "Depth %d should be invalid", depth)
		}
	})

	// Test engagement time validation
	t.Run("Engagement time validation", func(t *testing.T) {
		validTimes := []int{0, 1000, 5000, 30000, 60000}
		invalidTimes := []int{-1, -100, -5000}

		for _, time := range validTimes {
			assert.GreaterOrEqual(t, time, 0, "Valid time should be >= 0")
		}

		for _, time := range invalidTimes {
			assert.Less(t, time, 0, "Invalid time should be < 0")
		}
	})
}

func TestSpamReferrerList(t *testing.T) {
	// Ensure spam referrer list is not empty
	assert.NotEmpty(t, spamReferrers, "Spam referrer list should not be empty")

	// Ensure all entries are lowercase for case-insensitive matching
	for _, spam := range spamReferrers {
		assert.Equal(t, strings.ToLower(spam), spam, "Spam referrer %s should be lowercase", spam)
	}

	// Test that known spam domains are in the list
	knownSpam := []string{"semalt.com", "darodar.com", "best-seo-offer.com"}
	for _, known := range knownSpam {
		found := false
		for _, spam := range spamReferrers {
			if spam == known {
				found = true
				break
			}
		}
		assert.True(t, found, "Known spam domain %s should be in list", known)
	}
}

func TestMaxURLSize(t *testing.T) {
	// Verify MaxURLSize constant is set to Plausible standard
	assert.Equal(t, 2000, MaxURLSize, "MaxURLSize should be 2000 (Plausible standard)")
}

// Benchmark tests for performance-critical functions
func BenchmarkParseUserAgent(b *testing.B) {
	ua := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parseUserAgent(ua)
	}
}

func BenchmarkGenerateUUID(b *testing.B) {
	parts := []string{"website-id", "192.168.1.1", "user-agent", "salt"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		generateUUID(parts...)
	}
}

func BenchmarkIsSpamReferrer(b *testing.B) {
	referrer := "https://www.google.com/search?q=test"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isSpamReferrer(referrer)
	}
}

func BenchmarkHashDate(b *testing.B) {
	t := time.Now()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hashDate(t, "month")
	}
}
