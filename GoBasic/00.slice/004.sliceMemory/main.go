package main

import "fmt"

/**
1.s的长度和容量都为1,第一个append之后,s扩容,容量变为2,此时s=[5,7]

2.第一个append之后,s扩容,容量变为4,此时s=[5,7,9]

3.第三个append之后,s还有容量,不必扩容,此时x=[5,7,9,11],x指向的底层数组变为[5,7,9,11],x和s指向同一个底层数组。
s此时不变的原因是append函数执行完后，返回的是一个全新的 slice，并且对传入的 slice 并不影响,
并且append的返回值不是s。

4.最后一个append之后,此时s=[5,7,9],有容量不必扩容,，所以直接在底层数组索引为3的地方填上12,所以y=[5,7,9,12],底层数组变为[5,7,9,12],导致x=[5,7,9,12]。
最后x改变的原因是,x和y底层数组一样都是指向s的底层数组,并且x没有发生扩容,没有发生底层数组的迁移。
*/

func main() {
	s := []int{5}
	s = append(s, 7)
	s = append(s, 9)
	x := append(s, 11)
	y := append(s, 12)
	fmt.Println(s, x, y)
}
