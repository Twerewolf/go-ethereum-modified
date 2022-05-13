#!/bin/bash

GETHPATH=../../build/bin

$GETHPATH/geth \
--port 30303 \
--miner.threads 1 \
--datadir . \
--networkid 224 \
--ipcpath ./blockchain/geth.ipc \
--http  \
--http.port 8545 \
--miner.gasprice 1 \
--miner.etherbase "23ffb1da16604cfc0373feedbfa5fe60b928d5c7" \
--allow-insecure-unlock \
console
