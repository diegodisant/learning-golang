package main

import "fmt"

func main() {
  var firstConstraint bool = true
  var secondConstraint bool = false

  fmt.Println("a && b", firstConstraint && secondConstraint)
  fmt.Println("a || b", firstConstraint || secondConstraint)
}
