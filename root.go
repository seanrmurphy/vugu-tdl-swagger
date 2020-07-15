package main

import (
	"log"
	"net/url"
	"syscall/js"

	"github.com/google/go-querystring/query"
	pkce "github.com/nirasan/go-oauth-pkce-code-verifier"
)

func (c *Root) BeforeBuild() {

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
		LoginData.CodeVerifier = &pkce.CodeVerifier{
			Value: cv.String(),
		}
		log.Printf("LoginData.LoggedIn = %v\n", LoginData.LoggedIn)
	}
}

func (c *Root) createCognitoURI(p CognitoParameters) (u url.URL) {

	clientName := "initialtest"

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

func (c *Root) Logout() {
	log.Printf("Not implemented")
}

func (c *Root) Login() {
	log.Printf("About to redirect to login page...")

	//v, _ := pkce.CreateCodeVerifier()
	//codeVerifier := v.String()

	//sessionStorageSet("codeVerifier", codeVerifier)

	//log.Printf("Code verifiers = %v", codeVerifier)

	log.Printf("About to create challenge...")
	challenge := LoginData.CodeVerifier.CodeChallengeS256()
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
