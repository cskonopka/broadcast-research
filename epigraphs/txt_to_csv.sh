#!/bin/bash

# Check if a directory is provided as an argument
if [ $# -ne 1 ]; then
    echo "Usage: $0 <directory_path>"
    exit 1
fi

# Store the directory path
DIR="$1"

# Check if the directory exists
if [ ! -d "$DIR" ]; then
    echo "Error: Directory '$DIR' does not exist."
    exit 1
fi

# Loop through all .txt files in the specified directory
for file in "$DIR"/*.txt; do
    # Check if any .txt files exist
    if [ -f "$file" ]; then
        # Get the base filename without the .txt extension
        base_name=$(basename "$file" .txt)
        # Define the new filename with .csv extension
        new_file="$DIR/$base_name.csv"
        # Rename the file
        mv "$file" "$new_file"
        echo "Converted: $file -> $new_file"
    fi
done

# Check if any files were converted
if [ ! -f "$DIR"/*.txt ]; then
    echo "No .txt files found in the directory."
fi