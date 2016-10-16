package main

import "fmt"

func main() {
    //具有二个缓存的channel
    queue := make(chan string,2)
    //连续发送数据
    queue <- "one"
    queue <- "two"
    //然后关闭，关闭后居然还可以收数据,6666
    close(queue)
    //range遍历channel
    for elem := range queue {
        fmt.Println(elem)
    }
}
