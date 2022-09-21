package main

import "fmt"

/*
值接收者和指针接收者的区别是:指针接收者可以改变结构体内部的数据
*/

//定义一个对象
type treeNode struct {
	value       int       //成员变量
	left, right *treeNode //成员变量
}

//成员方法:值接收者
func (t treeNode) print() {
	fmt.Println(t.value)
}

//指针接收者
func (t *treeNode) setValue(value int) {
	t.value = value
}

func main() {
	//实例化对象
	root := treeNode{value: 3,
		left:  &treeNode{right: &treeNode{value: 2}},
		right: &treeNode{5, &treeNode{}, nil},
	}
	//用对象.来调用方法
	root.print()
	root.setValue(10)
	root.print()

}
