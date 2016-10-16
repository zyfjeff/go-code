package main

import "fmt"

type rect struct {
    width,height int
}

func (r *rect) area() int {
    return r.width * r.height
}

func (r rect) perim() int {
    r.width = 100;
    return 2 * r.width + 2*r.height
}


func main() {
    r := rect{width: 10,height: 5}
    fmt.Println("area: ",r.area())
    fmt.Println("perlim: ",r.perim())
    fmt.Println(r)

    rp := &r
    fmt.Println("area: ",rp.area())
    fmt.Println("perlim: ",rp.perim())

}

