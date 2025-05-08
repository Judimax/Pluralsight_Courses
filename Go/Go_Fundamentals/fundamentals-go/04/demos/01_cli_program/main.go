package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("What would you like me to scream?")
	in := bufio.NewReader(os.Stdin) // os.Stdin is raw keyboard input
	s, _ := in.ReadString('\n') // \n is the delimiter
	s = strings.TrimSpace(s)  // remove \n
	s = strings.ToUpper(s) 
	fmt.Println(s + "!")
}
