package main

import (
    "fmt"
    "net"
    "os"
)

//ResolveIPAddr 解析ip地址，转换为type ip[]形式
func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
        fmt.Println("Usage: ", os.Args[0], "hostname")
        os.Exit(1)
    }

    name := os.Args[1]
    //解析ip地址，将ip
    addr, err := net.ResolveIPAddr("ip", name)
    if err != nil {
        fmt.Println("Resolution error", err.Error())
        os.Exit(1)
    }
    fmt.Println("Resolved address is ", addr.String())
}
