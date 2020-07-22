module github.com/seanrmurphy/vugu-tdl-swagger/frontend

go 1.14

replace github.com/vugu/vgrouter => ../../vgrouter

require (
	github.com/fatih/structs v1.1.0
	github.com/go-openapi/runtime v0.19.19
	github.com/go-openapi/strfmt v0.19.5
	github.com/google/go-querystring v1.0.0
	github.com/google/uuid v1.1.1
	github.com/gopherjs/vecty/example v0.0.0-20200328200803-52636d1f7aba
	github.com/nirasan/go-oauth-pkce-code-verifier v0.0.0-20170819232839-0fbfe93532da
	github.com/seanrmurphy/vugu-tdl-swagger v0.0.0-20200720114405-ee4421fdb522
	github.com/vugu/vgrouter v0.0.0-00010101000000-000000000000
	github.com/vugu/vgrun v0.0.0-20200413095632-da7c2a7eb99c // indirect
	github.com/vugu/vjson v0.0.0-20200505061711-f9cbed27d3d9
	github.com/vugu/vugu v0.3.2
	golang.org/x/sys v0.0.0-20200720211630-cb9d2d5c5666 // indirect
)
