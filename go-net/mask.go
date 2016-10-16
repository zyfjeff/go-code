package main

import (
	"fmt"
	"net"
	"os"
)

//ParseIP ip解析，解析字符串的ip地址为type ip[]
//type ip[] 含有DefaultMask Mask Size等方法，用来获取默认掩码 网络位 掩码长度和网络位长度等
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	dotAddr := os.Args[1]
	//ip解析
	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}
	//默认掩码
	mask := addr.DefaultMask()
	//网络位
	network := addr.Mask(mask)
	//掩码长度,网络位长度
	ones, bits := mask.Size()
	fmt.Println("Address is ", addr.String(),
		"Default mask length is ", bits,
		"Leading ones count is ", ones,
		"Mask is (hex) ", mask.String(),
		" Network is ", network.String())
	os.Exit(0)
}
