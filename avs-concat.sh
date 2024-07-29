#!/bin/bash

# Check if input directory is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <input_directory>"
  exit 1
fi

# Input directory
input_dir="$1"

# Temporary file for concatenated content
temp_file=$(mktemp)

# Output file in the input directory
output_file="${input_dir}/concatenated.txt"

# Concatenate all files within the input directory
cat "$input_dir"/* > "$temp_file"

# Add header "name, date" at the top of the concatenated file
echo "date, name" | cat - "$temp_file" > "$output_file"

# Add a new line at the top of the concatenated file
echo "" | cat - "$output_file" > "${output_file}.tmp"
mv "${output_file}.tmp" "$output_file"

# Clean up temporary file
rm "$temp_file"

echo "Files have been concatenated and saved to $output_file"
