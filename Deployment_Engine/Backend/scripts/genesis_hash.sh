#!/bin/bash

hash_value=$(curl -sX POST --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x0", true],"id":1}' -H "Content-Type: application/json" http://192.168.253.108:8545 | jq -r '.result.hash')

# Output the hash value
echo "Hash Value: $hash_value"

# Save the hash value to a file
echo "$hash_value" > /mnt/shared-files/network/deposit_contract_block_hash.txt

sleep 2

