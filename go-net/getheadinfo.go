package main

import (
	"fmt"
	"io/ioutil"
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

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	//解析TCP地址 ,只要传入host:port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	//发起TCP请求
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	//写入和读取返回的都是字节，这是需要注意的
	//写入内容
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	//读取返回信息
	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}
