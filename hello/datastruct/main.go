package main

import "fmt"

func main() {
	var list [20]int
	list[0] = 10
	list[7] = 10

	fmt.Println("list of array", list)

	/////////slice/////////

	var lists []int
	lists = append(lists, 20)
	fmt.Println("lists", lists[0])
}
