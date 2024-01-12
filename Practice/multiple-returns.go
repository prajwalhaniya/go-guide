package main

import (
  "fmt"
)

func Names() (string, string) {
    return "Foo", "Bar"
}

func main() {
    p1, p2 := Names()
    fmt.Println(p1, p2)

    p3, _ := Names()
    fmt.Println(p3)
}
