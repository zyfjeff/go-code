package main

import "bytes"
import "fmt"
import "regexp"

func main() {
    //正则表达式匹配字符串
    match,_ := regexp.MatchString("p([a-z]+)ch","peach")
    //直接打印匹配的结果是否匹配到
    fmt.Println(match)
    //编译了一个正则表达式
    r,_ := regexp.Compile("p([a-z]+)ch")
    //进行匹配
    fmt.Println(r.MatchString("peach"))
    //查找匹配的字符串
    fmt.Println(r.FindString("peach punch"))
    fmt.Println(r.FindStringIndex("peach punch"))

    fmt.Println(r.FindStringSubmatch("peach punch"))
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))
    fmt.Println(r.FindAllString("peach punch pinch",-1))
    fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch",-1))

    fmt.Println(r.FindAllString("peach punch pinch",2))
    fmt.Println(r.Match([]byte("peach")))

    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println(r)
    fmt.Println(r.ReplaceAllString("a peach","<fruit>"))

    in := []byte("a peach")
    out := r.ReplaceAllFunc(in,bytes.ToUpper)
    fmt.Println(string(out))
}
