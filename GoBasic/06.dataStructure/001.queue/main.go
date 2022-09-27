package main

import (
	"fmt"
	"sync"
)

type queue struct {
	sync.Mutex
	data []interface{}
}

//入队列
func (q *queue) push(data interface{}) {
	q.Lock()
	defer q.Unlock()
	q.data = append(q.data, data)
}

//出队列
func (q *queue) pop() (interface{}, bool) {
	q.Lock()
	q.Unlock()
	if len(q.data) > 0 {
		o := q.data[0]
		q.data = q.data[1:] //留下标1后面的元素
		return o, true
	} else {
		return nil, false
	}
}

func main() {
	q := &queue{}
	q.push(111)
	q.push(222)
	q.push(333)

	fmt.Println(q.pop())
	fmt.Println(q.pop())
	fmt.Println(q.pop())
	fmt.Println(q.pop())

}
