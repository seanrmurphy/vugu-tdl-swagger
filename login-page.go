package main

import (
	"log"
	"syscall/js"

	"net/url"

	"github.com/google/go-querystring/query"
	pkce "github.com/nirasan/go-oauth-pkce-code-verifier"
)

func (c *LoginPage) BeforeBuild() {
	cv := sessionStorageGet("codeVerifier")

	log.Printf("cv = %v, type = %v", cv, cv.Type())
	if cv.Type() == js.TypeNull {
		v, _ := pkce.CreateCodeVerifier()
		LoginData.CodeVerifier = v

		log.Printf("Creating new code verifier for login = %v", LoginData.CodeVerifier.String())

		sessionStorageSet("codeVerifier", LoginData.CodeVerifier.String())

		cv := sessionStorageGet("codeVerifier")
		log.Printf("cv= %v", cv)
	} else {
		c.codeVerifier = &pkce.CodeVerifier{
			Value: cv.String(),
		}
	}
}

func (c *LoginPage) createCognitoURI(p CognitoParameters) (u url.URL) {

	clientName := "initial-test"

	u = url.URL{
		Scheme: "https",
		Host:   clientName + ".auth.eu-west-1.amazoncognito.com",
		Path:   "oauth2/authorize",
		Opaque: "//" + clientName + ".auth.eu-west-1.amazoncognito.com/oauth2/authorize",
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

	window.Call("open", q.String(), "_self")
}
