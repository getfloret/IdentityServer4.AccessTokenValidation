package options

import (
	"time"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/IdentityModel"
)
// Options class provides information needed to control Bearer Authentication handler behavior
type JwtBearerOptions struct {
	/// Gets or sets if HTTPS is required for the metadata address or authority.
	/// The default is true. This should be disabled only in development environments.
	RequireHttpsMetadata bool


	/// Gets or sets the discovery endpoint for obtaining metadata
	MetadataAddress string

	/// Gets or sets the Authority to use when making OpenIdConnect calls.
	Authority string

	/// Gets or sets the audience for any received OpenIdConnect token.
	/// The expected audience for any received OpenIdConnect token.
    Audience string

	/// Gets or sets the challenge to put in the "WWW-Authenticate" header.
	Challenge string

	/// <summary>
	/// The object provided by the application to process events raised by the bearer authentication handler.
	/// The application may implement the interface fully, or it may create an instance of JwtBearerEvents
	/// and assign delegates only to the events it wants to process.
	/// </summary>
//	public new JwtBearerEvents Events
//{
//	get {
//	return (JwtBearerEvents)base.Events;
//}
//	set {
//	base.Events = value;
//}
//}

	/// <summary>
	/// The HttpMessageHandler used to retrieve metadata.
	/// This cannot be set at the same time as BackchannelCertificateValidator unless the value
	/// is a WebRequestHandler.
	/// </summary>
	//public HttpMessageHandler BackchannelHttpHandler {
	//get;
//	//set;
//}

	/// Gets or sets the timeout when using the backchannel to make an http call.
	BackchannelTimeout time.Duration

	/// Configuration provided directly by the developer. If provided, then MetadataAddress and the Backchannel properties
	/// will not be used. This information should not be updated during request processing.
	Configuration IdentityModel.OpenIdConnectConfiguration

	//https://github.com/aspnet/Security/blob/beaa2b443d46ef8adaf5c2a89eb475e1893037c2/src/Microsoft.AspNetCore.Authentication.OpenIdConnect/OpenIdConnectOptions.cs
	/// <summary>
	/// Responsible for retrieving, caching, and refreshing the configuration from metadata.
	/// If not provided, then one will be created using the MetadataAddress and Backchannel properties.
	/// </summary>
//	public IConfigurationManager<OpenIdConnectConfiguration> ConfigurationManager {
//	get;
//	set;
//}

	/// Gets or sets if a metadata refresh should be attempted after a SecurityTokenSignatureKeyNotFoundException. This allows for automatic
	/// recovery in the event of a signature key rollover. This is enabled by default.
	RefreshOnIssuerKeyNotFound bool

	/// <summary>
	/// Gets the ordered list of <see cref="ISecurityTokenValidator"/> used to validate access tokens.
	/// </summary>
//	public IList<ISecurityTokenValidator> SecurityTokenValidators {
//	get;
//} = new List<ISecurityTokenValidator> {
//	new JwtSecurityTokenHandler()
//};

	/// <summary>
	/// Gets or sets the parameters used to validate identity tokens.
	/// </summary>
	/// <remarks>Contains the types and definitions required for validating a token.</remarks>
	/// <exception cref="ArgumentNullException">if 'value' is null.</exception>
//	public TokenValidationParameters TokenValidationParameters {
//	get;
//	set;
//} = new TokenValidationParameters();

	/// Defines whether the bearer token should be stored in the
	/// <see cref="Http.Authentication.AuthenticationProperties"/> after a successful authorization.
	SaveToken bool

	/// Defines whether the token validation errors should be returned to the caller.
	/// Enabled by default, this option can be disabled to prevent the JWT handler
	/// from returning an error and an error_description in the WWW-Authenticate header.
	IncludeErrorDetails bool
}
