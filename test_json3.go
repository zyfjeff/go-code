package main

import (
    "log"
    "os"
    "io"
    "fmt"
    "encoding/json"
)

type KeyValue struct {
    Key string
    Value string
}

func main() {
    file,err := os.Open("test")
    if err != nil {
        log.Fatal("oepn file failure")
        return
    }

    enc := json.NewDecoder(file)
    for {
        var body KeyValue
        if err := enc.Decode(&body); err == io.EOF {
            break;
        } else if err != nil {
            log.Fatal(err)
        }

        fmt.Println(body)
    }

}
