#!/bin/bash
API_URL="http://192.168.253.108:5052/eth/v1/node/identity"
JSON_DATA=$(curl -s $API_URL)
# Extracted values
peer_id=$(echo $JSON_DATA | jq -r '.data.peer_id')

# Print the results
echo "$peer_id"