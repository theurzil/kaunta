package geoip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookupIP(t *testing.T) {
	// Note: Without database loaded, LookupIP returns empty string for country (not "Unknown")
	// This is intentional to prevent CHAR(2) column constraint violations
	// This tests the behavior when database is not initialized
	tests := []struct {
		name      string
		ip        string
		wantError bool
	}{
		{
			name:      "Valid IP without DB",
			ip:        "8.8.8.8",
			wantError: false,
		},
		{
			name:      "Invalid IP format",
			ip:        "999.999.999.999",
			wantError: false,
		},
		{
			name:      "Empty string",
			ip:        "",
			wantError: false,
		},
		{
			name:      "Localhost",
			ip:        "127.0.0.1",
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			country, city, region := LookupIP(tt.ip)

			// When DB is not loaded, should return empty string for country
			// to comply with CHAR(2) database constraint
			if reader == nil {
				assert.Equal(t, "", country)
				assert.Equal(t, "", city)
				assert.Equal(t, "", region)
			}
		})
	}
}

func TestCloseWithoutInit(t *testing.T) {
	// Should not panic if Close called without Init
	err := Close()
	assert.NoError(t, err)
}

// Integration test: Tests actual database if available
func TestLookupIPWithDatabase(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	// This would require a valid GeoIP database file
	// Included as documentation of expected behavior
	expectedResults := map[string]struct {
		country string
		city    string
	}{
		"8.8.8.8":      {"US", "Mountain View"}, // Google DNS
		"1.1.1.1":      {"US", "Los Angeles"},   // Cloudflare DNS
		"9.9.9.9":      {"US", ""},              // Quad9 DNS
		"208.67.222.2": {"US", ""},              // OpenDNS
	}

	// Only run if database is loaded
	if reader == nil {
		t.Skip("GeoIP database not initialized")
	}

	for ip, expected := range expectedResults {
		t.Run(ip, func(t *testing.T) {
			country, city, _ := LookupIP(ip)
			assert.Equal(t, expected.country, country, "Country mismatch for %s", ip)
			// Note: City may vary based on MaxMind DB version
			if expected.city != "" {
				assert.NotEmpty(t, city, "Expected city for %s", ip)
			}
		})
	}
}
