package main

import "fmt"
import "sort"

//string slice的别名，给这个别名struct 添加方法
type ByLength []string

func (s ByLength) Len() int {
    return len(s)
}

func (s ByLength) Swap(i,j int) {
    s[i],s[j] = s[j],s[i]
}

func (s ByLength) Less(i,j int) bool {
    return len(s[i]) < len(s[j])
}

//sort接口需要实现 Swap Less和len即可
func main() {
    fruits := []string{"peach","banana","kiwi"}
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)
}
