package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time")
	a := time.Now()
	fmt.Println(a)
	fmt.Println(a.Format("01-02-2006 15:04:05 Monday")) // you have to specify the std foramt which is this

}
