package main

import "fmt"

// @Time    : 2019/12/12 16:35
// @Author  : Little柯南
// @File    : 5_自动转换.py

func main() {
	// 无显式声明的常量可以自动转换
	const s = 20
	var a int = 10
	var b = s + a
	fmt.Println(b)
	var c float32 = 1.32
	var d = s + c
	fmt.Println(d)
}
