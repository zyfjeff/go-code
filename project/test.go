package main

import (
	"fmt"
)

type int_t int

func (t int_t) String() string {
	return "dwdewde"
}

func main() {

	//　方法一: 构建二维数组
	/*
		picture := make([][]uint8, 10)
		for i := range picture {
			picture[i] = make([]uint8, 8)
		}
		fmt.Println(picture)
	*/

	//	方法二: 构建二维数组
	picture := make([][]uint8, 10)
	pixels := make([]uint8, 80)
	for i := range picture {
		picture[i], pixels = pixels[:8], pixels[8:]
	}

	tt := int_t(10)

	fmt.Println(picture)
	fmt.Println(tt)
}
