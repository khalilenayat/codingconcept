package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var a int
	a = 16
	var b int16
	b = int16(a)
	fmt.Println(b)

	num := "12345"
	numint, err := strconv.Atoi(num)
	fmt.Println("numint =", numint, "err =", err)
	fmt.Println("now thwe time is", time.Now())
}
