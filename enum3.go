package main

const (
    //同一个常量组中，可以提供多个iota，各自独立增长
    A,B = iota,iota << 10   //0,0 << 10
    C,D                     //1,1 << 10
)


func main(){

}
