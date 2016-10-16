package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

//错误检查的方法
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	service := ":1200" //解析成127.0.0.1:1200
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	//监听ip地址，开始accep链接，然后打印时间返回
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}
