package main

import (
	"fmt"
)

func tet(tt interface{}) {

	switch tt := tt.(type) {
	case int:
		fmt.Println("int", tt)
	case bool:
		fmt.Println("bool", tt)
	}
}

func main() {
	a := 10
	tet(a)
}
