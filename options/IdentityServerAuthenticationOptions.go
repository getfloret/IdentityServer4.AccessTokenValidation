package options

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/IdentityModel"
	"time"
)

type TokenClaimsProcessor interface {
	Process(claims jwt.MapClaims) (jwt.MapClaims, error)
}

type IdentityServerAuthenticationOptions struct {
	// Base-address of the token issuer
	Authority string
	/// Specifies whether HTTPS is required for the discovery endpoint
	RequireHttpsMetadata bool
	//
	DiscoveryHttpClientTimeout time.Duration
	//
	DiscoveryHttpClientTLSTimeout time.Duration

	/// Specifies which token types are supported (JWT, reference or both)
	SupportedTokens IdentityModel.SupportedTokens

	///// <summary>
	///// Callback to retrieve token from incoming request
	///// </summary>
	//public Func<HttpRequest, string> TokenRetriever { get; set; } = TokenRetrieval.FromAuthorizationHeader();

	// Name of the API resource used for authentication against introspection endpoint
	ApiName string

	// Secret used for authentication against introspection endpoint
	ApiSecret string

	/* Enable if this API is being secured by IdentityServer3, and if you need to support both JWTs and reference tokens.
	If you enable this, you should add scope validation for incoming JWTs.*/
	LegacyAudienceValidation bool

	ClaimsProcessor TokenClaimsProcessor

	// Claim type for name
	NameClaimType string

	// Claim type for role
	RoleClaimType string

	// Specifies inbound claim type map for JWT tokens (mainly used to disable the annoying default behavior of the MS JWT handler)
	InboundJwtClaimTypeMap map[string]string

	// Specifies whether caching is enabled for introspection responses (requires a distributed cache implementation)
	EnableCaching bool

	// Specifies ttl for introspection response caches
	CacheDuration time.Duration

	// Specifies the prefix of the cache key (token).
	CacheKeyPrefix string

	// Gets or sets the policay for the introspection discovery document.
	IntrospectionDiscoveryPolicy IdentityModel.DiscoveryPolicy

	// specifies whether the token should be saved in the authentication properties
	SaveToken bool

	// specifies the allowed clock skew when validating JWT tokens
	JwtValidationClockSkew time.Duration

	///// back-channel handler for JWT middleware
	//public HttpMessageHandler JwtBackChannelHandler { get; set; }

	///// <summary>
	///// back-channel handler for introspection endpoint
	///// </summary>
	//public HttpMessageHandler IntrospectionBackChannelHandler { get; set; }

	///// <summary>
	///// back-channel handler for introspection discovery endpoint
	///// </summary>
	//public HttpMessageHandler IntrospectionDiscoveryHandler { get; set; }

	// timeout for back-channel operations
	BackChannelTimeouts time.Duration

	//// todo
	///// <summary>
	///// events for JWT middleware
	///// </summary>
	//public JwtBearerEvents JwtBearerEvents { get; set; } = new JwtBearerEvents();

	/* Specifies how often the cached copy of the discovery document should be refreshed.
	If not set, it defaults to the default value of Microsoft's underlying configuration manager (which right now is 24h).
	If you need more fine grained control, provide your own configuration manager on the JWT options.*/
	DiscoveryDocumentRefreshInterval time.Duration



}

// Gets a value indicating whether JWTs are supported.
func (options *IdentityServerAuthenticationOptions) SupportsJwt()bool{
	return options.SupportedTokens == IdentityModel.SupportedTokens_Jwt || options.SupportedTokens == IdentityModel.SupportedTokens_Both
}

// Gets a value indicating whether reference tokens are supported.
func (options *IdentityServerAuthenticationOptions) SupportsIntrospection()bool{
	return options.SupportedTokens == IdentityModel.SupportedTokens_Reference || options.SupportedTokens == IdentityModel.SupportedTokens_Both
}

const (
	defaultDiscoveryHttpClientTimeout = 10 * time.Second
	defaultDiscoveryHttpClientTLSTimeout = 10 * time.Second
	defaultCacheDuration              = 10 * time.Minute
	defaultBackChannelTimeouts              = 60 * time.Second
	defaultDiscoveryDocumentRefreshInterval         = 24 * time.Hour
)
var (
	// AppSettings is a global variable that holds the application settings
	GlobalAuthenticationOptions = defaultOptions()
)
func defaultOptions() *IdentityServerAuthenticationOptions {
	return &IdentityServerAuthenticationOptions{
		DiscoveryHttpClientTimeout: defaultDiscoveryHttpClientTimeout,
		DiscoveryHttpClientTLSTimeout: defaultDiscoveryHttpClientTLSTimeout,
		CacheDuration:                     defaultCacheDuration,
		BackChannelTimeouts:              defaultBackChannelTimeouts,
		DiscoveryDocumentRefreshInterval:     defaultDiscoveryDocumentRefreshInterval,
		RequireHttpsMetadata: true,
		SupportedTokens: IdentityModel.SupportedTokens_Both,
		NameClaimType: "name",
		RoleClaimType: "role",
		EnableCaching: false,
		CacheKeyPrefix: "",
		SaveToken: true,
	}
}