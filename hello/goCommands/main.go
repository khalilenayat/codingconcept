package main

import (
	"fmt"
)

func main() {
	fmt.Println("testing go basic commands")
	mobno := 7377869902
	khalil := "khalil"
	fmt.Println("my name is khalil", khalil, "mobile no is ", mobno)
	var addno *string
	addno = &khalil
	fmt.Println("adress", addno)

}
