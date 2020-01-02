package main

import "fmt"

// @Time    : 2019/12/13 13:57
// @Author  : Litte柯南
// @File    : 8_breakcontinue.go

func main() {
outer:
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if y > 2 {
				fmt.Println()
				continue outer
			}
			if x > 2 {
				break
			}
			fmt.Printf("x= %d y=%d ", x, y)
		}

	}
}
