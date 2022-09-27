package main

import (
	"fmt"
	"sync"
)

type stack struct {
	sync.Mutex
	data []interface{}
}

func (s *stack) push(data interface{}) {
	s.Lock()
	defer s.Unlock()
	s.data = append([]interface{}{data}, s.data...)
}

func (s *stack) pop() (interface{}, bool) {
	s.Lock()
	defer s.Unlock()

	if len(s.data) > 0 {
		o := s.data[0]
		s.data = s.data[1:]
		return o, true
	}
	return nil, false
}

func main() {
	s := &stack{}
	s.push(111)
	s.push(222)
	s.push(333)

	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
}
