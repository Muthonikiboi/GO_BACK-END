package main

import (
	"log"
	"net/http"
)
	

func Hello(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte("<h1 style='color:blue'>Hello There This Is My First Web Server</h1>"))
}
func main() {
	http.HandleFunc("/hello",Hello) 
	//initialize http server
	//log.fatal helps see where the issue is
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
	
}