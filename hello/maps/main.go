package main

import "fmt"

func main() {

	khalil := make(map[string]int)

	// ///save data in map///
	khalil["xyz"] = 10
	khalil["bca"] = 20

	// /read data from map/////
	val, ok := khalil["xyz"]
	fmt.Println("value", val, "ok", ok)

}
