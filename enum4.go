package main


const (
    A = iota            //iota 是 0
    B                   //iota 是 1
    C = "c"             //iota被打断
    D                   //c
    E = iota            //显式恢复了 是4
    F                   //iota 是5
)

func main(){

}
