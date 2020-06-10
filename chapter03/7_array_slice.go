package main

import "fmt"

// @Time    : 2019/12/13 11:30
// @Author  : Litte柯南
// @File    : 7_array_slice.go

func main() {
	data := [3]int{10, 23, 45}

	for i, x := range data {
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Printf("x= %d, data=%d\n", x, data[i])
	}
	fmt.Println("data1=", data)
	fmt.Println("***************************************")
	for i, x := range data[:] {
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Printf("x= %d, data=%d\n", x, data[i])
	}
	fmt.Println("data2=", data)
}
