package main

import "fmt"

func main() {
	var emptyInterface interface{} = "Hello"
	fmt.Println(emptyInterface)

	// assertion de type
	var otherInterface, ok = emptyInterface.(string)

	fmt.Println(otherInterface, ok)
}
