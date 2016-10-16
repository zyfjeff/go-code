package main

func main(){
    x := []int{1,2,3}
    switch {
        case x[1] > 0:
            println("a")
        case x[1] < 0:
            println("b")
        default:
            println("c")
    }
}
