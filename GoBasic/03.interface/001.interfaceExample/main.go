package main

import "fmt"

/**
1. 如果实现了接收者是值类型的方法，会隐含地也实现了接收者是指针类型的方法,反之亦然。
2. 如果实现接口的方法一个是值接收者,一个是指针接收者,那么实例化的时候,如果实例化的类型为接口,那么必须用指针。
*/

type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p *Gopher) debug() {
	fmt.Printf("I am dubuging %s languahe\n", p.language)
}

func main() {
	//实例化struct类型
	var c = Gopher{
		language: "Golang",
	}
	c.code()
	c.debug()

	//实例化interface类型
	var b coder = &Gopher{"Java"}
	b.code()
	b.debug()
}
