#!/bin/bash

# Set the Ethereum account password
password=""

# Run geth to create a new account
geth account new --password <(echo "$password") --datadir "/Node-1/execution/"

# Check the exit code to see if the account was created successfully
if [ $? -eq 0 ]; then
  echo "Ethereum account created successfully"
else
  echo "Failed to create an Ethereum account"
fi

sleep 2


# Define the directory where the UTC-<timestamp> folders are located
keystore_dir="/Node-1/execution/keystore"

# Find the most recent UTC-<timestamp> folder in the directory
utc_folder=$(ls -t "$keystore_dir" | grep 'UTC-' | head -n 1)

# Check if a UTC-<timestamp> folder was found
if [ -z "$utc_folder" ]; then
  echo "No UTC-<timestamp> folder found in $keystore_dir"
  exit 1
fi

# Find the JSON file inside the most recent UTC-<timestamp> folder
json_file=$(find "$keystore_dir/$utc_folder" -type f -name '*')

# Check if a JSON file was found
if [ -z "$json_file" ]; then
  echo "No JSON file found in $keystore_dir/$utc_folder"
  exit 1
fi

# Extract the 'address' field from the JSON file using jq
address=$(jq -r '.address' "$json_file")

# Define the path for the text file to save the address
output_file="/shared-files/Node-1-address.txt"

# Save the address to the output file
echo "$address" > "$output_file"

# Print the address
echo "Address: $address"
echo "Address saved to $output_file"
