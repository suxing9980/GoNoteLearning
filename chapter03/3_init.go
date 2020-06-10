package main

import "fmt"

// @Time    : 2019/12/12 18:24
// @Author  : Litte柯南
// @File    : 3_init.go

func main() {
	type data struct {
		x int
		s string
	}

	var a data = data{12, "Little"}
	b := data{13, "Hello"}
	c := []int{1, 2}
	d := []int{1, 2,
		3,
		4}
	var e = []int{2, 5, 6}
	fmt.Println(a, b, c, d, e)
}
