package main

import "fmt"

/**
slice即是对数组的包装  动态数组
*/
func main() {
	//第一个方法是使用内建的函数 make。当我们使用 make 创建时，一个选项是可以指定 slice 的长度：
	slice1 := make([]string, 5)
	fmt.Println("slice1", slice1)
	//如果只指定了长度，那么容量默认等于长度。我们可以分别指定长度和容量：
	//不允许创建长度大于容量的 slice：
	slice2 := make([]int, 3, 5)
	fmt.Println("slice2", slice2)
	//字面量创建
	// 创建一个字符串 slice
	// 长度和容量都是 5
	slice3 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println("slice3", slice3)
	//在使用 slice 字面量创建 slice 时有一种方法可以初始化长度和容量，那就是初始化索引。下面是个例子：
	slice4 := []string{99: ""}
	fmt.Println("slice4", slice4)
	//创建一个 nil slice
	var slice5 []int
	fmt.Println("slice5", slice5 == nil)
	//创建 empty slice
	silce6 := make([]int, 0)
	slice7 := []int{}
	fmt.Println("empty slice", silce6 == nil, slice7)

	//Slice引用传递发生“意外”

	///上边我们一直在说，Slice是引用类型，指向的都是内存中的同一块内存，不过在实际应用中，有的时候却会发生“意外”，这种情况只有在像切片append元素的时候出现，Slice的处理机制是这样的，当Slice的容量还有空闲的时候，append进来的元素会直接使用空闲的容量空间，但是一旦append进来的元素个数超过了原来指定容量值的时候，内存管理器就是重新开辟一个更大的内存空间，用于存储多出来的元素，并且会将原来的元素复制一份，放到这块新开辟的内存空间。

	a := []int{1, 2, 3, 4}
	sa := a[1:3]
	fmt.Printf("%p\n", sa)
	sa = append(sa, 11, 22, 33)
	fmt.Printf("%p\n", sa)
	//可以看到执行了append操作后，内存地址发生了变化，说明已经不是引用传递。
}
