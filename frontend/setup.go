// setup.go

package main

import (
	"log"
	"syscall/js"

	pkce "github.com/nirasan/go-oauth-pkce-code-verifier"
	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"
)

var AuthenticationData AuthenticationDataType

func setupAuthentication() {
	AuthenticationData.ClientID = "6fst7hjms26vsahdp3c1pu95bc"
	AuthenticationData.ClientName = "todo-api-client"
	AuthenticationData.RestEndpoint = "https://2wwyvmz2zd.execute-api.eu-west-2.amazonaws.com/prod"
	AuthenticationData.RedirectURI = "http://localhost:8844"

	cv := sessionStorageGet("codeVerifier")

	if cv.Type() == js.TypeNull {
		v, _ := pkce.CreateCodeVerifier()
		AuthenticationData.LoginData.CodeVerifier = v

		log.Printf("Creating new code verifier for login = %v", AuthenticationData.LoginData.CodeVerifier.String())

		sessionStorageSet("codeVerifier", AuthenticationData.LoginData.CodeVerifier.String())

	} else {
		AuthenticationData.LoginData.CodeVerifier = &pkce.CodeVerifier{
			Value: cv.String(),
		}
	}
}

// OVERALL APPLICATION WIRING IN vuguSetup
func vuguSetup(buildEnv *vugu.BuildEnv, eventEnv vugu.EventEnv) vugu.Builder {

	setupAuthentication()

	// CREATE A NEW ROUTER INSTANCE
	router := vgrouter.New(eventEnv)

	// MAKE OUR WIRE FUNCTION POPULATE ANYTHING THAT WANTS A "NAVIGATOR".
	buildEnv.SetWireFunc(func(b vugu.Builder) {
		if c, ok := b.(vgrouter.NavigatorSetter); ok {
			c.NavigatorSet(router)
		}
	})

	// CREATE THE ROOT COMPONENT
	root := &Root{}
	buildEnv.WireComponent(root) // WIRE IT

	// ADD ROUTES FOR EACH PAGE.  NOTE THAT THESE ARE "EXACT" ROUTES.
	// YOU CAN ALSO ADD ROUTES THAT MATCH ANYTHING WITH THE SPECIFIED PREFIX.
	router.MustAddRouteExact("/",
		vgrouter.RouteHandlerFunc(func(rm *vgrouter.RouteMatch) {
			root.Body = &ToDoList{} // A COMPONENT WITH PAGE CONTENTS
		}))

	// TELL THE ROUTER TO LISTEN FOR THE BROWSER CHANGING URLS
	err := router.ListenForPopState()
	if err != nil {
		panic(err)
	}

	// GRAB THE CURRENT BROWSER URL AND PROCESS IT AS A ROUTE
	err = router.Pull()
	if err != nil {
		panic(err)
	}

	return root
}
