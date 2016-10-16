package main

import "time"
import "fmt"

func main() {
    c1 := make(chan string,1)
    //睡眠2秒，发送
    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "result 1"
    }()

    select {
    case res := <-c1:
        fmt.Println(res)
    //超时时间
    case <- time.After(time.Second * 1):
        fmt.Println("timeout 1")
    }

    c2 := make(chan string,1)
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "result 2"
    }()

    select {
    case res := <-c2:
        fmt.Println(res)
    //3秒超时
    case <- time.After(time.Second * 3):
        fmt.Println("timeout 2")
    }
}
