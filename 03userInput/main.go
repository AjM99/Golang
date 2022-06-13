package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a := "Welcome"
	fmt.Println(a)

	b := bufio.NewReader(os.Stdin)
	fmt.Println("Enter something")

	// comma ok, comma err/_

	c, _ := b.ReadString('\n') // ,_ is for backup if error occurs
	fmt.Println("Thanks", c)

}
