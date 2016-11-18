package main

import "fmt"

func main() {
  x, y := 1, 2
  x, y = y + 3, x + 2
  fmt.Println(x, y)
}
