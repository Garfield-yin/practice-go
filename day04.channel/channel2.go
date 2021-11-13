package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 单向通道，只能发，不能收
	ch1 := make(chan<- int, 1)
	ch1 <- 1
	//<-ch1 //invalid operation: <-ch1 (receive from send-only type chan<- int)

	// 只能收不能发
	ch2 := make(<-chan int, 1)
	//ch2<-1
	// <-ch2 阻塞
	_=ch2

	ch3 := make(chan int, 3)
	SendInt(ch3)

	ch4 := getChan()
	for ele := range ch4{
			fmt.Printf("The element in ch4: %v\n", ele)
	}

}

// 单向通道，约束函数行为
func SendInt(ch chan<- int){
	ch <- rand.Intn(1000)
}

func getChan() <-chan int{
	num := 3
	ch := make(chan int,num)
	for i:=0; i< num;i++ {
		ch<-i
	}
	close(ch)
	return ch
}