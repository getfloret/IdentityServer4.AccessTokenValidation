package IdentityModel

// Specifies how the client will transmit client ID and secret
type ClientCredentialStyle int
const (
	_ ClientCredentialStyle = iota
	// HTTP basic authentication
	AuthorizationHeader
	// Post values in body
	PostBody
)
