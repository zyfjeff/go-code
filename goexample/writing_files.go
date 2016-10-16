package main


import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    d1 := []byte("hello\ngo\n")
    //ioutil适合那些一次性读写的情况，不用打开和关闭
    //如果文件存在，不知道是追加还是覆写呢?
    err := ioutil.WriteFile("/tmp/dat1",d1,0644)
    check(err)

    //创建文件，返回文件句柄
    f,err := os.Create("/tmp/dat2")
    check(err)
    defer f.Close()

    //通过文件句柄写入
    d2 := []byte{115,111,109,101,10}
    n2,err := f.Write(d2)
    fmt.Printf("wrote %d bytes \n",n2)

    //写入一串字符串
    n3,err := f.WriteString("Writes\n")
    fmt.Printf("wrote %d bytes\n",n3)

    //同步数据到磁盘
    f.Sync()

    //bufio带用户缓冲的io操作
    w := bufio.NewWriter(f)
    n4,err := w.WriteString("buffered\n")
    fmt.Printf("wrote %d bytes\n",n4)
    w.Flush()
}
