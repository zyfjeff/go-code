package main

import "fmt"

func main() {
    //产生一个slice，长度为3
    s := make([]string,3)
    fmt.Println("emp:",s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:",s)
    fmt.Println("get:",s[2])
    fmt.Println("len:",len(s))
    //向slice，append数据
    s = append(s,"d")
    s = append(s,"e","f")
    fmt.Println("apd:",s)

    //slice拷贝
    c := make([]string,len(s))
    copy(c,s)
    fmt.Println("cpy:",c)

    //slice切片功能
    l := s[2:5]
    fmt.Println("sl1:",l)

    l = s[:5]
    fmt.Println("sl2:",l)

    l = s[2:]
    fmt.Println("sl3:",l)
    //slice的另一种方式
    t := []string{"g","h","i"}
    fmt.Println("dcl:",t)
    //二维slice,确定一维的长度
    twoD := make([][]int,3)
    for i := 0; i < 3; i++ {
        twoD[i] = make([]int,5)
        for j := 0;j < 5;j++ {
            twoD[i][j] = i + j
        }
    }
    /*
    for i := 0; i < 3; i++ {
        innerlen := i + 1
        //具体给第二维分配长度
        twoD[i] = make([]int,innerlen)
        for j := 0; j < innerlen; j++ {
            twoD[i][j] = i + j
        }
    }
    */
    fmt.Println("2d: ",twoD)
}
