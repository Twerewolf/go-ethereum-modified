# devp2p simulation examples

## ping-pong

`ping-pong.go` implements a simulation network which contains nodes running a
simple "ping-pong" protocol where nodes send a ping message to all their
connected peers every 10s and receive pong messages in return.

To run the simulation, run `go run ping-pong.go` in one terminal to start the
simulation API and `./ping-pong.sh` in another to start and connect the nodes:

```
$ go run ping-pong.go
INFO [08-15|13:53:49] using sim adapter
INFO [08-15|13:53:49] starting simulation server on 0.0.0.0:8888...
```

```
$ ./ping-pong.sh
---> 13:58:12 creating 10 nodes
Created node01
Started node01
...
Created node10
Started node10
---> 13:58:13 connecting node01 to all other nodes
Connected node01 to node02
...
Connected node01 to node10
---> 13:58:14 done
```

Use the `--adapter` flag to choose the adapter type:

```
$ go run ping-pong.go --adapter exec
INFO [08-15|14:01:14] using exec adapter                       tmpdir=/var/folders/k6/wpsgfg4n23ddbc6f5cnw5qg00000gn/T/p2p-example992833779
INFO [08-15|14:01:14] starting simulation server on 0.0.0.0:8888...
```
## TJW补充
1. 在/cmd/p2psim目录下编译p2psim程序，程序名修改为p2psim
    ```
    go build p2psim.go
    ```
2. 将执行文件copy到$PATH的范围内能直接全局使用，我使用的路径为/usr/local/go/bin/之下
3. 按照上面的操作说明进行，首先开一个terminal1 执行go run ping-pong.go进入执行状态，然后在terminal2执行 ./ping-pong.sh