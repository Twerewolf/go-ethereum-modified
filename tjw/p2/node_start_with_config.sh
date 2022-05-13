#!/bin/bash

GETHPATH=../../build/bin

$GETHPATH/geth \
--port 30304 \
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
