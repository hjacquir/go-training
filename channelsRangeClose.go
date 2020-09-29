package main

import "fmt"

func fibonacci(limit int, channel chan int) {
	y, x := 1, 0
	for i := 0; i < limit; i++ {
		channel <- x
		x = y
		y = x + y
	}

	close(channel)
}

func main() {
	var bufferSize int
	bufferSize = 10
	channelBufferedOfInt := make(chan int, bufferSize)
	channelCapacity := cap(channelBufferedOfInt)

	go fibonacci(channelCapacity, channelBufferedOfInt)

	for i := range channelBufferedOfInt {
		fmt.Println(i)
	}
}