package main

import "strings"
import "fmt"

//string slice和string，找到这个string的索引
func Index(vs []string,t string) int {
    for i,v := range vs {   //遍历string slice即可
        if v == t {
            return i
        }
    }
    return -1
}

//查看string是否在slice中
func Include(vs []string,t string) bool {
    return Index(vs,t) >= 0
}

//遍历string slice，对每个string调用func，如果返回true则返回true,有一真即是真
func Any(vs []string,f func(string) bool) bool {
    for _,v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

//有一个假即假
func All(vs []string,f func(string) bool) bool {
    for _,v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}

//返回符合条件的字符串slice
func Filter(vs []string,f func(string) bool) []string {
    vst := make([]string,0)
    for _,v := range vs {
        if f(v) {
            vst = append(vst,v)
        }
    }
    return vst
}

//返回每个字符串通过map运算后的结果
func Map(vs []string,f func(string) string) []string {
    vsm := make([]string,len(vs))
    for i,v := range vs {
        vsm[i] = f(v)
    }
    return vsm
}

func main() {

    var strs = []string{"peach","apple","pear","plum"}
    fmt.Println(Index(strs,"pear"))
    fmt.Println(Include(strs,"grape"))
    fmt.Println(Any(strs,func(v string) bool {
        return strings.HasPrefix(v,"p") //是否有p前缀,有一真为真，因此返回true
    }))

    fmt.Println(All(strs,func(v string) bool {
        return strings.HasPrefix(v,"p") //是否有p前缀，有一假即假，因此返回false
    }))

    fmt.Println(Filter(strs,func(v string) bool {
        return strings.Contains(v,"e")  //返回包含e
    }))

    fmt.Println(Map(strs,strings.ToUpper)) //大小写字符串转换
}
