package main

import "fmt"

type Car struct {
	Model string
	Version string
}

type Deutch Car

func main()  {
	var audi Deutch = Deutch{"Deutch", "Q3"}
	fmt.Println(audi)
}