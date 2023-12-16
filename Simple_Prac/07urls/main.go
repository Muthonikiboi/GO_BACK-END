package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://lco.dev:3000"

func main() {
	fmt.Println("URLs in Golang")
	fmt.Println(myURL)

	// Parse the URL to get the result
	result, err := url.Parse(myURL)
	if err != nil {
		panic(err)
	}

	// Print various components of the parsed URL
	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Port:", result.Port())
	fmt.Println("RawQuery:", result.RawQuery)

	//shows working with query parameters
	qparams := result.Query()
	fmt.Printf("The types of query params are: %T",qparams)

	fmt.Println(qparams["Port"])

	for _ , val := range qparams{
		fmt.Println("Param is:",val)
	}
	//creates a new url struct with specified components and then converts it to string 
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:"lco.dev",
		Path:"/tutcss",
		RawPath:"user=hitesh",//allows specifying  raw path including the query string
	}

	anotherURL := partsOfUrl.String()//converts the components to strings
	fmt.Println(anotherURL)

}
