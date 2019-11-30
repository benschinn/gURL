package main

import (
  "fmt"
  "os"
  "net/http"
  "net/http/cookiejar"
)

func main() {
  jar, _ := cookiejar.New(nil)
  url := os.Args[1]
  resp, err := http.Get(url)
  if err != nil {
    fmt.Println("Had a problem making request!", err.Error())
  }
  fmt.Println(resp)
}
