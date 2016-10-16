package main

import "strconv"
import "fmt"

func main() {
    //解析浮点型
    f,_ := strconv.ParseFloat("1.234",64)
    fmt.Println(f)
    //解析整型
    i,_ := strconv.ParseInt("123",0,64)
    fmt.Println(i)
    //解析整型,0表示根据字符串的前缀判断基数
    d,_ := strconv.ParseInt("0x1c8",0,64)
    fmt.Println(d)
    //解析无符号整型
    u,_ := strconv.ParseUint("789",0,64)
    fmt.Println(u)
    //10进制转换
    k,_ := strconv.Atoi("135")
    fmt.Println(k)
    _,e := strconv.Atoi("wat")
    fmt.Println(e)
}
