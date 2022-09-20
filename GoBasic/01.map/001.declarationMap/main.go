package main

import "fmt"

/**
不能向nilMap里添加元素,会引起panic
*/

func main() {
	var ageMp map[string]int
	fmt.Println("ageMp是一个nilMap", ageMp == nil)

	ageMp1 := make(map[string]int)
	fmt.Println("ageMp1是一个emptyMap", ageMp1 == nil)
}
