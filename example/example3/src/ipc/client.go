package ipc

import (
    "encoding/json"
)

// 发起一个请求，构建一个Request请求，然后使用json编码，最后通过chan传递给服务器端，
// 等待服务器响应，解码响应的数据构造Response响应体
type IpcClient struct {
    conn chan string
}
//链接
func NewIpcClient(server *IpcServer) *IpcClient {
    c := server.Connect()
    return &IpcClient{c}
}

func (client *IpcClient)Call(method,params string) (resp *Response,err error) {
    req := &Request{method,params}
    var b []byte
    b,err = json.Marshal(req)
    if err != nil {
        return
    }
    client.conn <- string(b)
    str := <-client.conn

    var resp1 Response
    err = json.Unmarshal([]byte(str),&resp1)
    resp = &resp1
    return
}

func (client *IpcClient)Close() {
    client.conn <- "CLOSE"
}
