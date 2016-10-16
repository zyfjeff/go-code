package main

const x,y int = 1,2 //多常量的初始化
const s = "Hello World!"

//常量组
const (
    a,b = 10,100
    c bool = false
    z   //不提供类型和初始化值，那么视作与上一常量相同
)


func main(){
    //未使用的局部变量不会引发编译错误
    const x = "xxx"
}
