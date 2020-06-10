package main

import "fmt"

// @Time    : 2020/1/17 16:54
// @Author  : Little柯南
// @File    : 5_defer.go

func main() {
	x, y := 1, 2
	defer func(a int) {
		fmt.Println("defer x,y:", a, y)
	}(x)
	x += 100
	y += 200
	fmt.Println(x, y)

}
