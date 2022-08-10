package main

import "fmt"

func resetNumber(number *int) {
	*number = 0
}

func main() {
	var number int = 5

	fmt.Println("number:", number)

	resetNumber(&number)

	fmt.Println("number:", number)
}
