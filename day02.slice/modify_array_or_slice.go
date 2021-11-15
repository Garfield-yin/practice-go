package main

import "fmt"

func main(){
	// 示例1。
	array1 := [3]string{"a", "b", "c"}
	fmt.Printf("The array: %v\n", array1)
	array2 := modifyArray(array1)
	fmt.Printf("The modified array: %v\n", array2)
	fmt.Printf("The original array: %v\n", array1)
	fmt.Println("------华丽的分割线1-----")

	slice1 := []string{"x", "y", "z"}
	fmt.Printf("The slice: %v\n", slice1)
	slice2 := modifySlice(slice1)
	fmt.Printf("The modified slice: %v\n", slice2)
	fmt.Printf("The original slice: %v\n", slice1)
	fmt.Println("------华丽的分割线2-----")

	complexArray1 := [3][]string{
		{"d", "e", "f"},
		{"g", "h", "i"},
		{"j", "k", "l"},
	}
	fmt.Printf("The complex array: %v\n", complexArray1)
	complexArray2 := modifyComplexArray(complexArray1)
	fmt.Printf("The modified complex array: %v\n", complexArray2)
	fmt.Printf("The original complex array: %v\n", complexArray1)

	fmt.Println("------华丽的分割线3-----")
	complexArray3 := modifyComplexSlice(complexArray1)
	fmt.Printf("The modified complex array: %v\n", complexArray3)
	fmt.Printf("The original complex array: %v\n", complexArray1)
	/*
	只改数组索引下标下的整个切片内容，
	因为会重新生成切片的底层数组，
	不会影响到原值。
	修改数组索引下标的切片的索引下标的单一字符串元素，
	 会影响到原值。
	*/

}

func modifyArray(arr [3]string) [3]string{
	arr[1] = "x"
	return arr;
}

func modifySlice(slice []string) []string{
	slice[1] = "x"
	return slice
}

func modifyComplexArray(arr [3][]string)[3][]string{
	arr[1][1] = "s"
	return arr
}

func modifyComplexSlice(arr [3][]string)[3][]string{
	arr[1] = []string{"x","y","z"}
	return arr
}