package main

import (
	"fmt"
	"time"
)

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %s\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go showMsg("Golang")
	time.Sleep(time.Millisecond * 2000)
}
