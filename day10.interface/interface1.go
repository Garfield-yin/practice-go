package main

import "fmt"

type Pet interface{
	SetName(name string)
	Name() string
	Category() string
}

type Pig struct{
	name string
}

func(pig *Pig) SetName(name string){
	pig.name = name
}

func(pig Pig)Name() string{
	return pig.name
}

func (pig Pig)Category()string{
	return "pig"
}

func main(){
	pig :=Pig{name:"xiaopp"}
	// &pig 的结果值就是pet 的动态值， *pig 的类型就是pet 的动态类型
	// 此时发生内存逃逸
	var pet Pet = &pig // 指针方法才包含了所有的接口方法
	fmt.Println(pet.Name())

	fmt.Println("-----华丽的分割线---------------")
	// pig 是否实现了 Pet 接口
	_,ok := interface{}(pig).(Pet)
	fmt.Printf("pig implements interface Pet: %v\n", ok)

		// &pig 是否实现了 Pet 接口
	_,ok = interface{}(&pig).(Pet)
	fmt.Printf("&pig implements interface Pet: %v\n", ok)

}

