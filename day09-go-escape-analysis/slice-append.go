package main

func main() {
	s1 := make([]int, 10000, 10000) // slice-append.go:4:12: make([]int, 10000, 10000) escapes to heap
	_ = s1
}
