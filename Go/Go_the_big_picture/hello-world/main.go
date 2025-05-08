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

	http.HandleFunc("/home",func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w,r, "../big-picture-go/03/demos/02-web-service/home.html")
	})

	http.ListenAndServe(":3000",nil)
}
