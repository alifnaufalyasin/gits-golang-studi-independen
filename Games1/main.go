package main

import (
	"fmt"
)

func main() {
	// M Alif Naufal Yasin - Universitas Telkom
	var name, uni, age, hobby string
	fmt.Print("Write your name : ")
	fmt.Scanln(&name)
	fmt.Print("Write your university : ")
	fmt.Scanln(&uni)
	fmt.Print("Write your age : ")
	fmt.Scanln(&age)
	fmt.Print("Write your hobby : ")
	fmt.Scanln(&hobby)

	fmt.Printf("Hello my name is %s from %s my age %s and my hobby %s.", name, uni, age, hobby)

}
