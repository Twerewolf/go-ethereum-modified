#!/bin/bash

LOGPATH=.

#-f 循环读取，用于查看递增的日志文件
tail -f $LOGPATH/gethConsole.log 
