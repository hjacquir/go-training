package main

import "fmt"

func main()  {
	days := []string{"m", "t", "w"}

	for _, v := range days {
		fmt.Println(v)
	}

	fmt.Println(days)
}