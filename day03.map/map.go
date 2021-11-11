package main

import "fmt"

// map
func main(){
	// 引用类型 m=nil
	var m map[string]int 
	val,ok := m["key_1"]
	fmt.Printf("map val:%d,key:%s,is ok:%v", val,"key_1", ok)

	delete(m,"key_1")
	// 对一个nil map 写会引发 panic,读或者删除不会出错
	//m["key_1"] = 1

	// 并发不安全
}