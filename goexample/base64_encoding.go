package main
import b64 "encoding/base64"
import "fmt"


func main(){
    data := "abc123!?$*&()'-=@~"
    //base64标志的Encoding
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)
    sDec,_ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    //base64 url Encoding
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec,_ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}
