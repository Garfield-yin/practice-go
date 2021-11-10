package main

import "fmt"

//https://halfrost.com/go_slice/
func main(){
	s1 := make([]int,5,6);
	fmt.Printf("The address of s1: %p\n", &s1)
	fmt.Printf("The length of s1: %d\n", len(s1)) // 5
	fmt.Printf("The capacity of s1: %d\n", cap(s1)) // 6
	fmt.Printf("The value of s1: %d\n", s1)//[0 0 0 0 0]
	

	fmt.Println("华丽的分割线-------------------")
	// 扩容后的容量
	s2 := append(s1, 1,2)
	fmt.Printf("The address of s2: %p\n", &s2)
	fmt.Printf("The length of s2: %d\n", len(s2)) // 7
	fmt.Printf("The capacity of s2: %d\n", cap(s2)) // 12
	fmt.Printf("The value of s2: %d\n", s2)//[0 0 0 0 0 1 2]
	
	
	/*一旦一个切片无法容纳更多的元素，
	Go 语言就会想办法扩容。生成一个容量更大的切片，
	然后将把原有的元素和新元素一并拷贝到新切片中。
	在一般的情况下，你可以简单地认为新切片的容量（以下简称新容量）
	将会是原切片容量（以下简称原容量）的 2 倍。
	*/

}