package main

import "fmt"

func main()  {
	var from string
	var to string

	from = "wiza@notifications.fr"
	to = "jacquirhatim@gmail.com"
	// create pointer
	toPointer := &to
	fromPointer := &from

	fmt.Println(from, to, *toPointer, *fromPointer)
}