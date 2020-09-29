package main

import (
	"fmt"
)

func main() {
	channelNotBuffered()
}

func channelBuffered() {
	stringChannels := make(chan string, 2)

	stringChannels <- "s1"
	stringChannels <- "s2"

	fmt.Println(<-stringChannels)
	fmt.Println(<-stringChannels)
}

func channelNotBuffered() {
	stringChannels := make(chan string)

	go func() {
		stringChannels <- "received"
	}()

	receiver := <- stringChannels

	fmt.Println(receiver)
}