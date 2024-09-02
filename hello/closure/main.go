package main

import "fmt"

func add() func() int {
	i := 5
	return func() int {
		i++
		return i
	}
}

func main() {
	result := add()
	fmt.Println("result", result())

}
