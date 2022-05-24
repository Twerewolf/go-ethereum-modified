package main

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

func main() {
	go func() { //ping协程
		for range time.Tick(1 * time.Second) {
			log.Info("sending ping")
			fmt.Println("sending ping")
			// if err := p2p.Send(rw, pingMsgCode, "PING"); err != nil {
			// 	errC <- err
			// 	return
			// }
		}
	}()
	go func() {
		time.Sleep(10 * time.Second)
	}()
	fmt.Println("test")
}
