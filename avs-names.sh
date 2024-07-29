#!/bin/bash

# Check if input directory is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <input_directory>"
  exit 1
fi

# Input directory
input_dir="$1"

# Get the base name of the input directory
input_dir_name=$(basename "$input_dir")

# Create the output directory two levels up
output_dir="$(dirname "$(dirname "$input_dir")")/AVS-TITLES"
mkdir -p "$output_dir"

# Output file in the output directory
output_file="${output_dir}/${input_dir_name}.txt"

# Clear the output file if it exists
> "$output_file"

# Iterate over files in the input directory
for file in "$input_dir"/*; do
  if [ -f "$file" ]; then
    # Get the file name
    filename=$(basename "$file")
    
    # Remove the file extensions (.mp4, .mov, .mkv)
    filename_no_ext="${filename%.mp4}"
    filename_no_ext="${filename_no_ext%.mov}"
    filename_no_ext="${filename_no_ext%.mkv}"
    
    # Split the filename into date and title
    date_part=$(echo "$filename_no_ext" | cut -d'_' -f1)
    title_part=$(echo "$filename_no_ext" | cut -d'_' -f2-)
    
    # Format as "date,title" and add to the output file
    echo "${date_part},${title_part}" >> "$output_file"
  fi
done

echo "Filenames have been processed and saved to $output_file"
