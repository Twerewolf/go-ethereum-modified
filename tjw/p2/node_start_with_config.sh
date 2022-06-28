#!/bin/bash

GETHPATH=../../build/bin

$GETHPATH/geth \
--port 30303 \
--miner.threads 1 \
--datadir . \
--networkid 224 \
--ipcpath ./blockchain/geth.ipc \
--http  \
--http.port 8546 \
--miner.gasprice 1 \
--miner.etherbase "e7fb044503784dcff18150e3297d723f17c69f81" \
--allow-insecure-unlock \
console 
# 2>>gethConsole.log
# --bootnodes "enode://ae795f7ed38aef3a29763216ec3673bcffa38938a77bc9811a2d56977bc495d8d93d6b4536d21b01b5f67662cea07dd89feabb0354abc42ecc33bd918a4573c2@127.0.0.1:30303" \
# --port 30304 \
# --seanet \
# console 2>>gethConsole.log
# --nodiscover \
# --verbosity 5 \
# --v5disc \
# --bootnodes "enode://7c67468cce9ed7d29b790f506c04ac010b2c1e9d0dd17384056183a9657747b4e6803182dce5fa75dcd4632a8ab7e3a0a249f963ceae6c2008de46083828e50d@127.0.0.1:0?discport=30301" \