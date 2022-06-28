package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

//可以使用自己的节点，也可以使用infura
const (
	// url = "https://mainnet.infura.io/v3/b4c05366e4c14e8a8304f0690aeae0e8"
	url = "https://mainnet.infura.io/v3/3588b39d059342aa95ee52fc97f49d65"
	// wss = "wss://mainnet.infura.io/ws/v3/b4c05366e4c14e8a8304f0690aeae0e8"
	wss = "wss://mainnet.infura.io/ws/v3/3588b39d059342aa95ee52fc97f49d65"
)

// failed to dial: websocket: bad handshake (HTTP status 401 Unauthorized)别人链接
// failed to dial: parse " https://mainnet.infura.io/v3/3588b39d059342aa95ee52fc97f49d65": first path segment in URL cannot contain colon 我的链接

func watch() {
	backend, err := ethclient.Dial(url)
	if err != nil {
		log.Printf("failed to dial: %v", err)
		return
	}
	//json rpc？？
	rpcCli, err := rpc.Dial(wss)
	if err != nil {
		log.Printf("failed to dial: %v", err)
		return
	}
	gcli := gethclient.New(rpcCli)
	//订阅待定的交易到txchannel
	txch := make(chan common.Hash, 100)
	_, err = gcli.SubscribePendingTransactions(context.Background(), txch)
	if err != nil {
		log.Printf("failed to SubscribePendingTransactions: %v", err)
		return
	}
	for {
		select {
		case txhash := <-txch:
			tx, _, err := backend.TransactionByHash(context.Background(), txhash)
			if err != nil {
				continue
			}
			data, _ := tx.MarshalJSON()
			log.Printf("tx: %v", string(data))
		}
	}
}

func main() {
	go watch()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan //输出
}
