package options

import (
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/IdentityModel"
	"time"
)

// Options class for the OAuth 2.0 introspection endpoint authentication handler
type OAuth2IntrospectionOptions struct {
/* Sets the base-path of the token provider.
If set, the OpenID Connect discovery document will be used to find the introspection endpoint.*/
Authority string

/* Sets the URL of the introspection endpoint.
If set, Authority is ignored.*/
IntrospectionEndpoint string

// Specifies the id of the introspection client (required).
 ClientId string

// Specifies the shared secret of the introspection client.
ClientSecret string

///// <summary>
///// Specifies the the client assertion to be used (optional replacement of simple client secret)
///// </summary>
//public ClientAssertion ClientAssertion { get; set; } = new ClientAssertion();

// Specifies how client id and secret are being sent
ClientCredentialStyle IdentityModel.ClientCredentialStyle

//todo
///// <summary>
///// Specifies the token type hint of the introspection client.
///// </summary>
//public string TokenTypeHint { get; set; } = OidcConstants.TokenTypes.AccessToken;

// Specifies the claim type to use for the name claim (defaults to 'name')
NameClaimType string

// Specifies the claim type to use for the role claim (defaults to 'role')
RoleClaimType string

///// <summary>
///// Specifies the authentication type to use for the authenticated identity.
///// If not set, the authentication scheme name is used as the authentication
///// type (defaults to null).
///// </summary>
// AuthenticationType string

// Specifies the policy for the discovery client
DiscoveryPolicy IdentityModel.DiscoveryPolicy

// Specifies whether tokens that contain dots (most likely a JWT) are skipped
SkipTokensWithDots bool

// Specifies whether the token should be stored in the context, and thus be available for the duration of the request
SaveToken bool

// Specifies whether the outcome of the toke validation should be cached. This reduces the load on the introspection endpoint at the STS
EnableCaching bool

// Specifies for how long the outcome of the token validation should be cached.
CacheDuration time.Duration

// Specifies the prefix of the cache key (token).
CacheKeyPrefix string

///// <summary>
///// Specifies the method how to retrieve the token from the HTTP request
///// </summary>
//public Func<HttpRequest, string> TokenRetriever { get; set; } = TokenRetrieval.FromAuthorizationHeader();

///// <summary>
///// Gets or sets the <see cref="OAuth2IntrospectionEvents"/> used to handle authentication events.
///// </summary>
//public new OAuth2IntrospectionEvents Events
//{
//get { return (OAuth2IntrospectionEvents)base.Events; }
//set { base.Events = value; }
//}

//internal AsyncLazy<IntrospectionClient> IntrospectionClient { get; set; }
//todo
//internal ConcurrentDictionary<string, AsyncLazy<TokenIntrospectionResponse>> LazyIntrospections { get; set; }
}

/// <summary>
/// Check that the options are valid. Should throw an exception if things are not ok.
/// </summary>
/// <exception cref="InvalidOperationException">
/// You must either set Authority or IntrospectionEndpoint
/// or
/// You must either set a ClientId or set an introspection HTTP handler
/// </exception>
/// <exception cref="ArgumentException">TokenRetriever must be set - TokenRetriever</exception>
func (option *OAuth2IntrospectionOptions) Validate()(){
	if option.Authority == "" && option.IntrospectionEndpoint == "" {
		//return errors.New("You must either set Authority or IntrospectionEndpoint")
		panic("You must either set Authority or IntrospectionEndpoint")
	}
	//
	//if TokenRetriever == null
	//{
	//	throw new ArgumentException("TokenRetriever must be set", nameof(TokenRetriever))
	//}
}


