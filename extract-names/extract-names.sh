#!/bin/bash

# Check if input directory is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <input_directory>"
  exit 1
fi

# Input directory
input_dir="$1"

# Output file in the root of the input directory
output_file="${input_dir}/file_names.txt"

# Clear the output file if it exists
> "$output_file"

# Iterate over directories in the input directory
for dir in "$input_dir"/*; do
  if [ -d "$dir/edits" ]; then
    folder_name=$(basename "$dir")
    # Iterate over files in the "edits" folder
    for file in "$dir/edits"/*; do
      if [ -f "$file" ]; then
        # Get the file name
        filename=$(basename "$file")
        
        # Add the file name and folder name to the output file
        echo "${filename},${folder_name}" >> "$output_file"
      fi
    done
  fi
done

echo "File names have been written to $output_file"
