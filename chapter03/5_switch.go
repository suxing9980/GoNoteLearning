package main

import "fmt"

// @Time    : 2019/12/13 9:39
// @Author  : Litte柯南
// @File    : 5_switch.go

func main() {
	a, b, c, x := 1, 2, 3, 2
	switch x {
	case a, b:
		fmt.Println("a|b")
	case c:
		fmt.Println("c")
	default:
		fmt.Println("4")
	}
}
