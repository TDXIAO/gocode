package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now();
	fmt.Printf("月：%v\n", now.Month());
	fmt.Printf("月：%v\n", int(now.Month()));
	fmt.Printf("%v\n", "dfdfd")
}
