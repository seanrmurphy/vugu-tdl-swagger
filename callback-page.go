package main

import (
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

	c.getTokens(codeVerifier.String(), url.Query().Get("code"))
}

type TokenParams struct {
	GrantType    string `url:"grant_type,omitempty"`
	ClientID     string `url:"client_id,omitempty"`
	CodeVerifier string `url:"code_verifier,omitempty"`
	Code         string `url:"code,omitempty"`
	RedirectURI  string `url:"redirect_uri,omitempty"`
}

func (c *CallbackPage) getTokens(v, code string) {

	u := url.URL{
		Scheme: "https",
		Host:   "initialtest.auth.eu-west-1.amazoncognito.com",
		Path:   "oauth2/token",
		Opaque: "//initialtest.auth.eu-west-1.amazoncognito.com/oauth2/token",
	}

	clientID := "7cvg3l59uc6u1kqdcejcdso6rh"

	t := TokenParams{
		GrantType:    "authorization_code",
		ClientID:     clientID,
		CodeVerifier: v,
		Code:         code,
		RedirectURI:  "http://localhost:8844/callback",
	}

	val, _ := query.Values(t)
	u.RawQuery = val.Encode()

	//req, _ := http.NewRequest("POST", u.String(), strings.NewReader(u.RawQuery))
	req, _ := http.NewRequest("POST", u.String(), nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Add("Authorization", "Basic "+clientID)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
