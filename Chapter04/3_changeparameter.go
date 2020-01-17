package main

import "fmt"

// @Time    : 2020/1/17 10:35
// @Author  : Little柯南
// @File    : 3_changeparameter.go

func test(s string, a ...int) {
	fmt.Printf("%T,%v\n", a, a) // 类型与值
}

func main() {
	test("student", 1, 2, 3, 4)
}
