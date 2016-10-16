package main

import "fmt"
import "sync"
//import "runtime"

var counter int = 0

func Count(lock *sync.Mutex) {
    lock.Lock() //临界区
    counter++
    fmt.Println(counter)
    lock.Unlock()
}

func main() {
    lock := &sync.Mutex{} //初始化锁
    for i := 0; i < 10;i++ {    //并发执行
        go Count(lock)
    }

    for {
        lock.Lock()         //是都达到临界值
        c := counter
        lock.Unlock()
     //   runtime.Gosched()
        if c >= 10 {
            break;
        }
    }
}
