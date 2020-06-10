package main

// shift + f10 :执行
import "fmt"

// @Time    : 2019/12/12 17:11
// @Author  : Litte柯南
// @File    : 2_zero_size.go
func main() {
	var a, b struct{}
	// 对象的地址是否相等与具体版本有关
	fmt.Println(&a, &b)
	fmt.Println(&a == &b, &a == nil)
}
