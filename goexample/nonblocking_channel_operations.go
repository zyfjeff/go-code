package main

import "fmt"

func main() {
    //默认是没有buffer的，只能阻塞的等待，因此非阻塞的情况下，导致send发送失败
    messages := make(chan string,1)
    signals := make(chan bool,1)

    //等待消息，select中是非阻塞的
    select {
    case msg := <-messages:
        fmt.Println("received message",msg)
    default:
        fmt.Println("no message received")
    }

    msg := "hi"
    select {
    case messages <-msg:    //为什么无法发送
        fmt.Println("sent message",msg)
    default:
        fmt.Println("no message sent")
    }

    select {
    case msg := <-messages:
        fmt.Println("received message",msg)
    case sig := <-signals:
        fmt.Println("received signal",sig)
    default:
        fmt.Println("no activity")
    }
}

