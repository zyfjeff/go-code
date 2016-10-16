package main

import (
    "encoding/json"
    "log"
    "os"
)

type KeyValue struct {
    Key string      `json:key`
    Value string    `json::value`
}

func main() {
    file,err  := os.Create("test")
    if err != nil {
        log.Println("Create file failure")
        return
    }

    enc := json.NewEncoder(file)
    enc.Encode(KeyValue{Key:"abc",Value:"def"})
    enc.Encode(KeyValue{Key:"abd",Value:"dee"})
    enc.Encode(KeyValue{Key:"abe",Value:"deg"})
    enc.Encode(KeyValue{Key:"abf",Value:"dea"})

    file.Close()
}
