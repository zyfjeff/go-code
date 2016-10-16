package main
import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

type readOp struct {
    key int
    resp chan int
}


type writeOp struct {
    key int
    val int
    resp chan bool
}

func main() {
    var ops int64 = 0
    //创建了两个channel
    reads := make(chan *readOp)
    writes := make(chan *writeOp)

    go func() {
        var state = make(map[int]int)
        for {
            select {
            //等到读，然后获取状态,发送到chnnel
            case read := <-reads:
                read.resp <- state[read.key]
            //等到写，然后设置对应key和value
            case write := <-writes:
                state[write.key] = write.val
                //最后响应
                write.resp <- true
            }
        }
    }()

    for r := 0; r < 100; r++ {
        go func() {
            for {
                //创建read，然后写入channel,等到写入成功的响应，然后计数
                read := &readOp {
                    key: rand.Intn(5),
                    resp: make(chan int)}
                reads <- read
                <-read.resp
                atomic.AddInt64(&ops,1)
            }
        }()
    }

    for w := 0;w < 10;w++ {
        go func() {
            for {
                //创建write，然后写入channel，等到响应，统计个数
                write := &writeOp{
                    key: rand.Intn(5),
                    val: rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <-write.resp
                atomic.AddInt64(&ops,1)
            }
        }()
    }
    time.Sleep(time.Second)
    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:",opsFinal)
}
