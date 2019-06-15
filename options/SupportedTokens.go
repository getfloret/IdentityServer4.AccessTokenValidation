package options

type SupportedTokens int
const (
	_ SupportedTokens = iota
	// JWTs and reference tokens
	SupportedTokens_Both
	// JWTs only
	SupportedTokens_Jwt
	// Reference tokens only
	SupportedTokens_Reference
)