package ipc

import (
    "encoding/json"
    "fmt"
)
//一次请求，包含了请求的方法和参数
type Request struct {
    Method string "method"
    Params string "params"
}

//一次响应包含了,响应的状态码和包体
type Response struct {
    Code string "code"
    Body string "body"
}

//
type Server interface {
    Name() string
    Handle(method,params string) *Response
}

type IpcServer struct {
    Server
}

func NewIpcServer(server Server) *IpcServer {
    return &IpcServer{server}
}

//返回一个session，包含了一个chan，返回session给客户端
func (server *IpcServer)Connect() chan string {
    session := make(chan string,0)
    go func(c chan string) {
        for{
            request := <-c
            if request == "CLOSE" { //关闭该链接
                break;
            }
            var req Request;
            err := json.Unmarshal([]byte(request),&req)
            if err != nil {
                fmt.Println("Invalid request format:",request)
            }
            resp := server.Handle(req.Method,req.Params)
            b,err := json.Marshal(resp)
            c <- string(b)
        }
    }(session)
    fmt.Println("A new session has been created successfully.")
    return session;
}
