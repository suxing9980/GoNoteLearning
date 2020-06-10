package main

// @Time    : 2019/12/13 9:11
// @Author  : Litte柯南
// @File    : 4_if_else.go

import . "fmt"

func main() {
	var x int = 3
	if x > 5 {
		println("a")
	} else if x > 0 && x < 5 {
		Println("b")
	} else {
		Println("c")
	}

	Println("**********************")

	// x,y局部变量范围包括整个ifelse范围
	if x, y := 1, 2; x > y {
		Println("x > y")
	} else {
		println("x < y")
	}
}
