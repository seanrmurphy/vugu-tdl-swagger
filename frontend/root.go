package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"syscall/js"

	"github.com/google/go-querystring/query"
)

func (c *Root) BeforeBuild() {

	url, _ := readBrowserURL()
	code := url.Query().Get("code")
	if code != "" && AuthenticationData.LoginData.LoggedIn == false {
		codeVerifier := sessionStorageGet("codeVerifier")
		log.Printf("code verifier = %v", codeVerifier.String())

		r := c.getTokens(codeVerifier.String(), url.Query().Get("code"))
		AuthenticationData.LoginData.ResponseParams = r
		AuthenticationData.LoginData.LoggedIn = true
	}
}

func (c *Root) createCognitoURI(p CognitoParameters) (u url.URL) {

	u = url.URL{
		Scheme: "https",
		Host:   AuthenticationData.ClientName + ".auth.eu-west-2.amazoncognito.com",
		Path:   "oauth2/authorize",
		Opaque: "//" + AuthenticationData.ClientName + ".auth.eu-west-2.amazoncognito.com/oauth2/authorize",
	}
	v, _ := query.Values(p)
	u.RawQuery = v.Encode()
	log.Printf("RawQuery = %v", u.RawQuery)
	return
}

func (c *Root) Logout() {
	log.Printf("Not implemented")
}

func (c *Root) Login() {

	challenge := AuthenticationData.LoginData.CodeVerifier.CodeChallengeS256()
	challengeMethod := "S256"

	p := CognitoParameters{
		ResponseType:        "code",
		ClientID:            AuthenticationData.ClientID,
		RedirectURI:         AuthenticationData.RedirectURI,
		State:               "initial-state",
		IdentityProvider:    "COGNITO",
		IDPProvider:         "",
		Scope:               "profile",
		CodeChallengeMethod: challengeMethod,
		CodeChallenge:       challenge,
	}

	q := c.createCognitoURI(p)

	window := js.Global().Get("window")
	window.Call("open", q.String(), "_self")
}

func (c *Root) getTokens(v, code string) (r ResponseParams) {

	u := url.URL{
		Scheme: "https",
		Host:   AuthenticationData.ClientName + ".auth.eu-west-2.amazoncognito.com",
		Path:   "oauth2/token",
		Opaque: "//" + AuthenticationData.ClientName + ".auth.eu-west-2.amazoncognito.com/oauth2/token",
	}

	t := TokenParams{
		GrantType:    "authorization_code",
		ClientID:     AuthenticationData.ClientID,
		CodeVerifier: v,
		Code:         code,
		RedirectURI:  AuthenticationData.RedirectURI,
	}

	val, _ := query.Values(t)
	u.RawQuery = val.Encode()

	req, _ := http.NewRequest("POST", u.String(), nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	r = ResponseParams{}
	json.Unmarshal(body, &r)
	fmt.Printf("Access Token: %v", r.AccessToken)
	return
}
