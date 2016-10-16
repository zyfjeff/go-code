package main
import "time"
import "fmt"

func main() {
    //创建了一个定时器，2秒后会发送事件到timer1.C channel
    timer1 := time.NewTimer(time.Second * 2)

    //等待定时器到期
    data := <-timer1.C
    fmt.Println("Timer 1 expired")
    fmt.Println("Timer 1 expired",data)
    timer2 := time.NewTimer(time.Second)
    go func() {
        <- timer2.C
        fmt.Println("Timer 2 expired")
    }()

    //关闭定时器
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}
