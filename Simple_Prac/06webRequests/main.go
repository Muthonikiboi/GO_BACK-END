package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//make url global
const url = "https://lco.dev"

func main(){

	//make simple request from net package
	response,err := http.Get(url)

	if err != nil{
		panic (err)
	}

	//check type of response
	fmt.Printf("The type of response is : %T \n",response)

	//close the public body of the response ie CONNECTION
	defer response.Body.Close()

	//read the response
	databyte, err := ioutil.ReadAll(response.Body)
	
	if err != nil{
		panic(err)
	}

	//check the response in string format
	content := string(databyte)
	fmt.Println(content)
}