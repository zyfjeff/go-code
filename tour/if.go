package main

import (
    "fmt"
    "math"
)

func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }
    //Sprint 数字转字符串
    return fmt.Sprint(math.Sqrt(x))
}

func main() {
    fmt.Println(sqrt(2),sqrt(-4))
}
