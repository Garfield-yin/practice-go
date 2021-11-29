package main

func foo() func() int {
	val := 1

	return func() int {
		val += 1
		return val
	}
}

func main() {
	foo()
}
