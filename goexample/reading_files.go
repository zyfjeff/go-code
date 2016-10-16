package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

//io/ioutil 和 io 以及os都有文件相关的操作
//一个错误检查的例子
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    //ioutil和os两种方法来打开文件
    //ioutil的ReadFile方法，不用open
    dat,err := ioutil.ReadFile("/tmp/dat")
    check(err)
    fmt.Print(string(dat))

    //调用os的系统调用
    f,err := os.Open("/tmp/dat")
    check(err)

    //申请一个slice使用,通过slice来接收
    b1 := make([]byte,5)
    n1,err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n",n1,string(b1))

    //通过Seek来移动文件指针
    o2,err := f.Seek(6,0)
    check(err)
    b2 := make([]byte,2)
    n2,err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n",n2,o2,string(b2))

    //通过io类来读取，参数是文件描述符
    o3,err := f.Seek(6,0)
    check(err)
    b3 := make([]byte,2)
    n3,err := io.ReadAtLeast(f,b3,2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n",n3,o3,string(b3))
    _,err = f.Seek(0,0)
    check(err)

    //通过带buf的bufio类也可以
    r4 := bufio.NewReader(f)
    b4,err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n",string(b4))
    f.Close()

}
