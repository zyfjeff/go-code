package main

var a struct {x int `a`}
var b struct {x int `ab`}

func main(){
    b = a
}
