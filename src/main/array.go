package main

import "fmt"

func main() {
	/**
	go函数中的数组传递是值传递带价很高
	*/

	var array1 [5]int
	fmt.Println("array1", array1)
	array2 := [5]int{555, 858, 5858, 5565, 777}
	fmt.Println("array2", array2)
	// 声明一个长度为5的整数数组
	// 为索引为1和2的位置指定元素初始化
	// 剩余元素为0值
	array3 := [5]int{1: 77, 2: 777}
	fmt.Println("array3", array3)
	//指针数组
	array4 := [5]*int{0: new(int), 1: new(int)}
	// 为索引为0和1的元素赋值
	*array4[0] = 7
	*array4[1] = 77
	fmt.Println("array4", *array4[0], *array4[1])
	//多维数组
	array5 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println("array5", array5)
}
