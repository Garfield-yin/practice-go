package main

import (
	"errors"
	"fmt"
)


type Operator func(x,y int)(int)
type calculateFunc func(x int, y int) (int, error)


func genCalculator(op Operator) calculateFunc{
	return func(x,y int)(int, error){
		if op ==nil{
			return 0, errors.New("invalid operation")
		}
		return op(x,y),nil
	}
}

func main() {
	x,y :=2,3
	op := func(x,y int)int{
		return x+y
	}
	add := genCalculator(op)
	result,err := add(x,y)
	fmt.Printf("result:%d,err:%v\n",result,err)
}