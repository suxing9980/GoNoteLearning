package main

import "fmt"

// @Time    : 2020/1/19 15:19
// @Author  : Little柯南
// @File    : 2_for.go

func main() {
	const words string = "什么是B恐怖支线剧情呢？很重要吗？"
	for i := 0; i < len(words); i++ {
		fmt.Printf("%d:[%c]", i, words[i])
	}
	fmt.Println()
	for k, v := range words {
		fmt.Printf("%d:[%c]", k, v)
	}
}
