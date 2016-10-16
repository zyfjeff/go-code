package main

import (
    "fmt"
    "net"
    "os"
)

//LookupPort 传入网络的类型，服务名称(tcp http)返回对应的port
func main() {
    if len(os.Args) != 3 {
        fmt.Fprintf(os.Stderr,
            "Usage: %s network-type service\n",
            os.Args[0])
        os.Exit(1)
    }
    networkType := os.Args[1]
    service := os.Args[2]
    //传入网络类型，和服务名称，返回对应的port
    port, err := net.LookupPort(networkType, service)
    if err != nil {
        fmt.Println("Error: ", err.Error())
        os.Exit(2)
    }

    fmt.Println("service port ", port)
    os.Exit(0)
}
