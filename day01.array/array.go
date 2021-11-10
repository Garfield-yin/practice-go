package main

import "fmt"

// 数组，值类型
//数组的长度是其类型的一部分。比如，[1]string和[2]string就是两个不同的数组类型。
func main() {
	arr := [5]int{1,2,3,4,5}
	for idx,val := range arr{
		fmt.Printf("array idx:%d, value of idx: %d\n", idx,val)
	}
	
	arrLen := len(arr)
	fmt.Printf("length of arr:%d\n", arrLen)
	// 数字的容量等于长度（能存放元素的数量）
	arrCap := cap(arr)
	fmt.Printf("cap of arr:%d\n", arrCap)
}