package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for our pizza")

	// store reader into a variable
	// no try-catch in Go
	// comma ok // err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for the ratings ->", input)
	fmt.Printf("type of rating is %T\n", input)

	// conversions assuming we want to add one to each of the ratings

	numRatings, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Add 1 to your ratings", numRatings+1)
	}
}
