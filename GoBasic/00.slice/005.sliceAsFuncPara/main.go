package main

import "fmt"

/**
1.f函数里的操作直接更改了slice的底层数据,所以会影响到实参s。
因为直接s[i]=s[i]+1,切片的底层结构体有指针元素,所以直接更改了slice的底层数据。
*/

func main() {
	s := []int{1, 1, 1}
	f(s)
	fmt.Println(s)
}

func f(s []int) {
	for i := range s {
		s[i] += 1
	}
}
