package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)


func main() {
	content :="My name is joy"

	file,err:= os.Create("./myDetails.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNil(err)
	//write a file 
	length,err := io.WriteString(file,content)

	// if err != nil{
	// 	panic(err)
	// }
	checkNil(err) 
	fmt.Println("The length is: ",length)
	defer file.Close()
}

//reading the file
func readFike(filname string){
	databyte, err :=ioutil.ReadFile(filname)

	// if err != nil {
	// 	panic(err)
	// }
	checkNil(err)
	fmt.Println("The text inside thr file",string(databyte))

}


//create a function the replaces the if err!= nil
func checkNil(err error){
	if err != nil {
		panic(err)
	}// we replace this code similar to this with ( >>>checkNil(err)<<<)
}
