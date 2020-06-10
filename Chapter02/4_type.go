package main

import "fmt"

type flags byte

const (
	read flags = 1 << iota
	write
	exec
)

// type组
type (
	user struct {
		name string
		age  uint8
	}

	event func(string) bool
)

func main() {

	var f flags = read | exec
	fmt.Printf("%b\n", f)
	fmt.Println("****************")
	//var u user
	//u.name = "tom"
	//u.age = 20
	u := user{age: 20, name: "Tom"}
	fmt.Println(u)
	var fu event = func(s string) bool {
		fmt.Println(s)
		return s != ""
	}
	fmt.Println(fu("tom"))
}
