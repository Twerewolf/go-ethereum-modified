# how to start a privatechain
0. be sure geth program is ready
1. run node_init.sh to initial the genesis state, written into geth database
2. run node_dumpconfig.sh to dump part of geth config into geth_config.toml file
3. run node_import.sh if there a backup 
4. remember to put your account key file into ./keystore/ dir
5. run node_start_localgeth2.sh
6. if need to keep an eye on the log, run log_monitor.sh in another terminal