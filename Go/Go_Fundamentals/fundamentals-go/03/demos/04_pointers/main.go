package main

import "fmt"

func main() {
	s := "Hello, world!"

	p := &s

	fmt.Println(p) // you get a memory address
	fmt.Println(*p)

	*p = "Hello, Gophers!" // defereince now equals hello gophers

	fmt.Println(s, *p)

	p = new(string)

	fmt.Println(p, *p) //can reference that memory indrectly now, value of p is an empty strinng

}
