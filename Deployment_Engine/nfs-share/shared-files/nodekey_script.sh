#!/bin/bash

sleep 10
# Source file
source_file="/Node-1/execution/geth/nodekey"
output_file="/shared-files/nodekey.txt"  # Updated to create the text file
# Destination file
destination_file="/shared-files/nodekey"

# Check if the source file exists
if [ -f "$source_file" ]; then
    # Copy the contents of the source file to the destination file
    cp "$source_file" "$destination_file"
    echo "File copied successfully."

    # Set permissions for the destination file
    chmod 777 "$destination_file"
    echo "Permissions set successfully."

    # Create a text file with the path to the nodekey
    cat "$source_file" > "$output_file"
    echo "Text file created successfully."
else
    echo "Source file does not exist."
fi

