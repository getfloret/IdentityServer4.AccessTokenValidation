package IdentityModel

import "github.com/getfloret/IdentityServer4.AccessTokenValidation/IdentityModel/jwk"
// https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata
// Contains OpenIdConnect configuration that can be populated from a json string.
type OpenIdConnectConfiguration struct {
	//Issuer  string `json:"issuer"`
	//JwksURI string `json:"jwks_uri"`
	///// <summary>
	///// Deserializes the json string into an <see cref="T:Microsoft.IdentityModel.Protocols.OpenIdConnect.OpenIdConnectConfiguration" /> object.
	///// </summary>
	///// <param name="json">json string representing the configuration.</param>
	///// <returns><see cref="T:Microsoft.IdentityModel.Protocols.OpenIdConnect.OpenIdConnectConfiguration" /> object representing the configuration.</returns>
	///// <exception cref="T:System.ArgumentNullException">If 'json' is null or empty.</exception>
	///// <exception cref="T:System.ArgumentException">If 'json' fails to deserialize.</exception>
	//public static OpenIdConnectConfiguration Create(string json)
	//{
	//if (string.IsNullOrEmpty(json))
	//throw LogHelper.LogArgumentNullException(nameof (json));
	//LogHelper.LogVerbose("IDX21808: Deserializing json into OpenIdConnectConfiguration object: '{0}'.", (object) json);
	//return new OpenIdConnectConfiguration(json);
	//}
	//
	///// <summary>
	///// Serializes the <see cref="T:Microsoft.IdentityModel.Protocols.OpenIdConnect.OpenIdConnectConfiguration" /> object to a json string.
	///// </summary>
	///// <param name="configuration"><see cref="T:Microsoft.IdentityModel.Protocols.OpenIdConnect.OpenIdConnectConfiguration" /> object to serialize.</param>
	///// <returns>json string representing the configuration object.</returns>
	///// <exception cref="T:System.ArgumentNullException">If 'configuration' is null.</exception>
	//public static string Write(OpenIdConnectConfiguration configuration)
	//{
	//if (configuration == null)
	//throw LogHelper.LogArgumentNullException(nameof (configuration));
	//LogHelper.LogVerbose("IDX21809: Serializing OpenIdConfiguration object to json string.");
	//return JsonConvert.SerializeObject((object) configuration);
	//}

	///// <summary>
	///// Initializes an new instance of <see cref="T:Microsoft.IdentityModel.Protocols.OpenIdConnect.OpenIdConnectConfiguration" /> from a json string.
	///// </summary>
	///// <param name="json">a json string containing the metadata</param>
	///// <exception cref="T:System.ArgumentNullException">If 'json' is null or empty.</exception>
	//public OpenIdConnectConfiguration(string json)
	//{
	//if (string.IsNullOrEmpty(json))
	//throw LogHelper.LogArgumentNullException(nameof (json));
	//try
	//{
	//LogHelper.LogVerbose("IDX21806: Deserializing json string into json web keys.", (object) json, (object) this);
	//JsonConvert.PopulateObject(json, (object) this);
	//}
	//catch (Exception ex)
	//{
	//throw LogHelper.LogExceptionMessage((Exception) new ArgumentException(LogHelper.FormatInvariant("IDX21815: Error deserializing json: '{0}' into '{1}'.", (object) json, (object) this.GetType()), ex));
	//}
	//}

	// When deserializing from JSON any properties that are not defined will be placed here.
	AdditionalData map[string]interface{} `json:"-"`

	/// <summary>Gets the collection of 'acr_values_supported'</summary>
	AcrValuesSupported []string `json:"acr_values_supported,omitempty"`

	/// <summary>Gets or sets the 'authorization_endpoint'.</summary>
	AuthorizationEndpoint string `json:"authorization_endpoint,omitempty"`

	/// <summary>Gets or sets the 'check_session_iframe'.</summary>
	CheckSessionIframe string `json:"check_session_iframe,omitempty"`

	/// <summary>Gets the collection of 'claims_supported'</summary>
	ClaimsSupported []string `json:"claims_supported,omitempty"`

	/// <summary>Gets the collection of 'claims_locales_supported'</summary>
	ClaimsLocalesSupported []string `json:"claims_locales_supported,omitempty"`

	/// <summary>Gets or sets the 'claims_parameter_supported'</summary>
	ClaimsParameterSupported bool `json:"claims_parameter_supported,omitempty"`

	/// <summary>Gets the collection of 'claim_types_supported'</summary>
	ClaimTypesSupported []string `json:"claim_types_supported,omitempty"`

	/// <summary>Gets the collection of 'display_values_supported'</summary>
	DisplayValuesSupported []string `json:"display_values_supported,omitempty"`

	/// <summary>Gets or sets the 'end_session_endpoint'.</summary>
	EndSessionEndpoint string `json:"end_session_endpoint,omitempty"`

	/// Gets or sets the 'frontchannel_logout_session_supported'.
	FrontchannelLogoutSessionSupported bool `json:"frontchannel_logout_session_supported,omitempty"`

	/// <summary>Gets or sets the 'frontchannel_logout_supported'.</summary>
	FrontchannelLogoutSupported bool `json:"frontchannel_logout_supported,omitempty"`

	/// <summary>Gets the collection of 'grant_types_supported'</summary>
	GrantTypesSupported []string `json:"grant_types_supported,omitempty"`

	/// Boolean value specifying whether the OP supports HTTP-based logout. Default is false.
	HttpLogoutSupported bool `json:"http_logout_supported,omitempty"`

	/// Gets the collection of 'id_token_encryption_alg_values_supported'.
	IdTokenEncryptionAlgValuesSupported []string `json:"id_token_encryption_alg_values_supported,omitempty"`

	/// Gets the collection of 'id_token_encryption_enc_values_supported'.
	IdTokenEncryptionEncValuesSupported []string `json:"id_token_encryption_enc_values_supported,omitempty"`

	/// Gets the collection of 'id_token_signing_alg_values_supported'.
	IdTokenSigningAlgValuesSupported []string `json:"id_token_signing_alg_values_supported,omitempty"`

	/// <summary>Gets or sets the 'issuer'.</summary>
	Issuer string `json:"issuer,omitempty"`

	/// <summary>Gets or sets the 'jwks_uri'</summary>
	JwksUri string `json:"jwks_uri,omitempty"`

	/// Gets or sets the <see cref="P:Microsoft.IdentityModel.Protocols.OpenIdConnect.OpenIdConnectConfiguration.JsonWebKeySet" />
	JsonWebKeySet jwk.JSONWebKeySet

	/// Boolean value specifying whether the OP can pass a sid (session ID) query parameter to identify the RP session at the OP when the logout_uri is used. Dafault Value is false.
	LogoutSessionSupported bool `json:"logout_session_supported,omitempty"`

	/// <summary>Gets or sets the 'op_policy_uri'</summary>
	OpPolicyUri string `json:"op_policy_uri,omitempty"`

	/// <summary>Gets or sets the 'op_tos_uri'</summary>
	OpTosUri string `json:"op_tos_uri,omitempty"`

	/// <summary>Gets or sets the 'registration_endpoint'</summary>
	RegistrationEndpoint string `json:"registration_endpoint,omitempty"`

	/// Gets the collection of 'request_object_encryption_alg_values_supported'.
	RequestObjectEncryptionAlgValuesSupported []string `json:"request_object_encryption_alg_values_supported,omitempty"`

	/// Gets the collection of 'request_object_encryption_enc_values_supported'.
	RequestObjectEncryptionEncValuesSupported []string `json:"request_object_encryption_enc_values_supported,omitempty"`

	/// Gets the collection of 'request_object_signing_alg_values_supported'.
	RequestObjectSigningAlgValuesSupported []string `json:"request_object_signing_alg_values_supported,omitempty"`

	/// <summary>Gets or sets the 'request_parameter_supported'</summary>
	RequestParameterSupported bool `json:"request_parameter_supported,omitempty"`

	/// <summary>Gets or sets the 'request_uri_parameter_supported'</summary>
	RequestUriParameterSupported bool `json:"request_uri_parameter_supported,omitempty"`

	/// <summary>Gets or sets the 'require_request_uri_registration'</summary>
	RequireRequestUriRegistration bool `json:"require_request_uri_registration,omitempty"`

	/// <summary>Gets the collection of 'response_modes_supported'.</summary>
	ResponseModesSupported []string `json:"response_modes_supported,omitempty"`

	/// <summary>Gets the collection of 'response_types_supported'.</summary>
	ResponseTypesSupported []string `json:"response_types_supported,omitempty"`

	/// <summary>Gets or sets the 'service_documentation'</summary>
	ServiceDocumentation string `json:"service_documentation,omitempty"`

	/// <summary>Gets the collection of 'scopes_supported'</summary>
	ScopesSupported []string `json:"scopes_supported,omitempty"`

	///// Gets the <see cref="T:System.Collections.Generic.ICollection`1" /> that the IdentityProvider indicates are to be used signing tokens.
	//
	//public ICollection<SecurityKey> SigningKeys { get; } = (ICollection<SecurityKey>) new Collection<SecurityKey>();

	/// <summary>Gets the collection of 'subject_types_supported'.</summary>
	SubjectTypesSupported []string `json:"subject_types_supported,omitempty"`

	/// <summary>Gets or sets the 'token_endpoint'.</summary>
	TokenEndpoint string `json:"token_endpoint,omitempty"`

	/// Gets the collection of 'token_endpoint_auth_methods_supported'.
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported,omitempty"`

	/// <summary>
	/// Gets the collection of 'token_endpoint_auth_signing_alg_values_supported'.
	/// </summary>
	TokenEndpointAuthSigningAlgValuesSupported []string `json:"token_endpoint_auth_signing_alg_values_supported,omitempty"`

	/// <summary>Gets the collection of 'ui_locales_supported'</summary>
	UILocalesSupported []string `json:"ui_locales_supported,omitempty"`

	/// <summary>Gets or sets the 'user_info_endpoint'.</summary>
	UserInfoEndpoint string `json:"userinfo_endpoint,omitempty"`

	/// Gets the collection of 'userinfo_encryption_alg_values_supported'
	UserInfoEndpointEncryptionAlgValuesSupported []string `json:"userinfo_encryption_alg_values_supported,omitempty"`

	/// Gets the collection of 'userinfo_encryption_enc_values_supported'
	UserInfoEndpointEncryptionEncValuesSupported []string `json:"userinfo_encryption_enc_values_supported,omitempty"`

	/// Gets the collection of 'userinfo_signing_alg_values_supported'
	UserInfoEndpointSigningAlgValuesSupported []string `json:"userinfo_signing_alg_values_supported,omitempty"`

	IntrospectionEndpoint string `json:"introspection_endpoint,omitempty"`
}

