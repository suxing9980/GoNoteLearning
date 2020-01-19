package main

import "fmt"

// @Time    : 2020/1/19 10:04
// @Author  : Little柯南
// @File    : 1_name.go

func main() {
	s := "雨痕\x61\142\u0041"
	fmt.Printf("%v\n", s)
	fmt.Printf("%v  %x\n", len(s), s)
}
