#!/bin/bash

GETHPATH=../../build/bin

$GETHPATH/geth \
--port 30305 \
--miner.threads 1 \
--datadir . \
--networkid 224 \
--ipcpath ./blockchain/geth.ipc \
--http  \
--http.port 8547 \
--miner.gasprice 1 \
--miner.etherbase "e7fb044503784dcff18150e3297d723f17c69f81" \
--allow-insecure-unlock \
console 2>>gethConsole.log
# --bootnodes "enode://ae795f7ed38aef3a29763216ec3673bcffa38938a77bc9811a2d56977bc495d8d93d6b4536d21b01b5f67662cea07dd89feabb0354abc42ecc33bd918a4573c2@127.0.0.1:30303" \

# --bootnodes "enode://ae795f7ed38aef3a29763216ec3673bcffa38938a77bc9811a2d56977bc495d8d93d6b4536d21b01b5f67662cea07dd89feabb0354abc42ecc33bd918a4573c2@127.0.0.1:30303" \