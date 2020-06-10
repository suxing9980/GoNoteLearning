package main

import (
	"errors"
	"fmt"
)

// @Time    : 2020/1/17 17:44
// @Author  : Little柯南
// @File    : 6_error.go

// 返回值有命名，那么只写一个return即可，具体表现以返回值的实际赋值为准
func divZero(num, num2 int) (ret int, err error) {
	if num2 == 0 {
		ret = 0
		err = errors.New("Error:division by zero")
	} else {
		ret = num / num2
		err = nil
	}
	return
}

func main() {
	a, err := divZero(12, 2)
	fmt.Println(a, err)
}
