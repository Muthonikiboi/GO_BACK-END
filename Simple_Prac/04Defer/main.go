package main

import "fmt"

func main(){
	fmt.Println("hello")
	defer fmt.Println("world")

	//when we interchange the output is world Hello
	fmt.Println("world")
	fmt.Println("hello")

	//assume me write defer as the first the output is Hello World
	defer fmt.Println("world") 
	fmt.Println("hello")

	//occurs in order it was deferred(last in first out)
	defer fmt.Println("world")
	defer fmt.Println("my")
	defer fmt.Println("to")   
	fmt.Println("hello")//output will be  ("Hello to my world")
	deferring()

	//with this function called here the output will be 
	//hello,4,3,2,1,0, to, my , world
}


//having an outside function
func deferring(){
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}
}
