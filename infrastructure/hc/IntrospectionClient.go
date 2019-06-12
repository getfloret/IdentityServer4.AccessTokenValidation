package hc

import (
	"bytes"
	"encoding/json"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/options"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

https://github.com/IdentityModel/IdentityModel/blob/master/src/Client/IntrospectionClient.cs
import (
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/options"
	"net"
	"net/http"
	"time"
)

var (
	// Default global instance of a custom http.Client using the defaults from the options package
	IntrospectionClient = DefaultIntrospectionHTTPClient()
	// UserAgent can be used to specify the User-Agent header sent on every request that used this package's
	// http.Client
	IntrospectionClientUserAgent = "IS4.IntrospectionClient"
)

// DefaultHTTPClient returns a new http.Client with KeepAlive disabled. That means no connection pooling.
// Use it only for one time requests where performance is not a concern
// It use some settings from the options package: options.HttpClientTimeout and options.HttpClientTlsTimeout
func DefaultIntrospectionHTTPClient() *http.Client {
	return NewIntrospectionClient(options.GlobalAuthenticationOptions.BackChannelTimeouts, options.GlobalAuthenticationOptions.BackChannelTimeouts)
}

// NewHTTPClient returns a new http.Client with specific timeouts from its arguments. KeepAlive is disabled.
// That means no connection pooling. Use it only for one time requests where performance is not a concern
func NewIntrospectionClient(timeout time.Duration, tlsTimeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			Proxy:               http.ProxyFromEnvironment,
			DisableKeepAlives:   true,
			Dial:                (&net.Dialer{Timeout: timeout}).Dial,
			TLSHandshakeTimeout: tlsTimeout}}
}

type TokenIntrospectionRequest struct {
	Token string //token
	TokenTypeHint string //token_type_hint
}

func Post(url string, data TokenIntrospectionRequest, contentType string) (content string) {
	bodyData := "token="+data.Token
	if(data.TokenTypeHint!=""){
		bodyData+="&token_type_hint="+data.TokenTypeHint
	}
	reqest,err := http.NewRequest("POST",my_url,strings.NewReader(bodyData))
	reqest.Header.Set("Content-Type","application/x-www-form-urlencoded")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(reqest)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	return
}






func extractIntrospectResult(r io.Reader) (*Result, error) {
	res := Result{
		Optionals: make(map[string]json.RawMessage),
	}

	if err := json.NewDecoder(r).Decode(&res.Optionals); err != nil {
		return nil, err
	}

	if val, ok := res.Optionals["active"]; ok {
		if err := json.Unmarshal(val, &res.Active); err != nil {
			return nil, err
		}

		delete(res.Optionals, "active")
	}

	return &res, nil
}

// Result is the OAuth2 Introspection Result
type Result struct {
	Active bool

	Optionals map[string]json.RawMessage
}




/// Models an OAuth 2.0 introspection response
public class TokenIntrospectionResponse : ProtocolResponse
{
/// Allows to initialize instance specific data.
protected override Task InitializeAsync(object initializationData = null)
{
if (!IsError)
{
var claims = Json.ToClaims(excludeKeys: "scope").ToList();

// due to a bug in identityserver - we need to be able to deal with the scope list both in array as well as space-separated list format
var scope = Json.TryGetValue("scope");

// scope element exists
if (scope != null)
{
// it's an array
if (scope is JArray scopeArray)
{
foreach (var item in scopeArray)
{
claims.Add(new Claim("scope", item.ToString()));
}
}
else
{
// it's a string
var scopeString = scope.ToString();

var scopes = scopeString.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);
foreach (var scopeValue in scopes)
{
claims.Add(new Claim("scope", scopeValue));
}
}
}

Claims = claims;
}
else
{
Claims = Enumerable.Empty<Claim>();
}

return Task.CompletedTask;
}

/// Gets a value indicating whether the token is active.
public bool IsActive => Json.TryGetBoolean("active").Value;

/// Gets the claims.
public IEnumerable<Claim> Claims { get; protected set; }

}



/// Request for OAuth token introspection
public class TokenIntrospectionRequest : ProtocolRequest
{
/// Gets or sets the token.
public string Token { get; set; }

/// Gets or sets the token type hint.
public string TokenTypeHint { get; set; }
}




