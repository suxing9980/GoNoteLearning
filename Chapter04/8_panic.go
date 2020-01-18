package main

import (
	"fmt"
	"log"
)

// @Time    : 2020/1/18 15:07
// @Author  : Little柯南
// @File    : 8_panic.go

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		}
	}()

	panic("something")
	fmt.Println("exit.")

}
