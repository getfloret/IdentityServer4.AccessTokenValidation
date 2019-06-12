package oidc

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/IdentityModel"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/IdentityModel/jwk"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/infrastructure"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/infrastructure/breaker"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/infrastructure/caching"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/options"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

// http://openid.net/specs/openid-connect-discovery-1_0.html#ProviderConfig
// https://planb-provider.example.org/.well-known/openid-configuration
// https://accounts.google.com/.well-known/openid-configuration
type cachingOpenIDProviderLoader struct {
	url      string
	keyCache *caching.Cache
}

var (
	errInvalidResponseStatusCode = errors.New("Invalid response status code")
	scheduleFunc                 = infrastructure.Schedule
)

// NewCachingOpenIDProviderLoader returns a KeyLoader that uses the configured URL to an OpenID
// endpoint where the URI for the JSON Web Keys Set is available
func NewCachingOpenIDProviderLoader(u *url.URL) KeyLoader {
	hystrix.ConfigureCommand("loadConfiguration", hystrix.CommandConfig{
		Timeout:               10000,
		MaxConcurrentRequests: 2,
		ErrorPercentThreshold: 25,
	})
	hystrix.ConfigureCommand("loadKeys", hystrix.CommandConfig{
		Timeout:               10000,
		MaxConcurrentRequests: 2,
		ErrorPercentThreshold: 25,
	})

	kl := &cachingOpenIDProviderLoader{url: u.String(), keyCache: caching.NewCache()}
	scheduleFunc(options.GlobalAuthenticationOptions.DiscoveryDocumentRefreshInterval, kl.refreshKeys)
	return kl
}

func (kl *cachingOpenIDProviderLoader) LoadKey(id string) (interface{}, error) {
	v := kl.keyCache.Get(id)
	if v == nil {
		return nil, fmt.Errorf("Key '%s' not found", id)
	}
	return v.(jwk.JSONWebKey).Key, nil
}

func (kl *cachingOpenIDProviderLoader) Keys() map[string]interface{} {
	return kl.keyCache.Snapshot()
}

// Example: https://www.googleapis.com/oauth2/v3/certs
func (kl *cachingOpenIDProviderLoader) refreshKeys() {
	log.Println("Refreshing keys..")

	log.Println("Loading configuration..")
	c, err := kl.loadConfiguration()
	if err != nil {
		log.Printf("Failed to get configuration from %q. %s\n", kl.url, err)
		return
	}

	log.Println("Configuration loaded successfully, loading JWKS..")
	resp, err := breaker.Get("loadKeys", c.JwksUri)
	if err != nil {
		log.Println("Failed to get JWKS from ", c.JwksUri)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read JWKS response body from %q: %v\n", c.JwksUri, err)
		return
	}

	log.Println("JWKS loaded successfully, parsing JWKS..")
	jwks := new(jwk.JSONWebKeySet)
	if err = json.Unmarshal(body, jwks); err != nil {
		log.Println("Failed to parse JWKS: ", err)
		return
	}

	// safety first: only remove public keys if our newly
	// received list contains at least one public key!
	// (we don't want our tokeninfo to run out of public keys
	// just because somebody cleared the provider database)
	numKeys := len(jwks.Keys)
	if numKeys < 1 {
		log.Println("No JWKS currently in the OpenID provider")
		//if c, ok := metrics.DefaultRegistry.GetOrRegister(metricsNoKeysError, metrics.NewCounter).(metrics.Counter); ok {
		//	c.Inc(1)
		//}
		return
	}

	//if g, ok := metrics.DefaultRegistry.GetOrRegister(metricsNumKeys, metrics.NewGauge).(metrics.Gauge); ok {
	//	g.Update(int64(numKeys))
	//}

	newKeys := jwks.ToMap()
	for kid, k := range newKeys {
		key := k.(jwk.JSONWebKey)
		existing := kl.keyCache.Get(kid)
		if existing == nil {
			log.Printf("Received new public key %q (%s)\n", kid, key.Algorithm)
		} else if !reflect.DeepEqual(existing, key) {
			// this is potentially dangerous: the key contents changed..
			// (but maybe the key wasn't used for signing yet, so it might be ok)
			log.Printf("Received a replacement public key for existing key %q (%s)", kid, key.Algorithm)
		}
	}

	log.Printf("Resetting key cache with %d key(s)..", numKeys)
	kl.keyCache.Reset(newKeys)
	log.Println("Refresh done..")
}

// https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderConfigurationResponse
func (kl *cachingOpenIDProviderLoader) loadConfiguration() (*IdentityModel.OpenIdConnectConfiguration, error) {
	resp, err := breaker.Get("loadConfiguration", kl.url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errInvalidResponseStatusCode
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	config := new(IdentityModel.OpenIdConnectConfiguration)
	err = json.Unmarshal(body, config)
	return config, err
}
