package main

import "fmt"

// @Time    : 2020/1/16 17:25
// @Author  : Little柯南
// @File    : 1_逃逸分析.go

func GetPoint() *int {
	var num int = 5
	return &num
}

func main() {
	var p = GetPoint()
	fmt.Println("p=", p, " *p=", *p)
}
