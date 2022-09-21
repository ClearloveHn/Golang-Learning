package main

import "fmt"

/**
1. 在调用方法中,值类型既可以调用值接收者方法,也可以调用指针接收者方法。指针类型同样如此。
2. 在调用指针接收者的方法时,无论什么类型调用,都会改变struct内部的数据。
*/

type Person struct {
	age int
}

func (p Person) howOld() int {
	return p.age
}

func (p *Person) growUp() {
	p.age += 1
}

func main() {
	//lhn是值类型
	lhn := Person{
		age: 18,
	}
	//值类型调用值接收者方法
	fmt.Println(lhn.howOld())
	//值类型调用指针接收者方法
	lhn.growUp()
	fmt.Println(lhn.howOld()) //struct内部值发生改变

	//xyr是指针类型
	xyr := &Person{
		age: 21,
	}
	//指针类型调用值接收者方法
	fmt.Println(xyr.howOld())
	//指针类型调用指针接收者方法
	xyr.growUp()
	fmt.Println(xyr.howOld())
}
