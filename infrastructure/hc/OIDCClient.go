package hc

import (
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/options"
	"net"
	"net/http"
	"time"
)

var (
	// Default global instance of a custom http.Client using the defaults from the options package
	OIDCDiscoveryClient = DefaultOIDCHTTPClient()
	// UserAgent can be used to specify the User-Agent header sent on every request that used this package's
	// http.Client
	UserAgent = "IS4.AccessTokenValidation"
)

// DefaultHTTPClient returns a new http.Client with KeepAlive disabled. That means no connection pooling.
// Use it only for one time requests where performance is not a concern
// It use some settings from the options package: options.HttpClientTimeout and options.HttpClientTlsTimeout
func DefaultOIDCHTTPClient() *http.Client {
	return NewHTTPClient(options.GlobalAuthenticationOptions.DiscoveryHttpClientTimeout, options.GlobalAuthenticationOptions.DiscoveryHttpClientTLSTimeout)
}

// NewHTTPClient returns a new http.Client with specific timeouts from its arguments. KeepAlive is disabled.
// That means no connection pooling. Use it only for one time requests where performance is not a concern
func NewHTTPClient(timeout time.Duration, tlsTimeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			Proxy:               http.ProxyFromEnvironment,
			DisableKeepAlives:   true,
			Dial:                (&net.Dialer{Timeout: timeout}).Dial,
			TLSHandshakeTimeout: tlsTimeout}}
}

// Get issues a GET to the specified URL. It follows redirects, up to a maximum of 10
//
// An error is returned if there were too many redirects or if there
// was an HTTP protocol error. A non-2xx response doesn't cause an
// error.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
//
// Get is a wrapper around Default.Get.
//
// A User-Agent header is set for every request from the UserAgent variable in the same package
func Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	return OIDCDiscoveryClient.Do(req)
}
