package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	name string
	age byte
}
type manager struct {
	user
	title string
}

func (u user)ToString() string  {
	return fmt.Sprintf("%+v", u)
}

func (u user)Print()  {
	fmt.Printf("%+v", u)
}

type Printer interface {
	Print()
}

type X int

func (x *X)inc()  {
	*x++
}

func task(id int32) {
	for i:=0;i<10;i++ {
		fmt.Printf("[%d : %d]", id, i)
		time.Sleep(time.Second)
	}
}
// go入口函数
func main() {
	var mm manager
	mm.user.age = 18
	mm.user.name = "sunmi"
	fmt.Println(mm.ToString())

	var p Printer = mm
	p.Print()

	fmt.Println()
	go task(12)
	go task(15)
	time.Sleep(time.Second * 11)

	//fmt.Println("Hello World")
	//var xs X
	//xs.inc()
	//fmt.Println(xs)
	//var mm manager
	//mm.user.age = 18
	//mm.user.name = "sunmi"
	//mm.title = "姓名"
	//fmt.Println(mm)
	//var num int32
	//var str string = "studying"
	//fmt.Println("num=",num," str=",str)
	//
	//for j := 0; j<5; j++ {
	//	fmt.Println(j)
	//}
	//fmt.Println("*************")
	//var i = 0
	//for i<5 {
	//	fmt.Println(i)
	//	i++
	//}
	//fmt.Println("*************")
	//var k = 5
	//for {
	//	if k < 0 {
	//		break
	//	}
	//	fmt.Println(k)
	//	k--
	//}
	//
	////x := []int{1,2,3}
	//var x  = []int{1,2,3}
 	//for index, number := range x{
	//	fmt.Println(index,"->", number)
	//}
	//
	//fmt.Println(div(3,2))
	//
	//var a, b int  = 10, 20
	//fmt.Println(a," ", b)
	//
	//var aaa = make(map[int]string)
	//aaa[1] = "at"
	//aaa[2] = "bt"
	//fmt.Println(aaa)
}

func div(a int, b int) (int, error)  {
	if b==0 {
		return 0, errors.New("0")
	}
	return a/b, nil
}
