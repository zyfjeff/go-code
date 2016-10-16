package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

//封装了参数和返回值
type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server")
		os.Exit(1)
	}

	serverAddress := os.Args[1]
	//发起http的rpc请求
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing: ", err)
	}
	//构造参数
	args := Args{17, 8}
	var reply int
	//远程调用，传递参数，接收返回值
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error: ", err)
	}

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("Arith error: ", err)
	}

	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
