#!/bin/bash
# 将已经运行生成的区块链数据进行备份输出
# GETHPATH=/home/tjw/CodeRoot/go-ethereum/build/bin
GETHPATH=../../build/bin
NAME=blockchain_backup_$(date +%Y-%m-%d-%H:%M)
$GETHPATH/geth export $NAME --datadir . --syncmode full
# geth export blockchain_backup --datadir ./data0/ --syncmode full
mv $NAME ./backup