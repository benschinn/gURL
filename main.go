package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
  "net/http/cookiejar"
	"net/url"
	"os"
  "log"
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
  cj, _ := cookiejar.New(nil)

  client := &http.Client{
    Jar: cj,
  }

	requestURL := url.URL{
		Scheme: "http",
		Host:   "dashboard.localhost:3000",
		Path:   "/sign_in",
	}

	requestHeaders := http.Header{
    "Accept":          {"*/*"},
    "Accept-Encoding": {"gzip, defalte"},
		"Content-Type":    {"application/x-www-form-urlencoded"},
		"Accept-Language": {"en-US,en;q=0.9"},
	}

	form := url.Values{"v1_analytics_user[email]": {"demo@dashboard.com"}, "v1_analytics_user[password]": {"password0"}, "authenticity_token": {csrfToken}}

	request := http.Request{
		Method:        "POST",
		URL:           &requestURL,
		Header:        requestHeaders,
		Body:          ioutil.NopCloser(bytes.NewReader([]byte(form.Encode()))),
		ContentLength: int64(len(form.Encode())),
	}
	dump, err := httputil.DumpRequest(&request, true)
	if err != nil {
		fmt.Println("dump err", err.Error())
	}

	fmt.Println("******** REQUEST ********")
	fmt.Println(string(dump))
  fmt.Println("CONTENT LENGTH", request.ContentLength)

	resp, err := client.Do(&request)

  //fmt.Println("ERROR", err.Error())
  fmt.Println(resp)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
  fmt.Println(resp.Body)

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Failed to read body", err.Error())
	}

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
