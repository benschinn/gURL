package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
)

func main() {
	jar, _ := cookiejar.New(nil)
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Had a problem making request!", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	html := string(body)
	fmt.Println(html)
	fmt.Println(jar)
	fmt.Println(resp)
}
