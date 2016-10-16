package main


func main(){
    s := "abc"
    println(&s)

    s,y := "hello",20
    println(&s,y)
    {
        s,z := 1000,30
        println(&s,z)
    }
}
