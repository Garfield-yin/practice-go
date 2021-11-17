package main

import "fmt"

type AnimalCategory struct{
	kingdom string //界
	phylum string //门
	class string // 纲
	order string // 目
	family string // 科
	genus string // 属
	species string // 种
}


type Animal struct {
  scientificName string // 学名
  AnimalCategory    // 动物基本分类
}

/*
方法隶属的类型其实并不局限于结构体类型，但必须是某个自定义的数据类型，并且不能是任何接口类型
*/
func(ac AnimalCategory) String() string {
	fmt.Println("call AnimalCategory String")
	 return fmt.Sprintf("%s%s%s%s%s%s%s",    ac.kingdom, ac.phylum, ac.class, ac.order,    ac.family, ac.genus, ac.species)
}

func (a Animal) Category() string {
  return a.AnimalCategory.String()
}

func(a Animal) String() string {
	fmt.Println("call Animal String")
	 return fmt.Sprintf("%s%s%s%s%s%s%s",    a.kingdom, a.phylum, a.class, a.order,    a.family, a.genus, a.species)
}

func (a Animal) String1() string {
  return fmt.Sprintf("%s (category: %s)",
    a.scientificName, a.AnimalCategory)
}

type Cat struct {
  name string
  Animal
}
// 优先判断cat 是否有 string 方法，有的话屏蔽 Animal 和 AnimalCategory String
func (cat Cat) String() string {
 	fmt.Println("call Cat String")
  return fmt.Sprintf("%s (category: %s, name: %q)",
    cat.scientificName, cat.Animal.AnimalCategory, cat.name)
}
// 值方法的接收者是该方法所属的那个类型值的一个副本。我们在该方法内对该副本的修改一般都不会体现在原值上，除非这个类型本身是某个引用类型（比如切片或字典）的别名类型。

func (cat *Cat) SetName(name string) {
  cat.name = name
}


func main() {
	category := AnimalCategory{species: "cat"}
	fmt.Println(category.String())
	animal := Animal{AnimalCategory: category,scientificName:"animal"}
	// 如果定义了 animal.String 则 AnimalCategory.String 会被屏蔽
	fmt.Printf("The animal: %s\n", animal)

	cat := Cat{name: "Garfield", Animal: animal}
	fmt.Printf("The cat: %s\n", cat)
	cat.SetName("miaomaio") // Go 会转义为 &cat.SetName
	fmt.Printf("The cat: %s\n", cat)

}