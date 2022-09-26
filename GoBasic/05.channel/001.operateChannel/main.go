package main

import (
	"fmt"
	"math/rand"
	"time"
)

var values = make(chan int) //创建一个无缓冲Channel

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send: %v\n", value)
	time.Sleep(time.Second * 5)
	values <- value //向channel发送数据
}

func main() {

	defer close(values) //defer 关闭channel
	go send()
	fmt.Println("waiting")
	value := <-values //从channel接受值
	fmt.Printf("receive: %v\n", value)
	fmt.Printf("end")
}
