package main

func main() {
	ch1 := make(chan *int, 1)
	y := 5
	py := &y
	ch1 <- py
}
