package main

import "fmt"

// @Time    : 2019/12/13 10:18
// @Author  : Litte柯南
// @File    : 6_for_range.go

func main() {
	var data = [3]string{"a", "b", "c"}
	for index, item := range data {
		fmt.Println(index, item)
		fmt.Println(&index, &item)
	}
}
