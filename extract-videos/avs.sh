# #!/bin/bash

# # Check if correct number of arguments is provided
# if [ "$#" -ne 2 ]; then
#     echo "Usage: $0 source_directory destination_directory"
#     exit 1
# fi

# SOURCE_DIR=$1
# DEST_DIR=$2

# # Check if source directory exists
# if [ ! -d "$SOURCE_DIR" ]; then
#     echo "Source directory does not exist."
#     exit 1
# fi

# # Check if destination directory exists, create it if it doesn't
# if [ ! -d "$DEST_DIR" ]; then
#     mkdir -p "$DEST_DIR"
# fi

# # Iterate over each folder in the source directory
# for folder in "$SOURCE_DIR"/*; do
#     if [ -d "$folder" ]; then
#         EDITS_DIR="$folder/edits"
#         if [ -d "$EDITS_DIR" ]; then
#             cp -r "$EDITS_DIR"/* "$DEST_DIR"
#         else
#             echo "No 'edits' folder found in $folder"
#         fi
#     fi
# done

# echo "All files copied successfully to $DEST_DIR"

#!/bin/bash

# Check if the correct number of arguments is provided
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 source_directory destination_directory"
    exit 1
fi

SOURCE_DIR=$1
DEST_DIR=$2

# Check if the source directory exists
if [ ! -d "$SOURCE_DIR" ]; then
    echo "Source directory does not exist."
    exit 1
fi

# Check if the destination directory exists, create it if it doesn't
if [ ! -d "$DEST_DIR" ]; then
    mkdir -p "$DEST_DIR"
fi

# Iterate over each folder in the source directory
for folder in "$SOURCE_DIR"/*; do
    if [ -d "$folder" ]; then
        EDITS_DIR="$folder/edits"
        if [ -d "$EDITS_DIR" ]; then
            folder_name=$(basename "$folder")
            for file in "$EDITS_DIR"/*; do
                if [ -f "$file" ]; then
                    file_name=$(basename "$file")
                    cp "$file" "$DEST_DIR/${folder_name}_$file_name"
                fi
            done
        else
            echo "No 'edits' folder found in $folder"
        fi
    fi
done

echo "All files copied successfully to $DEST_DIR"
