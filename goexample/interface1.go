package main

import (
    "fmt"
)

func PrintAll(vals []interface{}) {
    for _,val := range vals {
        fmt.Println(val)
    }
}

func main() {
    names := []string{"test1","test2","test3"}
    vals := make([]interface{},len(names))
    for i,v := range names {
        vals[i] = v
    }
    PrintAll(vals)
}
