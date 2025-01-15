package main

import (
	"fmt"
	// "gocode/testproject01/test"
)

func main() {
	// fmt.Println("hello world")
	// fmt.Println(test.StuNo)
	// var sum int = 0
	// for i := 1; i <= 5; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)
	// for {
	// 	fmt.Println(sum)
	// }
	// var str string= "hello 你好"
	// for i, value := range str {
	// 	// fmt.Println(i)
	// 	// fmt.Println(value)
	// 	fmt.Printf("i = %d, value = %c\n", i, value)
	// }

	lable1 : 
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == 2 && j == 2 {
				break lable1
			}
		}
	}
}
