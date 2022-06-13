package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	//symbils for pointer * &

	a := 2
	var b = &a                 // & is refence
	fmt.Println("value :", *b) // * gives the value
	fmt.Println("value:", b)   //

	*b = *b * 2
	fmt.Println("new value:", a)

}
