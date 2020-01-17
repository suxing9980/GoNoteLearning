package main

import "fmt"

// @Time    : 2020/1/17 10:47
// @Author  : Little柯南
// @File    : 4_sliceParameter.go

func mytest(a ...int) { // 可变长参数，本质是切片，但表现形式异于切片
	fmt.Println(a)
}

func mytest2(a [3]int) { // 数组作为参数
	fmt.Println(a)
}

func mytest3(a []int) { // 直接将切片作为参数
	fmt.Println(a)
}

func main() {
	a := [3]int{10, 20, 30}
	mytest(a[:]...)
	mytest2(a)
	mytest3(a[:])
}
