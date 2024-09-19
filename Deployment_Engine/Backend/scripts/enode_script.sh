#!/bin/bash

# # Specify the directory path
# directory="/mnt/shared-files"
# # Check if the directory exists
# sleep 2
# if [ -d "$directory" ]; then
#     # Change to the specified directory
#     cd "$directory"
    
#     # Run the command (e.g., bootnode)
#     bootnode -nodekey ./nodekey.txt -writeaddress
# else
#     echo "Directory not found: $directory"
# fi

response=$(curl -s -X POST -H "Content-Type: application/json" http://192.168.253.108:8545 --data '{"jsonrpc": "2.0", "id": 42, "method": "admin_nodeInfo", "params": []}')

# Extract the enode value using jq
enode_value=$(echo "$response" | jq -r '.result.enode')

# Print the enode value
echo "$enode_value"
