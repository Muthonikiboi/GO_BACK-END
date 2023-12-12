package main


func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080",srv)
}
