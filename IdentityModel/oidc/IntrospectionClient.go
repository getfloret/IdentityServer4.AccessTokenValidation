package oidc

import (
	"encoding/json"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/options"
	"io"
	"net"
	"net/http"
	"strings"
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

func Post(url string, data *TokenIntrospectionRequest) (tokenResult *IntrospectionResult, err error) {
	bodyData := "token="+data.Token
	if(data.TokenTypeHint!=""){
		bodyData+="&token_type_hint="+data.TokenTypeHint
	}

	req, err := http.NewRequest("POST",DefaultKeyLoader.OIDCConf().IntrospectionEndpoint,strings.NewReader(bodyData))
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	tokenResult, _ = extractIntrospectResult(resp.Body)
	return
}




func extractIntrospectResult(r io.Reader) (*IntrospectionResult, error) {
	res := IntrospectionResult{
		Claims: make(map[string]interface{}),
	}
	if err := json.NewDecoder(r).Decode(&res.Claims); err != nil {
		return nil, err
	}

	if val, ok := res.Claims["active"]; ok {
		if err := json.Unmarshal(val.([]byte), &res.Active); err != nil {
			return nil, err
		}
		delete(res.Claims, "active")
	}

	return &res, nil
}

// IntrospectionResult is the OAuth2 Introspection IntrospectionResult
type IntrospectionResult struct {
	Active bool

	Claims map[string]interface{}
}


