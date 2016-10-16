package main

func main() {
a := []int{0,0,0}   //直接提供初始化表达式
a[1] = 10

b := make([]int,3)  //makeslice,初始化成员结构,返回对象，并非指针
b[1] = 10

c := new([]int)     //返回的是指针，所以c[1] = 10是错误的
c[1] = 10
}
