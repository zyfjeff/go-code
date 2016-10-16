package main

import (
	"fmt"
	"net"
	"os"
)

//错误检查的方法
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		//读取内容，如果失败就返回
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		//响应请求
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

func main() {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}
