package main

import "fmt"

func main() {
  var numberSlice []int = make([]int, 0)

  fmt.Println(numberSlice)

  numberSlice = append(numberSlice, 5)

  fmt.Println(numberSlice)

  var initialNumberSlice []int = make([]int, 5)

  fmt.Println(initialNumberSlice)

  initialNumberSlice = append(initialNumberSlice, 10)

  fmt.Println(initialNumberSlice)

  var fixedNumberSlice []int = make([]int, 5, 5)

  fmt.Println(fixedNumberSlice)

  fixedNumberSlice = append(fixedNumberSlice, 10)

  fmt.Println(fixedNumberSlice)

  fmt.Println("Copy operation")

  var firstArray []int = make([]int, 3)

  firstArray[0] = 1
  firstArray[1] = 2
  firstArray[2] = 3

  var secondArray []int = make([]int, 3)

  fmt.Println(fmt.Sprintf(
    "Copy from %v to %v",
    firstArray,
    secondArray,
  ))

  copy(secondArray, firstArray)

  fmt.Println("Copy result:", secondArray)
}
