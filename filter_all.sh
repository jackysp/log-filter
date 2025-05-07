#!/usr/bin/env bash
set -euo pipefail

if [[ $# -lt 2 ]]; then
  echo "Usage: $0 START_TIME END_TIME"
  echo "Example: $0 \"2025/04/25 10:27:00\" \"2025/04/25 10:43:00\""
  exit 1
fi

START="$1"
END="$2"
INPUT_DIR="/Users/xxx/Downloads/xxx"
OUTPUT_DIR="$INPUT_DIR/filtered"

mkdir -p "$OUTPUT_DIR"

# Recursively find and filter all tidb.log, tikv.log, and pd.log
find "$INPUT_DIR" -type f \( -name "tidb.log" -o -name "tikv.log" -o -name "pd.log" \) | while read -r IN; do
  DIR=$(dirname "$IN")
  BASE=$(basename "$IN" .log)
  OUT="$DIR/${BASE}_filtered.log"
  echo "Filtering $IN → $OUT"
  ./log-filter \
    -input="$IN" \
    -start="$START" \
    -end="$END" \
    -output="$OUT"
done

echo "All done. Filtered logs written alongside originals."

