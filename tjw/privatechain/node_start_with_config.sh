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

# --nodiscover \
# --bootnodes "enode://7c67468cce9ed7d29b790f506c04ac010b2c1e9d0dd17384056183a9657747b4e6803182dce5fa75dcd4632a8ab7e3a0a249f963ceae6c2008de46083828e50d@127.0.0.1:0?discport=30301" \