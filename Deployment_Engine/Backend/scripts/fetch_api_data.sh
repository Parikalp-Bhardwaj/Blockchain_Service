#!/bin/bash

# sudo apt-get install jq

# Fetch data from the API and store it in a variable
API_URL="http://192.168.253.108:3500/eth/v1/node/identity"
JSON_DATA=$(curl -s $API_URL)

# Extract the p2p_addresses field
P2P_ADDRESSES=$(echo $JSON_DATA | jq -r '.data.p2p_addresses[1]')

# Print the extracted p2p_addresses
echo $P2P_ADDRESSES
