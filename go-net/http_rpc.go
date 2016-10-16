package main

//go的rpc可以基于HTTP,TCP,
//go rpc的一些限制
// 函数必须是公有的，也就是首字符要大写
// 至少有两个参数，第一个参数是一个指针指向客户端传递给服务器端的参数，第二个参数也是一个指针是指向服务器端返回给客户端的参数
// 有一个返回值类型为os.Error

// F(&T1,&T2) os.Error

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

//封装了参数和返回值
type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
