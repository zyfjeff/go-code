package main

import (
    "fmt"
    "unsafe"
)

const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(b)
)

func main(){
    fmt.Println("test")
}
