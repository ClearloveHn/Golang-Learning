package main

import "fmt"

type Person interface {
	job()
	growUp()
}

type student struct {
	age int
}

func (s student) job() {
	fmt.Println("I am a student")
}
func (s *student) growUp() {
	s.age += 1
}

type programmer struct {
	age int
}

func (p programmer) job() {
	fmt.Println("I am a programmer")
}

func (p *programmer) growUp() {
	p.age += 10
}

func whatJob(p Person) {
	p.job()
}

func growUp(p Person) {
	p.growUp()
}

func main() {
	a := student{
		age: 18,
	}
	whatJob(&a)
	growUp(&a)
	fmt.Println(a)

	var b Person = &programmer{
		age: 25,
	}
	whatJob(b)
	growUp(b)
	fmt.Println(b)
}
