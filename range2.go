package main
import (
    "fmt"
)

//range对引用对象不会复制，引用类型包括slice map channel

func main(){
    s := []int{1,2,3,4,5}
    for i,v := range s{
        if(i == 0) {
            s = s[:3]   //对slice修改，不影响range,底层数据不会被复制
            s[2] = 100
            fmt.Println(s)
        }
        println(i,v)
    }
    fmt.Println(s)
}
