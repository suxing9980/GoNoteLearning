package main

import (
	"log"
	"os"
)

// @Time    : 2020/1/17 16:54
// @Author  : Little柯南
// @File    : 5_defer.go

func main() {
	f, err := os.Open("./2_struct.go")
	defer f.Close()

	if err != nil {
		log.Fatalln(err)
	}

}
