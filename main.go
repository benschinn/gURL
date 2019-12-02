package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func extractToken(html string) string {
	r := regexp.MustCompile(`<meta name="csrf-token" content=".*?(.*)\/>`)
	metaTag := r.FindString(html)

	r1 := regexp.MustCompile(`content=".*?(.*)\/>`)
	partOfTag := r1.FindString(metaTag)

	r2 := regexp.MustCompile(`".*?(.*)\"`)
	token := r2.FindString(partOfTag)

	r3 := regexp.MustCompile(`[\w | \d].*?(.*)\=`)
	csrfToken := r3.FindString(token)

	return csrfToken

}

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Had a problem making request!", err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	html := string(body)

	csrf := extractToken(html)

	fmt.Println(csrf)
}
