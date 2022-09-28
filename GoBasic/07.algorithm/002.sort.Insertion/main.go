package main

import (
	"fmt"
	"math/rand"
)

func insertSort(arr *[]int) {
	n := len(*arr)
	for i := 1; i < n; i++ {
		value := (*arr)[i]                //要插入的数据
		j := i - 1                        //已排好序的元素下标
		for j >= 0 && value < (*arr)[j] { //插入的数据和已排好序的元素依次对比
			(*arr)[j+1] = (*arr)[j]
			j--
		}
		(*arr)[j+1] = value
	}
}

func main() {
	arrSize := 10
	arr := []int{}
	for i := 0; i < arrSize; i++ {
		arr = append(arr, rand.Intn(50))
	}
	insertSort(&arr)
	fmt.Println(arr)
}
