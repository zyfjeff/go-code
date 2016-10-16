package main

type Color int

const (
    Black Color = iota
    Red
    Blue
)

func test(c Color) {}

func main() {
    c := Black
    test(c)
    x := 1
    test(x) //类型不匹配,只能使用Color类型
    test(1) //常量会自动转换
}
