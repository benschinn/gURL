package main

import (
	"bytes"
	//"encoding/json"
  "net/http/httputil"
  "net/url"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

var signInUrl string = os.Args[1]

func getToken() string {
	resp, err := http.Get(signInUrl)

	if err != nil {
		fmt.Println("Had a problem making request!", err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	html := string(body)
	//grab meta tag with csrf token
	r := regexp.MustCompile(`<meta name="csrf-token" content=".*?(.*)\/>`)
	metaTag := r.FindString(html)

	//grab just the content attribute
	r_ := regexp.MustCompile(`content=".*?(.*)\/>`)
	contentAttribute := r_.FindString(metaTag)

	//grab token in quotes
	r__ := regexp.MustCompile(`".*?(.*)\"`)
	tokenInQuotes := r__.FindString(contentAttribute)

	//grab token without quotes
	r___ := regexp.MustCompile(`[\w | \d].*?(.*)\=`)
	token := r___.FindString(tokenInQuotes)
	d1 := []byte(string(body))
	errr := ioutil.WriteFile("./token.html", d1, 0644)
	if errr != nil {
		fmt.Println("error writing file")
	}

	return token
}
func authenticate(csrfToken string) {
  requestURL := url.URL{
    Scheme: "http",
    Host:   "dashboard.localhost:3000",
    Path:   "/sign_in",
  }

  // Build the headers
  requestHeaders := http.Header{
    "Content-Type": {"application/x-www-form-urlencoded"},
    "Accept-Language": {"en-US,en;q=0.9"},
  }
  var reqBody = []byte(fmt.Sprintf(`{
		"v1_analytics_user[email]":    "demo@dashboard.com",
		"v1_analytics_user[password]": "password0",
		"authenticity_token":          %s,
		"commit":                      "SIGN IN",
	}`, csrfToken))

  request := http.Request{
    Method:        "POST",
    URL:           &requestURL,
    Header:        requestHeaders,
    Body:          ioutil.NopCloser(bytes.NewReader(reqBody)),
    ContentLength: int64(len(reqBody)),
  }
  dump, err := httputil.DumpRequest(&request, true)
  if err != nil {
    fmt.Println("dump err", err.Error())
  }

  // Make the request
  fmt.Println("******** REQUEST ********")
  fmt.Println(string(dump))
  resp, err := http.DefaultClient.Do(&request)

  defer resp.Body.Close()

  if err != nil {
    fmt.Println("Request didn't make it!", err.Error())
  }

  // Parse the response
  responseBody, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    fmt.Println("Failed to read body", err.Error())
  }
  fmt.Println(resp)

	d1 := []byte(string(responseBody))
	errr := ioutil.WriteFile("./response.html", d1, 0644)
	if errr != nil {
		fmt.Println("error writing file")
	}
}

func main() {
	csrf := getToken()
	authenticate(csrf)
}
