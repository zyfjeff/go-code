package main

import "fmt"
import "time"

func main() {
    i := 2
    //不换行
    fmt.Print("write ",i," as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }
    //time库很好用啊
    switch time.Now().Weekday() {
    case time.Saturday,time.Sunday:
        fmt.Println("it's the weekend")
    default:
        fmt.Println("it's a weekday")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("it's before noon")
    case false:
        fmt.Println("continue")
    default:
        fmt.Println("it's after noon")
    }
}
