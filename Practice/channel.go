package main

import (
  "fmt"
  "time"
)

func printCount(c chan int) {
  num := 0
  for num >= 0 {
    num = <-c
    fmt.Println(num, " ")
  }
}

func main() {
  c := make(chan int)
  a := []int{8, 6, 7, 1, 2, 3}

  go printCount(c)

  for _, v := range a {
    c <- v
  }

  time.Sleep(time.Millisecond * 1)
  fmt.Println("END OF MAIN")
}
