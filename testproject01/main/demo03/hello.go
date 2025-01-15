package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p *Person) test() *Person {
	p.Name = "啦啦"
	fmt.Println(p.Name)
	p2 := Person{"haha"}
	return &p2
}

// func (p *Person) String() string {
// 	return fmt.Sprintf("Name = %v", p.Name)
// }

func (p Person) String() string {
	return fmt.Sprintf("Name = %v", p.Name)
}

func main() {
	p := Person{"丽丽"}
	p2 := p.test()
	fmt.Println(p2)
	fmt.Println(p.Name)
	fmt.Printf("p的地址为：%p\n", &p);
	fmt.Println(p);
}
