package main

import (
	"fmt"
	"net"
	"os"
	"reflect"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	//将ip地址转换为
	addr := net.ParseIP(name)

	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println(reflect.TypeOf(addr)) //反射,查看其类型
		fmt.Println("The address is", reflect.TypeOf(addr.String()))
	}
	os.Exit(0)
}
