package main

import "fmt"

func main() {
  var dataMap map[string]int32 = make(map[string]int32)

  dataMap["x0"] = 9999
  dataMap["x1"] = 7129

  fmt.Println(dataMap)
}
