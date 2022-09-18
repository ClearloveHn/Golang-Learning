package main

import "fmt"

func main() {
	//定义切片
	var slice []int
	fmt.Println("slice是一个nil切片", slice == nil)

	slice1 := []int{}
	fmt.Println("slice1是一个empty切片", slice1 == nil)

}
