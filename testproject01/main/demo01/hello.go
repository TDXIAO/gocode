package main

import (
	"fmt"
	"gocode/testproject01/test"
)

var num int = fun()

func fun() int {
	fmt.Println("fun函数执行...")
	return 10
}

func init() {
	fmt.Println("init函数执行...")
}

func main() {
	fmt.Println("main函数执行...")
	fmt.Println(test.StuNo)
}
