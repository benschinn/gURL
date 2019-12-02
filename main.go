package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"net/http/httputil"
	"bytes"
	"encoding/json"
	"os"
	"regexp"
	"time"
	//"log"
	//"os/exec"
)

type UserCred struct {
	email    string
	password string
}

var signInUrl string = os.Args[1]

func getToken() string {
	resp, err := http.Get(signInUrl)

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
func authenticate(csrfToken string) {
	reqBody, err := json.Marshal(map[string]string{
		"v1_analytics_user[email]":    "demo@dashboard.com",
		"v1_analytics_user[password]": "password0",
		"authenticity_token":          csrfToken,
		"commit":                      "SIGN IN",
	})

	if err != nil {
		fmt.Println("Had a problem authenticating!", err.Error())
	}

	timeout := time.Duration(7 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", signInUrl, bytes.NewBuffer(reqBody))
	request.Header.Set("Content-type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept-language", "en-US,en;q=0.9")

	if err != nil {
		fmt.Println("Http client has problems.", err.Error())
	}

	resp, err := client.Do(request)

	if err != nil {
		fmt.Println("Error making request", err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body", err.Error())
	}

	fmt.Println(string(body))
	fmt.Println(bytes.NewBuffer(reqBody))
	fmt.Println(request.Body)
	fmt.Println(request)
	fmt.Println(resp)
}

func main() {
	csrf := getToken()
	authenticate(csrf)
}
