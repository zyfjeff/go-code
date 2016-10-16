package main

import "time"
import "fmt"

func main() {
    //创建了定时器,每time.Millisecond * 500就产生事件，发送到chnanel
    ticker := time.NewTicker(time.Millisecond * 500)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at",t)
        }
    }()
    //睡上一段事件，然后关闭
    time.Sleep(time.Millisecond * 1600)
    ticker.Stop()
    fmt.Println("Ticker stopped")
}

