package main

import (
  "fmt"
  "math/big"
  "os"
  "strings"
)

func processHex(hexStr string) string {
  return strings.TrimPrefix(hexStr, "0x")
}

func main() {
  if len(os.Args) == 2 {
    inputHexNumber := os.Args[1]
    inputHexNumber = strings.ToLower(inputHexNumber)

    i := new(big.Int)

    if _, ok := i.SetString(processHex(inputHexNumber), 16); !ok {
      fmt.Println("Invailed hexdecimal number!")
    }

    fmt.Println(i)
  } else {
    fmt.Println("usage:", os.Args[0], "[0xYournumber]")
  }
}
