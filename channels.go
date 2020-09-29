package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func sum2(s []int) int {
	sum2 := 0

	for _,v := range s {
		sum2 += v
	}

	return sum2
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	var slice1 = s[0:len(s)/2]
	var slice2 = s[len(s)/2:]

	go sum(slice1, c)
	go sum(slice2, c)
	x, y := <-c, <-c // receive from c

	//x2 := sum2(slice1)
	//y2 := sum2(slice2)


	fmt.Println(x, y, x+y)
	//fmt.Println(x2, y2)
}
