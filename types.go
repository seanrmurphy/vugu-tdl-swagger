package main

import pkce "github.com/nirasan/go-oauth-pkce-code-verifier"

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

type AuthenticationDataType struct {
	ClientID     string
	ClientName   string
	LoginData    LoginDataType
	RestEndpoint string
	RedirectURI  string
}

type LoginDataType struct {
	LoggedIn       bool
	CodeVerifier   *pkce.CodeVerifier
	ResponseParams ResponseParams
}
