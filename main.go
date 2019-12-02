package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func getToken() string {
	url := os.Args[1]
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Had a problem making request!", err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	html := string(body)
	r := regexp.MustCompile(`<meta name="csrf-token" content=".*?(.*)\/>`)
	metaTag := r.FindString(html)

	r_ := regexp.MustCompile(`content=".*?(.*)\/>`)
	contentAttribute := r_.FindString(metaTag)

	r__ := regexp.MustCompile(`".*?(.*)\"`)
	tokenInQuotes := r__.FindString(contentAttribute)

	r___ := regexp.MustCompile(`[\w | \d].*?(.*)\=`)
	token := r___.FindString(tokenInQuotes)

	return token
}

func main() {
	csrf := getToken()

	fmt.Println(csrf)
}
