package main

import "fmt"

func main() {
	var balance []int
	for i := 0; i < 10; i++ {
		balance = append(balance, 100+i)
	}
	// var res []int
	balance = balance[:3] //从0到0获取，即一个都不保留
	size := len(balance)
	fmt.Println("size:", size)
	for num, bal := range balance {
		fmt.Println(num, bal)
	}

}
