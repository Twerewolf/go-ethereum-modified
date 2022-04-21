#!/bin/bash
#constraint setting can not have space between '='
#need more than just config
# need mine log settings
# using 提前生成好的config文件
# --trace value                       Write execution trace to the given file
GETHPATH=../../build/bin

$GETHPATH/geth \
--config ./geth_config.toml \
--mine \
--miner.threads 1 \
--log.json \
--verbosity 5 \
console 2>>gethConsole.log
