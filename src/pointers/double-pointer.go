package main

import "fmt"

func main() {
  var number int = 5
  var numberPointer *int = &number

  fmt.Println("number", number)

  fmt.Println("numberPointer", numberPointer)
  fmt.Println("&numberPointer", &numberPointer)
  fmt.Println("*numberPointer", *numberPointer)

  var dNumberPointer **int = &numberPointer;

  fmt.Println("dNumberPointer", dNumberPointer)
  fmt.Println("&dNumberPointer", &dNumberPointer)
  fmt.Println("**dNumberPointer", **dNumberPointer)
}
