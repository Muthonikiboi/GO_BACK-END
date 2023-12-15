package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


const url = "https://lco.dev"

func main(){
	response,err := http.Get(url)

	if err != nil{
		panic (err)
	}
	fmt.Printf("The type of response is : %T \n",response)
	defer response.Body.Close()

	//to read response we use ioutil.ReadAll(response.Body)

	databyte ,err := ioutil.ReadAll(response.Body)
	
	if  err != nil {
		panic(err)
	}

	getContent := string(databyte)
	fmt.Println(getContent)
}