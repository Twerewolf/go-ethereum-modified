package main

import (
	"fmt"
	"time"
)

func test1(ch chan int) {
	fmt.Printf("%T,%p\n", ch, ch)
}

func main() {
	ch1 := make(chan int)
	fmt.Printf("%T,%p\n", ch1, ch1)
	fmt.Printf("%d\n", ch1) //错误类型但能输出

	test1(ch1)
	time.Sleep(2 * time.Second)
}
