package main

import (
	"fmt"
	"sync"
)

//sync.WaitGroup 是保证Goroutine之间同步的一个工具

var wg sync.WaitGroup

func hello(msg string) {
	defer wg.Done()
	fmt.Printf("hello %v\n", msg)

}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 计数器,每启动一个goroutine就+1
		go hello("golang")
		wg.Wait()
	}

}
