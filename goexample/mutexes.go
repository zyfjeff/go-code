package main

import (
    "fmt"
    "math/rand"
    "runtime"
    "sync"
    "sync/atomic"
    "time"
)

func main() {
    var state = make(map[int]int)
    //var mutex = &sync.Mutex{}
    var mutex = sync.Mutex{}
    var ops int64 = 0
    //100个协程，每个协程产生随机数，统计各个数字出现的频率，通过mutex包含map,ops统计总个数
    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                atomic.AddInt64(&ops,1)
                runtime.Gosched()
            }
        }()
    }

    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddInt64(&ops,1)
                runtime.Gosched()
            }
        }()
    }

    time.Sleep(time.Second)
    //打印次数
    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:",opsFinal)
    mutex.Lock()
    fmt.Println("state:",state)
    mutex.Unlock()
}
