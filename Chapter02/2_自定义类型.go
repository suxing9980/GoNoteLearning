package main

import "fmt"

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
}
