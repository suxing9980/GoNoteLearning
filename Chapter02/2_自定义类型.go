package main

import (
	"fmt"
	"math"
)

type color byte

const (
	black color = iota
	green
	white
	pink
)

func test(c color) {
	fmt.Println("color = ", c)
}

//
func main() {

	fmt.Printf("%v %v %v %v\n", black, green, white, pink)
	test(100)
	var x color = 2
	test(x)

	a, b, c := 100, 0144, 0x64
	fmt.Println(a, b, c)
	fmt.Printf("0b%b %#o %#x \n ", a, b, c)
	fmt.Println(math.MinInt8, math.MaxInt8)
}
