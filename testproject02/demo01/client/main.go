package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {
	fmt.Println("客户端启动了。。")
	conn,err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("客户端连接失败，err:", err)
		return
	}
	fmt.Println("连接成功，conn:", conn)

	reader := bufio.NewReader(os.Stdin)

	str, err2 := reader.ReadString('\n')
	if err2 != nil {
		fmt.Println("终端输入失败，err2:", err2)
	}

	n, err3 := conn.Write([]byte(str))
	if err3 != nil {
		fmt.Println("连接失败， err3:", err3)
	}

	fmt.Println("终端数据通过客户端发送成功，一共发送了%d字节的数据,并退出", n)

}