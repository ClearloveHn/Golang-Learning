package main

import "fmt"

/**
map的增删改查操作
*/

func main() {
	namesMap := make(map[string]string)
	//增
	namesMap["li"] = "李"
	namesMap["liu"] = "刘"
	fmt.Println("增", namesMap)
	//删
	delete(namesMap, "li")
	fmt.Println("删", namesMap)
	//改
	namesMap["liu"] = "留"
	fmt.Println("改", namesMap)
	//查
	for k, v := range namesMap {
		fmt.Println(k, v)
	}
}
