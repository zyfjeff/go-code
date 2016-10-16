package main

import "fmt"
import "os"

func main() {
    f := createFile("/tmp/defer.txt")
    defer closeFile(f)//在main函数执行完成后，自动执行closeFile
    writeFile(f)
}
//返回的是一个指针?
func createFile(p string) *os.File {
    fmt.Println("creating")
    f,err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f,"data")
}

func closeFile(f *os.File) {
    fmt.Println("closing")
    f.Close()
}
