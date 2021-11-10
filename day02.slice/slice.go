package main

import "fmt"

// slice(切片特点)，可变长,引用类型
// 在每个切片的底层数据结构中，一定会包含一个数组
//数组可以被叫做切片的底层数组，而切片也可以被看作是对数组的某个连续片段的引用。
func main(){
	s1 := make([]int, 5)// 如果不指明其容量，那么它就会和长度一致
	fmt.Printf("The length of s1: %d\n", len(s1)) // 5
	fmt.Printf("The capacity of s1: %d\n", cap(s1)) // 5
	fmt.Printf("The value of s1: %d\n", s1) //[0,0,0,0,0]
	s2 := make([]int, 5, 8) // 切片的容量实际上代表了它的底层数组的长度
	fmt.Printf("The length of s2: %d\n", len(s2)) // 5
	fmt.Printf("The capacity of s2: %d\n", cap(s2)) // 8
	fmt.Printf("The value of s2: %d\n", s2) //[0 0 0 0 0]

	s3 := []int{1,2,3,4,5,6}
	fmt.Printf("The length of s3: %d\n", len(s3)) // 6
	fmt.Printf("The capacity of s3: %d\n", cap(s3)) // 6
	fmt.Printf("The value of s3: %d\n", s3) //[1 2 3 4 5 6]
	s4 := s3[3:5]
	fmt.Printf("The length of s4: %d\n", len(s4)) // 2
	fmt.Printf("The capacity of s4: %d\n", cap(s4)) // 3
	fmt.Printf("The value of s4: %d\n", s4) //[4 5]
	// 将切片扩展到最大
	s5 := s3[0:cap(s3)]
	fmt.Printf("The length of s5: %d\n", len(s5)) // 6
	fmt.Printf("The capacity of s5: %d\n", cap(s5)) // 6
	fmt.Printf("The value of s5: %d\n", s5) //[1 2 3 4 5 6]
}