package main

import "fmt"

func main() {
	a, b := 5, 2

	fmt.Println(a + b) //
	fmt.Println(a-b) // 3
	fmt.Println(a*b) // 10
	fmt.Println(a/b) //2
	fmt.Println(a%b) // 1

	fmt.Println(float32(a)/float32(b)) // 2/5

	fmt.Println(a == b) //false
	fmt.Println(a < b) // false
	fmt.Println(a > b) //false
}
