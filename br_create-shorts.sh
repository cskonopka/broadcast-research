#!/bin/bash

# Check if input directory is provided
if [ -z "$1" ]; then
  echo "Please provide the input directory containing .mp4 files."
  exit 1
fi

# Set the input and output directories
INPUT_DIR="$1"
OUTPUT_DIR="${INPUT_DIR}_shorts"

# Create the output directory if it doesn't exist
mkdir -p "$OUTPUT_DIR"

# Regular expression to match date formats (e.g., 20210825, 2021-08-25)
DATE_REGEX="(_)?(19|20)?[0-9]{2}[-_]?([0-9]{2}[-_]?[0-9]{2})"

# Process each .mp4 file in the input directory
for FILE in "$INPUT_DIR"/*.mp4; do
  # Get the filename without extension
  FILENAME=$(basename "$FILE" .mp4)
  
  # Split the filename at the first underscore
  CLEAN_FILENAME=$(echo "$FILENAME" | sed -E "s/^([^_]*)_//")

  # Remove the date from the filename
  CLEAN_FILENAME=$(echo "$CLEAN_FILENAME" | sed -E "s/$DATE_REGEX//g")

  # Set the output filename
  OUTPUT_FILE="${OUTPUT_DIR}/${CLEAN_FILENAME}.mp4"
#   OUTPUT_FILE="${OUTPUT_DIR}/${CLEAN_FILENAME}_shorts.mp4"
  
  # Convert and stretch video to 9:16 aspect ratio (1080x1920)
  ffmpeg -i "$FILE" -vf "scale=1080:1920,setsar=1:1" "$OUTPUT_FILE"
  
done

echo "Conversion process complete! All videos saved in '$OUTPUT_DIR'."
