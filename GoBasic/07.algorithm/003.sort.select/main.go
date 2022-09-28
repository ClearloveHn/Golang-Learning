package main

import (
	"fmt"
	"math/rand"
)

func selectSort(arr *[]int) {
	length := len(*arr)
	for i := 0; i < length-1; i++ {
		min := i                          //定义最小元素的坐标为0
		for j := i + 1; j < length; j++ { //遍历数组和最小元素进行对比
			if (*arr)[min] > (*arr)[j] {
				min = j
			}
		}
		(*arr)[i], (*arr)[min] = (*arr)[min], (*arr)[i]
	}
}

func main() {
	arrSize := 10
	arr := []int{}
	for i := 0; i < arrSize; i++ {
		arr = append(arr, rand.Intn(50))
	}
	selectSort(&arr)
	fmt.Println(arr)

}
