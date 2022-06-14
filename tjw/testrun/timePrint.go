package main

import (
	"fmt"
	mrand "math/rand"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

// func nextRevalidateTime() time.Duration {
// 	tab.mutex.Lock()
// 	defer tab.mutex.Unlock()

// 	return time.Duration(tab.rand.Int63n(int64(revalidateInterval)))
// }
var (
	revalidateInterval = 10 * time.Second
)

// type logger struct {
// 	l log.Logger
// 	// t time.Time
// }

func main1() {
	// fmt.Println()
	// t := time.Timer
	// var Logger logger

	rand := mrand.New(mrand.NewSource(0))

	dur := time.Duration(rand.Int63n(int64(revalidateInterval)))
	// str := dur.String()
	// Logger.l.Info("time duration: ")
	log.Info(dur.String())
	fmt.Println(dur.String())
	// log.Info()
	// Logger.l.Info("Revalidated node", "b", "id", "checks")
}
