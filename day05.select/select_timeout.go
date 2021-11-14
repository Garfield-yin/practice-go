package main

import (
	"fmt"
	"time"
)


var done = make(chan struct{},1)
func main(){

	go func(){
		doSomething()
	}()
	select {
	case <- done:
		fmt.Println("do some thing done")
		return;

	case <- time.After(5*time.Second):
		fmt.Println("timeout")
		return
	}

}

func doSomething()([]byte,error){
	time.Sleep(3*time.Second)
	done <- struct{}{}
	return nil,nil
}