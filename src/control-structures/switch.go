package main

import "fmt"

func main() {
  var number int = 10

  switch(number) {
    case 0:
      fmt.Println("zero")
    case 9:
      fmt.Println("nine")
    default:
      fmt.Println("unknown number")
  }
}
