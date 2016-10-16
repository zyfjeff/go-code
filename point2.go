package main

import (
    "unsafe"
)

func main(){
    x := 0x12345678
    //z := &x //z无法++,做运算
    p := unsafe.Pointer(&x)
    n := (*[4]byte)(p)
    println(n[0])
    println(n[1])
    println(n[2])
    println(n[3])

}
