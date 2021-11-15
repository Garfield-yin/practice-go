package main

import "fmt"

// 函数签名
type Printer func(str string)(n int, err error)

//只要两个函数的参数列表和结果列表中的元素顺序及其类型是一致的，我们就可以说它们是一样的函数，
//或者说是实现了同一个函数类型的函数。
// 函数类型属于引用类型，它的值可以为nil，而这种类型的零值恰恰就是nil。
func PrinterToStd(str string)(n int, err error){
	return fmt.Println(str)
}

func main(){
	var printer Printer
	printer = PrinterToStd;
	printer("hello world")
}