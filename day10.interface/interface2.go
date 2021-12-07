package main

import "fmt"

/*
iface的实例会包含两个指针，一个是指向类型信息的指针，另一个是指向动态值的指针。
这里的类型信息是由另一个专用数据结构的实例承载的，
其中包含了动态值的类型，以及使它实现了接口的方法和调用它们的途径，
*/
type Pet interface {
  Name() string
  Category() string
}

type Dog struct{
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main(){
	dog := Dog{name:"little pig"}
	// %q    该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	fmt.Printf("The dog's name is %s\n", dog.Name()) // The dog's name is little pig

	/*如果我们使用一个变量给另外一个变量赋值，那么真正赋给后者的，并不是前者持有的那个值，而是该值的一个副本。*/
	var pet Pet = dog
	dog.SetName("big pig pig")
	fmt.Printf("The dog's name is %s\n", dog.Name()) //The dog's name is big pig pig
	fmt.Printf("The pet's name is %s\n", pet.Name())//The dog's name is little pig

	//接口类型本身是无法被值化的。在我们赋予它实际的值之前，它的值一定会是nil

	dog1 := Dog{"little pig"}
	fmt.Printf("The name of first dog is %s.\n", dog1.Name())
	dog2 := dog1
	fmt.Printf("The name of second dog is %q.\n", dog2.Name())
	dog1.name = "monster"
	fmt.Printf("The name of first dog is %q.\n", dog1.Name())
	fmt.Printf("The name of second dog is %q.\n", dog2.Name())
	fmt.Println()

	dog = Dog{"little pig"}
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	pet = &dog
	dog.SetName("big pig")
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
}