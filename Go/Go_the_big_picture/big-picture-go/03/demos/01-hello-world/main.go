package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, World")

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Web services are easy with go")
	})

	http.ListenAndServe(":3000",nil)
}
