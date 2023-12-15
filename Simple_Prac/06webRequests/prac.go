package main

import (
	"fmt"
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
}