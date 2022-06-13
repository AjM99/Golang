package main

import "fmt"

func main() {

	// array without numver in box is slice
	var a = []string{"1", "2", "3"}
	fmt.Println(a)

	a = append(a, "4")
	fmt.Println(a)

	a = append(a[1:3]) //will print postion 1 to 2  not 3
	fmt.Println(a)

	score := make([]int, 4)

	score[0] = 0
	score[1] = 1
	score[2] = 2
	score[3] = 3

	score = append(score, 9, 5, 6) // even though limit is 4 for slice it appends, frist 4 are default /fixed rest are added.

	fmt.Println(score)
}
