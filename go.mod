module github.com/vugu-examples/simple

go 1.14

replace github.com/vugu/vgrouter => ../../vugu/vgrouter

replace github.com/seanrmurphy/go-vecty-swagger/models => ../../seanrmurphy/go-vecty-swagger/models

require (
	github.com/fatih/structs v1.1.0
	github.com/go-openapi/runtime v0.19.19 // indirect
	github.com/go-openapi/strfmt v0.19.5
	github.com/google/go-querystring v1.0.0
	github.com/google/uuid v1.1.1
	github.com/gopherjs/vecty/example v0.0.0-20200328200803-52636d1f7aba
	github.com/nirasan/go-oauth-pkce-code-verifier v0.0.0-20170819232839-0fbfe93532da
	github.com/seanrmurphy/go-vecty-swagger v0.0.0-20200703103421-3b3d2d4e8515
	github.com/vugu/vgrouter v0.0.0-20200428001807-e4d5549422bc
	github.com/vugu/vjson v0.0.0-20200505061711-f9cbed27d3d9
	github.com/vugu/vugu v0.3.2
)
