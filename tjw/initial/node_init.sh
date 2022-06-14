# GETHPATH=/home/tjw/CodeRoot/go-ethereum/build/bin
GETHPATH=../../build/bin
$GETHPATH/geth --datadir . init genesis0307.json
# deleted the gaslimit setting in genesis file