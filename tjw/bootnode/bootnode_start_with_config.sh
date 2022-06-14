#!/bin/bash

# GETHPATH=../../build/bin
BOOTNODEPATH=../../cmd/bootnode

$BOOTNODEPATH/bootnode \
-nodekey $BOOTNODEPATH/mykey \
-addr 127.0.0.1:30086 \
-verbosity=5 \
-v5 \
-nat upnp:172.31.95.63
