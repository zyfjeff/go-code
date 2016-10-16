package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {
    //os.Signal类型的chn
    sigs := make(chan os.Signal,1)
    done := make(chan bool,1)
    //通过Norify来注册信号，
    signal.Notify(sigs,syscall.SIGINT,syscall.SIGTERM)

    //协成来收集信号，然后发送done chan来表示完成
    go func() {
        sig :=  <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}
