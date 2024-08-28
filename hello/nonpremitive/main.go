package main

import "fmt"

type student struct {
	name  string
	roll  int
	class int
}

func main() {
	var vishal student
	vishal.class = 9
	vishal.roll = 9
	vishal.name = "khalil"
	fmt.Println("struct ka print ", vishal)
	var khalil student
	khalil.roll = 10
	khalil.name = "sam"
	fmt.Print("print khalil st", khalil)

	val := 132
	val2 := "hsjkdk"

	var example interface{}
	example = val
	fmt.Print("interfacae ", example)

	var secexp interface{}
	secexp = val2
	fmt.Print("secexample", secexp)

	var varadresss *student
	varadresss = &vishal
	fmt.Print("vishaladresss", varadresss)

}
