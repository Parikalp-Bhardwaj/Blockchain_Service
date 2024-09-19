#!/bin/bash

ETH2_TESTNET_GENESIS="./shared-files/eth2-testnet-genesis"
CONFIG_PATH="./shared-files/network/config.yaml"
ETH1_CONFIG_PATH="./shared-files/network/genesis.json"
MNEMONICS_PATH="./shared-files/network/mnemonic.yaml"
STATE_OUTPUT_PATH="./shared-files/network/genesis.ssz"
TRANCHES_DIR="./shared-files/network/tranches"

$ETH2_TESTNET_GENESIS merge \
    --config "$CONFIG_PATH" \
    --eth1-config "$ETH1_CONFIG_PATH" \
    --mnemonics "$MNEMONICS_PATH" \
    --state-output "$STATE_OUTPUT_PATH" \
    --tranches-dir "$TRANCHES_DIR"
