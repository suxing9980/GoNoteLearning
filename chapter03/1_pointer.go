package main

import "fmt"

// @Time    : 2019/12/12 17:07
// @Author  : Litte柯南
// @File    : 1_pointer.go
func main() {
	a := struct {
		x int
	}{}
	a.x = 10
	p := &a
	p.x += 100
	fmt.Println(*p)
}
