package main

import "fmt"

func main() {
    jobs := make(chan int,5)
    done := make(chan bool)

    go func() {
        for {
            //通过more可以知道是否还有没有接收的消息
            j,more := <-jobs
            if more {
                fmt.Println("received job",j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()
    //发送消息
    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job",j)
    }
    //关闭channel
    close(jobs)
    fmt.Println("sent all jobs")
    //通过done来表明发送结束
    <-done
}
