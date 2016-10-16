package main

import "fmt"

// 传值
func zeroval(ival int) {
    ival = 0
}

//传引用
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {

    i := 1
    fmt.Println("initial:",i)
    zeroval(i)  //1
    fmt.Println("zeroval:",i) //1

    zeroptr(&i)
    fmt.Println("zeroptr:",i) //0
    fmt.Println("pointer:",&i)
}
