package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Rate")

	a := bufio.NewReader(os.Stdin)

	input, _ := a.ReadString('\n')
	fmt.Println("thanks", input)

	//convert string to int  and using comma error
	c, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Added 1", c+1)
	}
}
