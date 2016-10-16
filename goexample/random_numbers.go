package main

import "time"
import "fmt"
import "math/rand"

func main() {
    //打印100以内的随机数
    fmt.Print(rand.Intn(100),",")
    fmt.Print(rand.Intn(100))
    fmt.Println()

    //打印一个随机的浮点数
    fmt.Println(rand.Float64())
    fmt.Print((rand.Float64() * 5 + 5),",")
    fmt.Print((rand.Float64() * 5) + 5)
    fmt.Println()

    //创建种子
    s1 := rand.NewSource(time.Now().UnixNano())
    //根据种子创建随机数发生器
    r1 := rand.New(s1)
    //产生100的整数
    fmt.Print(r1.Intn(100),",")
    fmt.Print(r1.Intn(100))
    fmt.Println()

    //创建另外一个随机种子
    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    fmt.Print(r2.Intn(100),",")
    fmt.Print(r2.Intn(100))
    fmt.Println()
    //再次创建
    s3 := rand.NewSource(42)
    r3 := rand.New(s3)
    fmt.Print(r3.Intn(100),",")
    fmt.Print(r3.Intn(100))
}
