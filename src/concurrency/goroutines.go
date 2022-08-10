package main

import "fmt"

func printSequence(max int) {
	for number := 1; number <= max; number += 1 {
		fmt.Println(number)
	}
}

func main() {
	go printSequence(1000)

	var input string

	fmt.Scanln(&input)
}
