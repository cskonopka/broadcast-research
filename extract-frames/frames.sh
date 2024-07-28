#!/bin/bash

# Check if the correct number of arguments is provided
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 source_directory output_directory"
    exit 1
fi

SOURCE_DIR=$1
OUTPUT_DIR=$2

# Check if the source directory exists
if [ ! -d "$SOURCE_DIR" ]; then
    echo "Source directory does not exist."
    exit 1
fi

# Create the output directory if it does not exist
if [ ! -d "$OUTPUT_DIR" ]; then
    mkdir -p "$OUTPUT_DIR"
fi

# Iterate over each video file in the source directory
for video_file in "$SOURCE_DIR"/*; do
    if [ -f "$video_file" ]; then
        video_name=$(basename "$video_file")
        video_base_name="${video_name%.*}"
        output_folder="$OUTPUT_DIR/$video_base_name"

        # Create a new folder for each video
        mkdir -p "$output_folder"

        # Extract video frames using ffmpeg and add folder name to each exported .png file
        ffmpeg -i "$video_file" "$output_folder/${video_base_name}_frame_%04d.png"
    fi
done

echo "All video frames extracted successfully to $OUTPUT_DIR."
