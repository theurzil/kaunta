package main

import (
	_ "embed"
	"log"
	"strings"

	"github.com/seuros/kaunta/internal/cli"
)

//go:embed VERSION
var versionFile string

//go:embed assets/kaunta.min.js
var trackerScript []byte

//go:embed assets/vendor/alpine.min.js
var alpineJS []byte

//go:embed assets/vendor/chart.min.js
var chartJS []byte

//go:embed assets/vendor/leaflet-1.9.4.min.js
var leafletJS []byte

//go:embed assets/vendor/leaflet-1.9.4.min.css
var leafletCSS []byte

//go:embed assets/vendor/topojson-client-3.1.0.min.js
var topojsonJS []byte

//go:embed assets/data/countries-110m.json
var countriesGeoJSON []byte

//go:embed dashboard.html
var dashboardTemplate []byte

func main() {
	// Extract version from embedded file
	version := strings.TrimSpace(versionFile)

	// Execute CLI with embedded assets
	if err := cli.Execute(
		version,
		trackerScript,
		alpineJS,
		chartJS,
		leafletJS,
		leafletCSS,
		topojsonJS,
		countriesGeoJSON,
		dashboardTemplate,
	); err != nil {
		log.Fatal(err)
	}
}