///// <summary>
///// Gets a bool that determines if the 'acr_values_supported' (AcrValuesSupported) property should be serialized.
///// This is used by Json.NET in order to conditionally serialize properties.
///// </summary>
///// <return>true if 'acr_values_supported' (AcrValuesSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeAcrValuesSupported()
//{
//return this.AcrValuesSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'claims_supported' (ClaimsSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'claims_supported' (ClaimsSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeClaimsSupported()
//{
//return this.ClaimsSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'claims_locales_supported' (ClaimsLocalesSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'claims_locales_supported' (ClaimsLocalesSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeClaimsLocalesSupported()
//{
//return this.ClaimsLocalesSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'claim_types_supported' (ClaimTypesSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'claim_types_supported' (ClaimTypesSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeClaimTypesSupported()
//{
//return this.ClaimTypesSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'display_values_supported' (DisplayValuesSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'display_values_supported' (DisplayValuesSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeDisplayValuesSupported()
//{
//return this.DisplayValuesSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'grant_types_supported' (GrantTypesSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'grant_types_supported' (GrantTypesSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeGrantTypesSupported()
//{
//return this.GrantTypesSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'id_token_encryption_alg_values_supported' (IdTokenEncryptionAlgValuesSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'id_token_encryption_alg_values_supported' (IdTokenEncryptionAlgValuesSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeIdTokenEncryptionAlgValuesSupported()
//{
//return this.IdTokenEncryptionAlgValuesSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'id_token_encryption_enc_values_supported' (IdTokenEncryptionEncValuesSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'id_token_encryption_enc_values_supported' (IdTokenEncryptionEncValuesSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeIdTokenEncryptionEncValuesSupported()
//{
//return this.IdTokenEncryptionEncValuesSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'id_token_signing_alg_values_supported' (IdTokenSigningAlgValuesSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'id_token_signing_alg_values_supported' (IdTokenSigningAlgValuesSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeIdTokenSigningAlgValuesSupported()
//{
//return this.IdTokenSigningAlgValuesSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'request_object_encryption_alg_values_supported' (RequestObjectEncryptionAlgValuesSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'request_object_encryption_alg_values_supported' (RequestObjectEncryptionAlgValuesSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeRequestObjectEncryptionAlgValuesSupported()
//{
//return this.RequestObjectEncryptionAlgValuesSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'request_object_encryption_enc_values_supported' (RequestObjectEncryptionEncValuesSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'request_object_encryption_enc_values_supported' (RequestObjectEncryptionEncValuesSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeRequestObjectEncryptionEncValuesSupported()
//{
//return this.RequestObjectEncryptionEncValuesSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'request_object_signing_alg_values_supported' (RequestObjectSigningAlgValuesSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'request_object_signing_alg_values_supported' (RequestObjectSigningAlgValuesSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeRequestObjectSigningAlgValuesSupported()
//{
//return this.RequestObjectSigningAlgValuesSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'response_modes_supported' (ResponseModesSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'response_modes_supported' (ResponseModesSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeResponseModesSupported()
//{
//return this.ResponseModesSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'response_types_supported' (ResponseTypesSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'response_types_supported' (ResponseTypesSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeResponseTypesSupported()
//{
//return this.ResponseTypesSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'scopes_supported' (ScopesSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'scopes_supported' (ScopesSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeScopesSupported()
//{
//return this.ScopesSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'subject_types_supported' (SubjectTypesSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'subject_types_supported' (SubjectTypesSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeSubjectTypesSupported()
//{
//return this.SubjectTypesSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'token_endpoint_auth_methods_supported' (TokenEndpointAuthMethodsSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'token_endpoint_auth_methods_supported' (TokenEndpointAuthMethodsSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeTokenEndpointAuthMethodsSupported()
//{
//return this.TokenEndpointAuthMethodsSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'token_endpoint_auth_signing_alg_values_supported' (TokenEndpointAuthSigningAlgValuesSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'token_endpoint_auth_signing_alg_values_supported' (TokenEndpointAuthSigningAlgValuesSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeTokenEndpointAuthSigningAlgValuesSupported()
//{
//return this.TokenEndpointAuthSigningAlgValuesSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'ui_locales_supported' (UILocalesSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'ui_locales_supported' (UILocalesSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeUILocalesSupported()
//{
//return this.UILocalesSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'userinfo_encryption_alg_values_supported' (UserInfoEndpointEncryptionAlgValuesSupported ) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'userinfo_encryption_alg_values_supported' (UserInfoEndpointEncryptionAlgValuesSupported ) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeUserInfoEndpointEncryptionAlgValuesSupported()
//{
//return this.UserInfoEndpointEncryptionAlgValuesSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'userinfo_encryption_enc_values_supported' (UserInfoEndpointEncryptionEncValuesSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'userinfo_encryption_enc_values_supported' (UserInfoEndpointEncryptionEncValuesSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeUserInfoEndpointEncryptionEncValuesSupported()
//{
//return this.UserInfoEndpointEncryptionEncValuesSupported.Count > 0;
//}

/// <summary>
/// Gets a bool that determines if the 'userinfo_signing_alg_values_supported' (UserInfoEndpointSigningAlgValuesSupported) property should be serialized.
/// This is used by Json.NET in order to conditionally serialize properties.
/// </summary>
/// <return>true if 'userinfo_signing_alg_values_supported' (UserInfoEndpointSigningAlgValuesSupported) is not empty; otherwise, false.</return>
//[EditorBrowsable(EditorBrowsableState.Never)]
//public bool ShouldSerializeUserInfoEndpointSigningAlgValuesSupported()
//{
//return this.UserInfoEndpointSigningAlgValuesSupported.Count > 0;
//}}