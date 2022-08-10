package main

import (
	"fmt"
	"strings"
)

func main() {
  var firstString string = "Hello my name is "
  var secondString string = "Diego"

  fmt.Println(firstString + secondString)

  fmt.Println(strings.Compare(secondString, "Diego"))
  fmt.Println(strings.Count(firstString, "l"))

  secondString = replaceChartAtIndex(secondString, 'e', 1)

  fmt.Println(secondString)
}

func replaceChartAtIndex(
  replaceString string,
  replaceChar rune,
  index int,
) string {
  chars := []rune(replaceString)

  chars[index] = replaceChar

  return string(chars)
}
