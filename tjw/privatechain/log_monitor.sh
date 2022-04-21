#!/bin/bash

LOGPATH=/home/tjw/CodeRoot/go-ethereum/tjw/privatechain

tail -f $LOGPATH/gethConsole.log #-f 循环读取，用于查看递增的日志文件