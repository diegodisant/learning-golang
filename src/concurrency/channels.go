package main

// channels are the way to communbicate two go routines
// and synchronize the execution

import (
	"fmt"
	"time"
)

type stringSignal chan string

func pinger(signal stringSignal) {
	for {
		signal <- "ping"

		time.Sleep((time.Second * 1))
	}
}

func ponger(signal stringSignal) {
	for {
		signal <- "pong"

		time.Sleep((time.Second * 1))
	}
}

func printer(signal stringSignal) {
	//var message string = ""

	for {
		//message = <-signal

		//fmt.Println(message)

    // all is same to
    fmt.Println(<-signal)
	}
}

func main() {
	var pingPongSignal stringSignal = make(stringSignal)

	go pinger(pingPongSignal)
	go ponger(pingPongSignal)
	go printer(pingPongSignal)

	var input string

	fmt.Scanln(&input)
}
