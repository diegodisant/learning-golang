package main

import "fmt"

func main() {
  var numbers []int = make([]int, 10)

  for
    numberIndex := 0;
    numberIndex < len(numbers);
    numberIndex += 1 {
    fmt.Println(numberIndex, numbers[numberIndex])
  }

  for numberIndex, number := range numbers {
    fmt.Println(numberIndex, number)
  }
}
