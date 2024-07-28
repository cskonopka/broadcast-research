#!/bin/bash

# Check if input directory is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <input_directory>"
  exit 1
fi

# Input directory
input_dir="$1"

# Output file in the root of the input directory
output_file="${input_dir}/concatenated_output.txt"

# Clear the output file if it exists
> "$output_file"

# Iterate over .txt files in the input directory
for file in "$input_dir"/*.txt; do
  if [ -f "$file" ]; then
    # Append the content of each file to the output file without adding a newline
    cat "$file" >> "$output_file"
  fi
done

echo "Files have been concatenated into $output_file"
