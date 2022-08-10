package main

import "fmt"

func main() {
  var number int
  var numberPointer *int

  numberPointer = &number // address or reference operator

  fmt.Println(numberPointer) // print the address
  fmt.Println(*numberPointer) // dereference and print the value

  // change the original value
  number = 5

  fmt.Println(numberPointer)

  // print the changed value cause pointer is only pointing to the reference
  fmt.Println(*numberPointer)
}
