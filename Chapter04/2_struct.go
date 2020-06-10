package main

import (
	"fmt"
	"log"
	"time"
)

// @Time    : 2020/1/17 10:03
// @Author  : Little柯南
// @File    : 2_struct.go

type serverOption struct {
	address string
	port    int
	path    string
	timeOut time.Duration
	log     *log.Logger
}

func newOption() *serverOption {
	return &serverOption{
		address: "0.0.0.0",
		port:    8080,
		path:    "/var/test",
		timeOut: time.Second * 5,
		log:     nil,
	}
}

func server(s *serverOption) {
	fmt.Println(s)
}

func main() {
	opt := newOption()
	opt.port = 8080
	server(opt)
}
