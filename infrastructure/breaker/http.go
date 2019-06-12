package breaker

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/infrastructure/hc"
)

// Get will fetch the HTTP resource from url using a GET method, wrapped in a circuit breaker named name
func Get(name string, url string) (*http.Response, error) {
	return GetWithFallback(name, url, nil)
}

// GetWithFallback will fetch the HTTP resource from url using a GET method, wrapped in a circuit breaker named name.
// If the operation fails, the fallback function f is called with the previous error as an argument
func GetWithFallback(name string, url string, f func(error) error) (resp *http.Response, err error) {
	err = hystrix.Do(name, func() error {
		var internalError error
		if resp, internalError = hc.OIDCDiscoveryClient.Get(url); internalError == nil {
			log.Trace("breaker success: ",name)
		} else {
			log.Warn("breaker failure: ",name)
		}
		return internalError
	}, f)
	return
}
