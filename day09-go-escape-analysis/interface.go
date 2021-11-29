package main

type foo interface {
	fooFunc()
}
type foo1 struct{}

func (f1 foo1) fooFunc() {}

func main() {
	var f foo
	f = foo1{}
	f.fooFunc()
}
