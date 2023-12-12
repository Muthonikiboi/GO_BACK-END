package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for our pizza")

	//store reader into a variable 
	//no try catch in go 
	//comma ok // err ok

	input,_:=reader.ReadString('\n')
	fmt.Println("thanks for the ratings ->",input)
	fmt.Printf(`type of rating is %T`,input)

	 
}