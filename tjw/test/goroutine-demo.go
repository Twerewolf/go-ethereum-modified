package main

import (
	"fmt"
	"time"
)

//hello 和 hi 两个func做main 的协程，执行先后顺序不固定
func hello() {
	fmt.Println("Hello world goroutine")
	for i := 0; i < 100; i++ {
		// str := strconv.Itoa('a' + i)
		// r := rune('a')
		fmt.Println("a")
	}
}
func hi() {
	fmt.Println("Hi")
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}
func main1() {
	go hi()
	fmt.Println("main function")
	go hello()
	// fmt.Println("goroutine 显示出来的可能是随机的？")

	// fmt.Println("增加等待")
	time.Sleep(time.Microsecond) //增加sleep后所有协程都能够执行结束
	// fmt.Println("end")
}
