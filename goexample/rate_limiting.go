package main

import "fmt"
import "time"

func main() {
    requests := make(chan int,5)
    //发送五条消息
    for i := 1; i <= 5;i++ {
        requests <- i
    }

    close(requests)
    //创建了定时器，然后遍历channel，打印信息
    limiter := time.Tick(time.Millisecond * 200)
    for req := range requests {
        <-limiter   //每隔time.Millisecond * 200，起到了限速的作用
        fmt.Println("requests",req,time.Now())
    }

    //创建了另外一个time.Time类似的channel
    burstyLimiter := make(chan time.Time,3)
    //发送三个
    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    go func() {
        for t := range time.Tick(time.Millisecond * 200) {
            burstyLimiter <- t //每200 * time.Millisecond 就发送一个事件到burstyLimiter
        }
    }()


    burstyRequests := make(chan int,5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i //发送五个数据
    }
    close(burstyRequests)
    for req := range burstyRequests { //现在开始限速读取
        <-burstyLimiter//通过Limiter限速
        fmt.Println("request",req,time.Now())
    }
}
