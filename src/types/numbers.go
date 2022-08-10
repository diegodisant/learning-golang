package main

import "fmt"

func main() {
  var normalNumber int = 20

  fmt.Println(normalNumber)

  var int8Number int8 = 127
  var uInt8Number uint8 = 255

  fmt.Println(int8Number)
  fmt.Println(uInt8Number)

  var int16Number int16 = 32765
  var uInt16Number uint16 = 65535

  fmt.Println(int16Number)
  fmt.Println(uInt16Number)

  var int32Number int32 = -2147483648
  var uInt32Number uint32 = 2147483648

  fmt.Println(int32Number)
  fmt.Println(uInt32Number)

  var int64Number int64 = -2147483648;
  var uInt64Number uint64 = 2147483648

  fmt.Println(int64Number)
  fmt.Println(uInt64Number)

  // operators

  fmt.Println("bit operators")

  fmt.Println("u8 & 1 =", uInt8Number & 1)
  fmt.Println("u8 & 128 =", uInt8Number & 128)

  fmt.Println("u8 | 1 =", uInt8Number | 1)
  fmt.Println("u8 | 128 =", uInt8Number | 128)

  fmt.Println("u8 ~ 1 =", uInt8Number ^ 1)
  fmt.Println("u8 ~ 128 =", uInt8Number ^ 128)

  fmt.Println("u8 >> 1 =", uInt8Number >> 1)
  fmt.Println("u8 << 1 =", uInt8Number << 1)

  fmt.Println("reset all bits -> u8 >> 10 =", uInt8Number >> 10)

  fmt.Println("arithmetical operators")

  fmt.Println("u8 - 1 =", uInt8Number - 1)
  fmt.Println("u8 - 100 =", uInt8Number - 100)

  uInt8Number <<= 10

  fmt.Println("u8 + 200", uInt8Number + 200)
}
