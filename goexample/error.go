package main

import "errors"
import "fmt"

//返回一个错误,通过errors.New创建一个错误对象
func f1(arg int) (int,error) {
    if arg == 42 {
        return -1,errors.New("Can't work with 42")
    }

    return arg + 3,nil
}

type argError struct {
    arg int
    prob string
}

//给结构体定义了一个错误
func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s",e.arg,e.prob)
}

//只要实现了Error方法都可以通过error接受，error就是一个接口
func f2(arg int) (int,error) {
    if arg == 42 {
        return -1,&argError{arg,"can't work with it"}
    }
    return arg + 3,nil
}

func main() {
    //遍历slice，对每个值调用f1和f2的方法
    for _,i := range []int{7,42} {
        if r,e := f1(i); e != nil {
            fmt.Println("f1 failed:",e)
        } else {
            fmt.Println("f1 worked:",r)
        }
    }
    for _,i := range []int{7,42} {
        if r,e := f2(i); e!= nil {
            fmt.Println("f2 failed:",e)
        } else {
            fmt.Println("f2 workded:",r)
        }
    }
}
