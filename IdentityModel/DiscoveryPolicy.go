package IdentityModel

type DiscoveryPolicy struct {
	// Gets or sets the Authority on which the policy checks will be based on
	Authority string

	//// Method of comparison for issuer and authority names. Defaults to <see cref="StringComparison.Ordinal" />
	//public StringComparison AuthorityNameComparison { get; set; } = StringComparison.Ordinal;

	// Specifies if HTTPS is enforced on all endpoints. Defaults to true.
	RequireHttps bool

	// Specifies if HTTP is allowed on loopback addresses. Defaults to true.
	AllowHttpOnLoopback bool

	///// <summary>
	///// Specifies valid loopback addresses, defaults to localhost and 127.0.0.1
	///// </summary>
	//public ICollection<string> LoopbackAddresses = new HashSet<string> { "localhost", "127.0.0.1" };

	// Specifies if the issuer name is checked to be identical to the authority. Defaults to true.
	ValidateIssuerName bool

	// Specifies if all endpoints are checked to belong to the authority. Defaults to true.
	ValidateEndpoints bool

	// Specifies a list of endpoints that should be excluded from validation
	EndpointValidationExcludeList []string

	// Specifies a list of additional base addresses that should be allowed for endpoints
	AdditionalEndpointBaseAddresses []string

	// Specifies if a key set is required. Defaults to true.
	RequireKeySet bool
}
