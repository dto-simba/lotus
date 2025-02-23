#!/usr/bin/env bash

set -e

# echo "[*] Generate a new wallet for validator and make default address"
# eudico wallet set-default `eudico wallet new`
# echo "[*] Importing wallet with funds in root"

eudico wait-api --timeout=300s
sleep 5
eudico wallet import --as-default --format=json-lotus  /scripts/ipc/src/wallet.key
echo "[*] Initializing validator config and adding validator"
eudico mir validator config init
validator_addr=`eudico mir validator config validator-addr | grep -vE '(/ip6/)' | grep -E '/ip4/127.0.0.1/tcp/1347'`
eudico mir validator config add-validator $validator_addr
echo "[*] Starting validator"
eudico mir validator run --nosync
