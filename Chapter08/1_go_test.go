package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello world!" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test()
	func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Hello goland!" + strconv.Itoa(i))
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second)
}
