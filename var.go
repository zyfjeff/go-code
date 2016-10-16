package main

var a = 0

func main(){
    a :=  3
    println(&a)
    {
        a := 4
        println(&a)
    }
    println(&a)
}
