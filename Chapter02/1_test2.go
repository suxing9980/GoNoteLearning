package main

import (
	"fmt"
)

func main() {
	x := 100
	fmt.Println(&x, x)
	{
		x, y := 200, 300
		fmt.Println(&x, y)
	}
	const (
		e uint = 64
		y
		s = "abc"
		z
	)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)

	const (
		_  = iota
		KB = 1 << (iota * 10)
		MB
		GB
		TB
	)
	fmt.Printf("%v %v %v %v\n", KB, MB, GB, TB)
	fmt.Println("*********************************")
	const (
		_, _ = iota, iota * 10
		a1, b1
		a2, b2
		a3, b3
	)
	fmt.Printf("a1 = %v, b1 = %v\n", a1, b1)
	fmt.Printf("a2 = %v, b2 = %v\n", a2, b2)
	fmt.Printf("a3 = %v, b3 = %v\n", a3, b3)
	fmt.Println("*********************************")
	const (
		q         = iota
		w float32 = iota
		t
		r
	)
	fmt.Printf("q= %T %v\n", q, q)
	fmt.Printf("w= %T %v\n", w, w)
	fmt.Printf("t= %T %v\n", t, t)
	fmt.Printf("r= %T %v\n", r, r)

}
