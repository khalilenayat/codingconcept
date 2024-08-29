package main

import "fmt"

type UserA struct {
	name   string
	adress string
	cont   int
}
type UserB struct {
	name   string
	adress string
	cont   int
}

func main() {
	var user UserA
	user.adduser(5)
	user.adress = "ctc"
	var khalil UserB
	khalil.adduser(6)
	khalil.adress = "slr"

}

func (user UserA) adduser(a int) {
	fmt.Println("user ko print krha", user)

}
func (user UserB) adduser(a int) {
	fmt.Println("user ko print krha", user)
}
