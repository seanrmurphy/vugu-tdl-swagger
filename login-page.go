package main

import (
	"log"
	"syscall/js"

	"net/url"

	"github.com/google/go-querystring/query"
	pkce "github.com/nirasan/go-oauth-pkce-code-verifier"
)

type CognitoParameters struct {
	ResponseType        string `url:"response_type,omitempty"`
	ClientID            string `url:"client_id,omitempty"`
	RedirectURI         string `url:"redirect_uri,omitempty"`
	State               string `url:"state,omitempty"`
	IdentityProvider    string `url:"identity_provider,omitempty"`
	IDPProvider         string `url:"idp_provider,omitempty"`
	Scope               string `url:"scope,omitempty"`
	CodeChallengeMethod string `url:"code_challenge_method,omitempty"`
	CodeChallenge       string `url:"code_challenge,omitempty"`
}

func (c *LoginPage) BeforeBuild() {

	cv := sessionStorageGet("codeVerifier")

	log.Printf("cv = %v, type = %v", cv, cv.Type())
	if cv.Type() == js.TypeNull {
		v, _ := pkce.CreateCodeVerifier()
		c.codeVerifier = v

		log.Printf("Creating new code verifier for login = %v", c.codeVerifier.String())

		sessionStorageSet("codeVerifier", c.codeVerifier.String())

		cv := sessionStorageGet("codeVerifier")
		log.Printf("cv= %v", cv)
	} else {
		c.codeVerifier = &pkce.CodeVerifier{
			Value: cv.String(),
		}
	}
}

func (c *LoginPage) createCognitoURI(p CognitoParameters) (u url.URL) {
	u = url.URL{
		Scheme: "https",
		Host:   "initialtest.auth.eu-west-1.amazoncognito.com",
		Path:   "oauth2/authorize",
		Opaque: "//initialtest.auth.eu-west-1.amazoncognito.com/oauth2/authorize",
	}
	v, _ := query.Values(p)
	u.RawQuery = v.Encode()
	log.Printf("RawQuery = %v", u.RawQuery)
	return
}

func (c *LoginPage) GotoLogin() {
	log.Printf("About to redirect to login page...")
	//v, _ := pkce.CreateCodeVerifier()
	//codeVerifier := v.String()

	//sessionStorageSet("codeVerifier", codeVerifier)

	//log.Printf("Code verifiers = %v", codeVerifier)

	log.Printf("About to create challenge...")
	challenge := c.codeVerifier.CodeChallengeS256()
	challengeMethod := "S256"

	clientID := "7cvg3l59uc6u1kqdcejcdso6rh"

	p := CognitoParameters{
		ResponseType:        "code",
		ClientID:            clientID,
		RedirectURI:         "http://localhost:8844/callback",
		State:               "initial-state",
		IdentityProvider:    "COGNITO",
		IDPProvider:         "",
		Scope:               "profile",
		CodeChallengeMethod: challengeMethod,
		CodeChallenge:       challenge,
	}

	q := c.createCognitoURI(p)

	log.Printf("Redirecting to...%v", q.String())
	window := js.Global().Get("window")
	log.Printf("window = %v", window)
	//params := [2]string{"https://www.google.com", "_self"}
	//params = append(params, "https://www.google.com", "_self")
	window.Call("open", q.String(), "_self")
	//window.Call("open", q.String())
	//dispatcher.Dispatch(&actions.Login{})
}
