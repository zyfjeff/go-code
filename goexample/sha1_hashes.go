package main

import "crypto/sha1"
import "fmt"

func main() {
    s := "sha1 this string"
    h := sha1.New()
    h.Write([]byte(s))  //强转为字节流写入
    bs := h.Sum(nil)    //参数用于追加已经存在的序列到h中

    fmt.Println(s)
    fmt.Printf("%x\n",bs)
}
