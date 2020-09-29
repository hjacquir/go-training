package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go setChannelWaiting(2, channel2)
	go setChannelWaiting(1, channel1)

	messages := [2]string{"first", "second"}

	for _, v := range messages {
		fmt.Println(v)
		select {
		case v := <-channel2:
			fmt.Println("received", v)
		case v := <-channel1:
			fmt.Println("received", v)
		}
	}
}

func setChannelWaiting(wait int, currentChannel chan string) {
	duration := time.Duration(wait) * time.Second
	time.Sleep(duration)
	currentChannel <- " after " + strconv.Itoa(wait) + " second"
}