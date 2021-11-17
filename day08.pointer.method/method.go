package main

import "fmt"

type Cat struct {
	name           string // 名字。
	scientificName string // 学名。
	category       string // 分类。
}

func New(name, scientificName, category string) Cat {
	return Cat{
		name:           name,
		scientificName: scientificName,
		category:       category,
	}
}

func (cat *Cat) SetName(name string) {
	cat.name = name
}

func (cat Cat) SetNameOfCopy(name string) {
	cat.name = name
}

func (cat Cat) Name() string {
	return cat.name
}

func (cat Cat) ScientificName() string {
	return cat.scientificName
}

func (cat Cat) Category() string {
	return cat.category
}
func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.category, cat.name)
}

func main() {
	cat := New("小火锅", "美短", "cat")
	cat.SetName("garfield")
	fmt.Printf("The cat: %s\n", cat)

	cat.SetNameOfCopy("笑哈哈")
	fmt.Printf("The cat: %s\n", cat) // 不会改变原值

	type Pet interface {
		SetName(name string) //指针方法
		Name() string
		Category() string
		ScientificName() string
	}

	
	_, ok := interface{}(cat).(Pet) // SetName 没有被实现
	fmt.Printf("Cat implements interface Pet: %v\n", ok)

	_, ok = interface{}(&cat).(Pet) // 指针方法包含了类型的值方法
	fmt.Printf("*Cat implements interface Pet: %v\n", ok)
}