package main

import "fmt"

func main() {

	// declare size
	var a [5]string
	a[0] = "1"
	a[1] = "2"
	//there will be a blank space for 2 in output
	a[3] = "4"
	fmt.Println("array is :", a)
	fmt.Println("array size :", len(a))

	var b = [5]string{"a", "b", "c"}
	fmt.Println("array 2:", b)
	fmt.Println("array 2:", len(b))
}
