package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

func (c *CallbackPage) BeforeBuild() {
	log.Printf("Before builder...")

	url, _ := readBrowserURL()
	log.Printf("url = %v", url)

	for k, v := range url.Query() {
		log.Printf("param = %v, val = %v", k, v)
	}

	codeVerifier := sessionStorageGet("codeVerifier")
	log.Printf("code verifier = %v", codeVerifier.String())

	r := c.getTokens(codeVerifier.String(), url.Query().Get("code"))
	AuthenticationData.LoginData.ResponseParams = r
	AuthenticationData.LoginData.LoggedIn = true

	// go to main landing page...
	//c.Navigate("/", nil)
}

func (c *CallbackPage) getTokens(v, code string) (r ResponseParams) {

	u := url.URL{
		Scheme: "https",
		Host:   AuthenticationData.ClientName + ".auth.eu-west-1.amazoncognito.com",
		Path:   "oauth2/token",
		Opaque: "//" + AuthenticationData.ClientName + ".auth.eu-west-1.amazoncognito.com/oauth2/token",
	}

	t := TokenParams{
		GrantType:    "authorization_code",
		ClientID:     AuthenticationData.ClientID,
		CodeVerifier: v,
		Code:         code,
		RedirectURI:  "http://localhost:8844/callback",
	}

	val, _ := query.Values(t)
	u.RawQuery = val.Encode()

	req, _ := http.NewRequest("POST", u.String(), nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("body = %v", string(body))

	r = ResponseParams{}
	json.Unmarshal(body, &r)
	fmt.Printf("Access Token: %v", r.AccessToken)
	return
}
