package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex
var wg sync.WaitGroup
var m = 100

func add() {
	defer wg.Done()
	defer lock.Unlock()
	lock.Lock()
	m += 1
	fmt.Printf("m++:%v\n", m)

}

func sub() {
	defer wg.Done()
	defer lock.Unlock()
	lock.Lock()
	m -= 1
	fmt.Printf("m--:%v\n", m)

}

func main() {
	for i := 0; i < 100; i++ {
		go add()
		wg.Add(1)
		go sub()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("end:", m)
}
