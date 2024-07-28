#!/bin/bash

# Check if the correct number of arguments is provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 commands_file"
    exit 1
fi

COMMANDS_FILE=$1

# Check if the commands file exists
if [ ! -f "$COMMANDS_FILE" ]; then
    echo "Commands file does not exist."
    exit 1
fi

# Read and execute each command from the file
while IFS= read -r command; do
    if [ -n "$command" ]; then  # Check if the command is not empty
        echo "Executing command: $command"
        eval "$command"
        wait  # Ensure the command finishes before proceeding
        echo "Waiting for 5 seconds..."
        sleep 5
    fi
done < "$COMMANDS_FILE"

echo "All commands executed successfully."
