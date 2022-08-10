package main

import (
	"fmt"
	"time"
)

func primaryProcessing(channel chan<- string) {
	for {
		channel <- "from primary"

		time.Sleep(time.Second * 2)
	}
}

func secondaryProcessing(channel chan<- string) {
	for {
		channel <- "from secondary"

		time.Sleep(time.Second * 3)
	}
}

func channelSelector(
	pChan <-chan string,
	sChan <-chan string,
) {
	for {
		select {
		case primaryMessage := <-pChan:
			fmt.Println(primaryMessage)
		case secondaryMessage := <-sChan:
			fmt.Println(secondaryMessage)
		case <-time.After(time.Second): //default case
			fmt.Println("timeout")
		}
	}
}

func main() {
	var mainChannel = make(chan string)
	var secondaryChannel = make(chan string)

	go primaryProcessing(mainChannel)
	go secondaryProcessing(secondaryChannel)
	go channelSelector(
		mainChannel,
		secondaryChannel,
	)

	var input string

	fmt.Scanln(&input)
}
