package main

import "fmt"
import "time"

//bool的chan，睡眠一段时间
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}


func main() {
    done := make(chan bool,1)
    go worker(done)

    <-done
}
