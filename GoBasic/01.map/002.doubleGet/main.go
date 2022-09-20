package main

import "fmt"

/**
两种从map种get数据的方法
*/
func main() {
	ageMp := make(map[string]int)
	ageMp["a"] = 18
	ageMp["b"] = 19

	//不带comma用法
	age1 := ageMp["a"]
	fmt.Println(age1)

	//带comma用法
	age2, ok := ageMp["b"]
	fmt.Println(age2, ok)
}
