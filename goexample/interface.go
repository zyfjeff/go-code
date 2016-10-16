package main

import "fmt"
import "math"

//rect和circle的共同接口，只要实现其中一个即可
type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width,height float64
}

type circle struct {
    radius float64
}

//rect实现了两个
func (r rect) area() float64 {
    return r.width * r.height
}

func (r rect) perim() float64 {
    return 2 * r.width + 2 * r.height
}

//circle实现了两个
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}


//通过接口g就可以抽象r和c
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3,height: 4}
    c := circle{radius: 5}
    measure(r)
    measure(c)
}
