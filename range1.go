package main

import (
    "fmt"
)


func main(){
    a := [3]int{1,2,3}
    for i,v := range a{         //i,v是a的复制品,range导致了对象的复制
        if i == 0 {
            a[1],a[2] = 999,999
            fmt.Println(a)
        }
        a[i] = v + 100
    }
    fmt.Println(a)
}
