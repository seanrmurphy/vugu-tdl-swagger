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

	c.getTokens(codeVerifier.String(), url.Query().Get("code"))
}

type TokenParams struct {
	GrantType    string `url:"grant_type,omitempty"`
	ClientID     string `url:"client_id,omitempty"`
	CodeVerifier string `url:"code_verifier,omitempty"`
	Code         string `url:"code,omitempty"`
	RedirectURI  string `url:"redirect_uri,omitempty"`
}

type ResponseParams struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempt"`
	TokenType    string `json:"token_type,omitempty"`
}

func (c *CallbackPage) getTokens(v, code string) {

	clientName := "todo-api-client"
	clientID := "6vqii43vld9jg0odddtom70gse"

	u := url.URL{
		Scheme: "https",
		Host:   clientName + ".auth.eu-west-1.amazoncognito.com",
		Path:   "oauth2/token",
		Opaque: "//" + clientName + ".auth.eu-west-1.amazoncognito.com/oauth2/token",
	}

	t := TokenParams{
		GrantType:    "authorization_code",
		ClientID:     clientID,
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

	authResponse := ResponseParams{}
	json.Unmarshal(body, authResponse)
	fmt.Printf("Access Token: %v", authResponse.AccessToken)

}
