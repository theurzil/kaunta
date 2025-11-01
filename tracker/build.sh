#!/bin/bash
set -e

echo "🔨 Building Kaunta tracker..."
echo ""

# Check for terser
if ! command -v terser &> /dev/null; then
    echo "❌ terser not found. Installing..."
    npm install -g terser
fi

# Minify
echo "📦 Minifying..."
terser kaunta.js \
  --compress passes=3,drop_console=false \
  --mangle \
  --comments false \
  --output kaunta.min.js

echo "✅ Minified"

# Generate SRI hash
echo ""
echo "🔐 SRI Hash (sha384):"
echo "   sha384-$(openssl dgst -sha384 -binary kaunta.min.js | openssl base64 -A)"

# Gzip for size testing
gzip -c kaunta.min.js > kaunta.min.js.gz

# Report sizes
echo ""
echo "📊 Sizes:"
printf "   Original: %'8d bytes (%.2f KB)\n" $(wc -c < kaunta.js) $(echo "scale=2; $(wc -c < kaunta.js)/1024" | bc)
printf "   Minified: %'8d bytes (%.2f KB)\n" $(wc -c < kaunta.min.js) $(echo "scale=2; $(wc -c < kaunta.min.js)/1024" | bc)
printf "   Gzipped:  %'8d bytes (%.2f KB)\n" $(wc -c < kaunta.min.js.gz) $(echo "scale=2; $(wc -c < kaunta.min.js.gz)/1024" | bc)

# Calculate savings
original_size=$(wc -c < kaunta.js)
minified_size=$(wc -c < kaunta.min.js)
gzipped_size=$(wc -c < kaunta.min.js.gz)

minified_percent=$(echo "scale=1; 100 - ($minified_size * 100 / $original_size)" | bc)
gzipped_percent=$(echo "scale=1; 100 - ($gzipped_size * 100 / $original_size)" | bc)

echo ""
echo "💾 Savings:"
echo "   Minified: -${minified_percent}%"
echo "   Gzipped:  -${gzipped_percent}%"

# Cleanup
rm kaunta.min.js.gz

echo ""
echo "✨ Done! Output: kaunta.min.js"
