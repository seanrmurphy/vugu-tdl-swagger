package main

import (
	"log"
	"net/url"
	"syscall/js"

	"github.com/google/go-querystring/query"
)

func (c *Root) BeforeBuild() {
}

func (c *Root) createCognitoURI(p CognitoParameters) (u url.URL) {

	u = url.URL{
		Scheme: "https",
		Host:   AuthenticationData.ClientName + ".auth.eu-west-1.amazoncognito.com",
		Path:   "oauth2/authorize",
		Opaque: "//" + AuthenticationData.ClientName + ".auth.eu-west-1.amazoncognito.com/oauth2/authorize",
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
		RedirectURI:         "http://localhost:8844/callback",
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
