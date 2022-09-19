package main

import "fmt"

/*
1.函数myAppend->因为传递的是slice,并且是append等非直接改变slice底层数据的操作,所以不会改变实参。
2.函数myAppendPtr->因为传递的是slice的指针,所有改变了实参。
*/

func myAppend(s []int) []int {
	// 这里 s 虽然改变了，但并不会影响外层函数的 s
	s = append(s, 100)
	return s
}

func myAppendPtr(s *[]int) {
	// 会改变外层 s 本身
	*s = append(*s, 100)
	return
}

func main() {
	s := []int{1, 1, 1}
	newS := myAppend(s)

	fmt.Println(s)
	fmt.Println(newS)

	s = newS

	myAppendPtr(&s)
	fmt.Println(s)
}
