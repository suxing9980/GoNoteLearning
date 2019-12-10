package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func makeSlice() []int {
	s := make([]int, 0, 10)
	fmt.Println(s, cap(s), len(s), reflect.TypeOf(s))
	s = append(s, 10)
	fmt.Println(s, cap(s), len(s), reflect.TypeOf(s))
	return s
}

func makmap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m)
	return m
}

func main() {
	a, _ := strconv.ParseInt("1100100", 2, 32)
	b, _ := strconv.ParseInt("0144", 8, 32)
	c, _ := strconv.ParseInt("64", 16, 32)
	fmt.Println(a, b, c)

	fmt.Println("0b" + strconv.FormatInt(a, 10))
	fmt.Println("0" + strconv.FormatInt(b, 8))
	fmt.Println("0x" + strconv.FormatInt(c, 16))
	fmt.Println("******************************")
	makeSlice()
	makmap()
}
