package main

import (
	"errors"
	"fmt"
)

// @Time    : 2020/1/18 9:19
// @Author  : Little柯南
// @File    : 7_errorVar.go

var errDivByZero = errors.New("division by zero")

func div(x, y int) (ret int, err error) {
	if y == 0 {
		ret = 0
		err = errDivByZero
	} else {
		ret = x / y
		err = nil
	}
	return
}

func main() {
	_, err := div(5, 0)
	if err == errDivByZero {
		fmt.Println(errDivByZero)
	}
}
