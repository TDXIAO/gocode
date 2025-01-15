package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func writeData(intChan chan int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		intChan<- i
		fmt.Println("写入的数据为：", i)
	}

	close(intChan)
}

func main() {
	intChan := make(chan int, 10)
	wg.Add(1)
	go writeData(intChan)

	wg.Wait()	
}
