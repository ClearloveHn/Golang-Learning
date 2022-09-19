package main

import "fmt"

/**
1.使用make创建一个切片时,会分别传入,type,len和cap三个参数,可以只传入两个参数,这种情况下len==cap。
2.如果len不为0,使用append向slice追加函数时,要补0
3.补零的原因:因为make操作是完成了对切片底层数组的创建和初始化,0是默认值。(len为x就有x个默认值)
*/
func main() {
	slice := make([]int, 10, 10)
	slice = append(slice, 1)
	fmt.Println(slice)
}
