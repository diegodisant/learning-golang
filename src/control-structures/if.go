package main

import "fmt"

func main() {
  var mainResult bool = true

  mainResult = !mainResult

  if mainResult {
    fmt.Println("The result is true")
  } else {
    fmt.Println("The result is false")
  }
}
