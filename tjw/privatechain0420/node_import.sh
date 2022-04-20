#!/bin/bash

GETHPATH=/home/tjw/CodeRoot/go-ethereum/build/bin
# 此处文件名需要修改，在稳定版内export时可直接赋为没有时间的
geth import ./backup/blockchain_backup_2022-03-31-16:00 --datadir . --syncmode full

