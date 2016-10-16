package main

import "fmt"

func main() {
    //创建了一个map,key是string，value是int
    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13
    //打印这个map，很方便，比C&C++ 方便多了,调试神器。
    fmt.Println("map:",m)

    // 得到value
    v1 := m["k1"]
    fmt.Println("v1: ",v1)
    fmt.Println("len:",len(m))

    //删除
    delete(m,"k2")
    fmt.Println("map",m)

    //得到,第二个参数表明是否存在这个键
    _,prs := m["k2"]
    fmt.Println("prs:",prs)

    n := map[string]int{"foo":1,"bar":2}
    fmt.Println("map:",n)
}
