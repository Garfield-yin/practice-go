package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
select语句只能与通道联用，
它一般由若干个分支组成。
每次执行这种语句的时候，
一般只有一个分支中的代码会被运行。*/

// select 阻塞试
func main(){
	intChannels := [3]chan int{
		make(chan int ,1),
		make(chan int ,1),
		make(chan int ,1),
	}
	idx := rand.Intn(3)
	fmt.Println("The idx:", idx)
	intChannels[idx] <- idx

	// 1. 该select 不会阻塞,有 default
	// 2. 如果没有default,且没有表达式满足，将阻塞，直到有满足条件的case
	select {
	case <- intChannels[0]:
		fmt.Println("The first candidate case is selected.")
	case <- intChannels[1]:
		fmt.Println("The second candidate case is selected.")
	case <- intChannels[2]:
		fmt.Println("The third candidate case is selected.")
	default:
			fmt.Println("The default candidate case is selected.")
	}

	ch1 := make(chan int,1)
	time.AfterFunc(2*time.Second, func(){
		close(ch1)
	})
	select {
	case _,ok := <- ch1:
		if !ok {
			fmt.Println("The candidate case is closed.") 
			break
	}
	}

}