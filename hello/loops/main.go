package main

import "fmt"

func main() {
	for i := 20; i <= 100; i++ {
		fmt.Println("numbers", i)
	}

	////startpoint,endpoint,changepoint//////

	////range in array/////
	var list [8]int
	list[0] = 10
	list[7] = 10

	//fmt.Println("list of array", list)
	for _, j := range list {
		fmt.Println("value is", j)
	}

}
