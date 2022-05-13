#!/bin/bash
# 使用备份文件进行还原
# GETHPATH=/home/tjw/CodeRoot/go-ethereum/build/bin
GETHPATH=../../build/bin
# 此处文件名需要修改，在稳定版内export时可直接赋为没有时间的
$GETHPATH/geth import ./backup/blockchain_backup_2022-03-31-16:00 --datadir . --syncmode full

