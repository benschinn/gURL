package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
  //"net/http/httputil"
	"os"
	"regexp"
  //"strings"
  //"net/url"
  //"bytes"
  //"log"
  //"os/exec"
)

type UserCred struct {
  email string
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
/*
func authenticate(csrfToken string) {
  requestURL := url.URL{
    Scheme: "https",
    Host:   "dashboard.qa.internal.mx",
    Path:   "/sign_in",
  }
  reqBody := strings.NewReader(`{
    "v1_analytics_user": {
      "email":   {"demo@dashboard.com"},
      "password: {"password0"},
    },
    "authenticity_token": {csrfToken}
  }`)

  requestHeaders := http.Header{
    "Content-Type": {"application/vnd.moneydesktop.sso.v3+json"},
    "Accept":       {"application/x-www-form-urlencoded"},
    "X-CSRF-Token":   {csrfToken},
  }

  jsonBody := []byte(string(reqBody))

  request := http.Request{
    Method:        "POST",
    URL:           &requestURL,
    Header:        requestHeaders,
    Body:          ioutil.NopCloser(bytes.NewReader(jsonBody)),
    ContentLength: int64(len(jsonBody)),
  }
  dump, err := httputil.DumpRequest(&request, true)

  if err != nil {
    fmt.Println("dump err", err.Error())
  }

  // Make the request
  fmt.Println("******** REQUEST ********")
  fmt.Println(string(dump))
}
*/

func main() {
	csrf := getToken()
  fmt.Println(csrf)
  /*
  //authenticate(csrf)
  //cmd := exec.Command("curl", "https://dashboard.qa.internal.mx/sign_in")
  curlReq := "https://dashboard.qa.internal.mx/sign_in"
  userCred := "v1_analytics_user[email]=demo@dashboard.com&v1_analytics_user[password]=password0"
  token :="authenticity_token=\"" + csrf + "\""
  fmt.Println(curlReq)
  cmd := exec.Command("curl", curlReq, "--data", userCred, "--data-urlencode", token, "--cookie", "cookie", "--cookie-jar", "cookie")
  //cmd.Stdin = strings.NewReader("some input")
  var out bytes.Buffer
  cmd.Stdout = &out
  err := cmd.Run()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf(out.String())
  */
}
