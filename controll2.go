package main


func main(){
    s := "abc"
    for i := range s {
        println(s[i])
    }

    for _,c := range s {
        println(c)
    }

    for range s{
        println("test")
    }

    m := map[string]int{"a":1,"b":2}
    for k,v := range m {
        println(k,v)
    }
}
