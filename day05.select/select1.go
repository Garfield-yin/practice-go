package main

import "fmt"


var channels = [3]chan int{
		nil,
		make(chan int),
		nil,
	}
var indexs = [3]int{0,1,2}

func main(){
	// 每个case 都会被求值
	// 从左到右，从上到下
	
	select {
	case getChan(0)<- getIndex(0):
		fmt.Println("The first candidate case is selected.")
		case getChan(1)<- getIndex(1):
		fmt.Println("The second candidate case is selected.")
		case getChan(2)<- getIndex(2):
		fmt.Println("The third candidate case is selected.")
		default:
		fmt.Println("No candidate case is selected!")

	}


}
func getChan(idx int) chan<- int{
	fmt.Printf("channels[%d]\n", idx)
	return channels[idx]
}

func getIndex(idx int) int {
	fmt.Printf("indexs[%d]\n", idx)
	return indexs[idx]
}