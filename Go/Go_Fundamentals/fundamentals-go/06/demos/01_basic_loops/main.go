package main

import (
	"fmt"
)

func main() {
	i := 99

	for { //infinite loop
		fmt.Println(i)
		i += 1
		break
	}

	i = 5
	for i < 10 { //loop till
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
