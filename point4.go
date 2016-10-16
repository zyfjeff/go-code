package main


import (
    "fmt"
    "unsafe"
)

func main() {
    d := struct {
        s   string
        x   int
    }{"abc",100}

    p := uintptr(unsafe.Pointer(&d))//转换为数值进行运算
    p += unsafe.Offsetof(d.x)

    p2 := unsafe.Pointer(p) //转换为Pointer来索引，取值
    px := (*int)(p2)
    *px = 200   //d.x = 200

    fmt.Printf("%#v",d)
}
