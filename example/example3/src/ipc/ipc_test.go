package ipc
import (
    "testing"
)

type EchoServer struct {

}

func (server *EchoServer) Handle(method,request string) *Response {
    return &Response{"200","ECHO:"+request}
}

func (server *EchoServer) Name() string {
    return "EchoServer"
}

func TestIpc(t *testing.T) {
    server := NewIpcServer(&EchoServer{})
    client1 := NewIpcClient(server)
    client2 := NewIpcClient(server)

    resp1,err := client1.Call("get","From Client")
    if err != nil {
        t.Error("client call server error")
    }
    resp2,err := client1.Call("get","From Client2")
    if err != nil {
        t.Error("client call server error")
    }

    if resp1.Body != "ECHO:From Client" || resp2.Body != "ECHO:From Client2" {
        t.Error("IpcClient.Call failed. resp1:",resp1,"resp2:",resp2)
    }

    client1.Close()
    client2.Close()
}
