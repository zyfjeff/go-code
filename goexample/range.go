package main

import "fmt"

func main() {
    nums := []int{2,3,4}
    sum := 0
    //遍历 slice，第一个是索引值第二个是value
    for _,num := range nums {
        sum += num
    }


    fmt.Println("sum:",sum)
    for i,num := range nums {
        if num == 3 {
            fmt.Println("index:",i)
        }
    }
    //遍历map
    kvs := map[string]string{"a":"apple","b":"banana"}
    for k,v := range kvs {
        fmt.Printf("%s -> %s\n",k,v)
    }

    //遍历字符串,打印的是ascii,可以通过string 强制转换
    for i,c := range "go" {
        fmt.Println(i,string(c))
    }
}
