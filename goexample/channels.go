package main

import "fmt"

func main() {
    //创建了一个chan
    messages := make(chan string)
    //输入一个ping
    go func() {messages <- "ping"} ()

    //接收
    msg := <- messages
    fmt.Println(msg)
}
