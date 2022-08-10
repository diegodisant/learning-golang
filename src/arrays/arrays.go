package main

import "fmt"

func main() {
  var numbers [5]int // fixed array size

  fmt.Println(numbers)

  numbers[0] = 100

  fmt.Println(numbers)
}
