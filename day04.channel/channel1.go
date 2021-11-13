package main

import "fmt"

func main(){
	ch1 := make(chan int,1)
	ch1<-1
	//ch1<-2 // 阻塞，通道已满
	ch2 := make(chan int,1)
	// 通道满时，长度等于容量
	fmt.Printf("length:%d,cap:%d\n",len(ch2),cap(ch2))

	//val,ok := <-ch2 //如果通道已空，那么对它的所有接收操作都会被阻塞
	//fmt.Printf("val:%d,is ok:%v",val,ok)
	ch2 <- 1


	var ch3 chan int
	//ch3 <- 1 // 阻塞
	// <-ch3 // 阻塞
	//通道的值为nil，读写都会造成永久的阻塞！
	_=ch3

	// 非缓冲通道
	ch4 := make(chan int)
	//ch4 <-1 阻塞
	// <-ch4 阻塞
  _ = ch4
}