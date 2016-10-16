package main

const (
    _   = iota                  //iota 是0
    KB   int64 = 1 << (10*iota) //此时iota是1
    MB
    GB
    TB
    //和KB的表达式相同，但是iota的值不一样
)

func main(){
    println(KB)
    println(MB)
    println(GB)
    println(TB)
}
