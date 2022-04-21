#!/bin/bash
# 生成geth_config文件，代替每次terminal中输入大量参数设置
# GETHPATH=/home/tjw/CodeRoot/go-ethereum/build/bin
GETHPATH=../../build/bin

$GETHPATH/geth --datadir . --networkid 224 \
--port 30311 --nodiscover \
--ipcpath ./blockchain/geth.ipc \
--http  \
--miner.gasprice 1 \
--miner.etherbase "23ffb1da16604cfc0373feedbfa5fe60b928d5c7" \
--allow-insecure-unlock \
dumpconfig > ./geth_config.toml
# console 2>>gethConsole.log
# dumpconfig > ./geth_config
# --mine \
# --miner.threads 1 \
# --verbosity 5 \
