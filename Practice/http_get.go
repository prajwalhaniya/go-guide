package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

func main() {
    resp, _ := http.Get("http://prajwalhaniya.in/")
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
    resp.Body.Close()
}
