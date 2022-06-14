package main

import (
	"fmt"
	"strconv"
)

func parseBool(s string) bool {
	f, _ := strconv.ParseBool(s)
	return f
}

func main() { //一个main包里也只能由一个main函数
	str := "2"
	fmt.Println("the string is: " + str)
	b := parseBool("2")
	if b == true {
		fmt.Println("it's true")
	} else {
		fmt.Println("it's not a correct string")
	}

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
}
