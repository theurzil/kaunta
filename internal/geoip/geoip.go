package geoip

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/oschwald/geoip2-golang"
)

var (
	reader *geoip2.Reader
	dbPath string
)

// Init initializes the GeoIP database
// Downloads GeoLite2-City if not present locally (optional - warns if missing)
func Init(dataDir string) error {
	dbPath = filepath.Join(dataDir, "GeoLite2-City.mmdb")

	// Download if missing
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Printf("GeoIP database not found at %s, attempting download...", dbPath)
		if err := downloadDatabase(dbPath); err != nil {
			log.Printf("⚠️  Warning: GeoIP database download failed: %v", err)
			log.Println("⚠️  GeoIP lookups will return 'Unknown'. To enable geolocation:")
			log.Printf("   1. Download from https://geoip.maxmind.com/")
			log.Printf("   2. Place at: %s", dbPath)
			// Don't fail - continue without GeoIP
			return nil
		}
		log.Println("✓ GeoIP database downloaded successfully")
	}

	// Open database
	var err error
	reader, err = geoip2.Open(dbPath)
	if err != nil {
		log.Printf("⚠️  Warning: Could not load GeoIP database: %v", err)
		log.Println("⚠️  GeoIP lookups will return 'Unknown'")
		// Don't fail - continue without GeoIP
		return nil
	}

	log.Println("✓ GeoIP database loaded")
	return nil
}

// LookupIP returns country, city, and region for an IP address
func LookupIP(ipStr string) (country, city, region string) {
	if reader == nil {
		return "Unknown", "", ""
	}

	ip := net.ParseIP(ipStr)
	if ip == nil {
		return "Unknown", "", ""
	}

	record, err := reader.City(ip)
	if err != nil {
		log.Printf("GeoIP lookup error for %s: %v", ipStr, err)
		return "Unknown", "", ""
	}

	country = record.Country.IsoCode
	if country == "" {
		country = "Unknown"
	}

	city = record.City.Names["en"]
	region = record.Subdivisions[0].Names["en"]

	return country, city, region
}

// Close closes the GeoIP database
func Close() error {
	if reader != nil {
		return reader.Close()
	}
	return nil
}

// downloadDatabase downloads GeoLite2-City database from MaxMind
// Source: https://github.com/maxmind/geoipupdate
func downloadDatabase(dbPath string) error {
	// Create directory if needed
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Use GitSquared/node-geolite2-redist as source
	// This is what Umami uses as default
	url := "https://raw.githubusercontent.com/GitSquared/node-geolite2-redist/master/redist/GeoLite2-City.mmdb"

	log.Printf("Downloading GeoIP database from %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status %d", resp.StatusCode)
	}

	// Write to file
	out, err := os.Create(dbPath)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("failed to write database: %w", err)
	}

	return nil
}
