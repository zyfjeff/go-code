package main

import "fmt"


type KeyValue struct {
    key string
    value string
}

func main() {
    name := []KeyValue{{key:"swqds",value:"wqd"}}
    for _,kv := range name {
        fmt.Println(kv.key)
    }
}
