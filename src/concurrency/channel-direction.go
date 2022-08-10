package main

// channels are the way to communbicate two go routines
// and synchronize the execution

import (
	"fmt"
	"time"
)

// i/o signal:    myChan chan string
// input signal:  myChan chan<- string
// output signal: myChan <-chan string

func pinger(signal chan<- string) {
	for {
		signal <- "ping"

		time.Sleep((time.Second * 1))
	}
}

func ponger(signal chan<- string) {
	for {
		signal <- "pong"

		time.Sleep((time.Second * 1))
	}
}

func printer(signal <-chan string) {
	//var message string = ""

	for {
		//message = <-signal

		//fmt.Println(message)

    // all is same to
    fmt.Println(<-signal)
	}
}

func main() {
	var pingPongSignal = make(chan string)

	go pinger(pingPongSignal)
	go ponger(pingPongSignal)
	go printer(pingPongSignal)

	var input string

	fmt.Scanln(&input)
}
