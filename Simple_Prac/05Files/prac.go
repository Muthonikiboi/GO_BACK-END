package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Code practice on File handling")

	//file content 
	fileContent:="This is my file handling practice lets see to what level have I understood"
	//create a file
	file , err :=os.Create("./myfile.txt",)

	if err != nil{
		panic(err)
	}

	//write into file
	length, err :=io.WriteString(file,fileContent)

	if err != nil{
		panic(err)
	}

	//output length of file 
	fmt.Println("The length of the file is",length)

	//close file //defer to make it be the last to output
	defer file.Close()
}

//we create a function to read file (filename in form of string)
func readFile(filename string){
	byteCode, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("the text inside file",string(byteCode))
}