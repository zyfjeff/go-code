package main

import "fmt"

//只写的chan，只能呢过写入
func ping(pings chan<- string,msg string) {
    pings <- msg
}

//pings是只读的，pongs是只写的
func pong(pings <-chan string,pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string,1)
    pongs := make(chan string,1)
    //发送信息
    ping(pings,"passed message")
    //接收信息
    pong(pings,pongs)
    //读取
    fmt.Println(<-pongs)
}
