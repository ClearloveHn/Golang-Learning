package main

import (
	"fmt"
	"time"
)

/*
1.select是一个控制结构,用于实现异步IO。
2.select里面的case语句必须是channel操作。
3.当所有的case都准备就绪的时候,会随机选择一个channel运行。
4.如果没有可运行的case语句,若有default,走default,无default,产生阻塞。
*/

var chanInt = make(chan int)
var chanStr = make(chan string)

func main() {
	go func() {
		defer close(chanInt)
		defer close(chanStr)
		chanInt <- 100
		chanStr <- "hello"
	}()

	for {
		select {
		case r := <-chanInt:
			fmt.Printf("chanInt%v\n", r)
		case r := <-chanStr:
			fmt.Printf("chanStr%v\n", r)
		default:
			fmt.Println("default......")
		}
		time.Sleep(time.Second)
	}

}
