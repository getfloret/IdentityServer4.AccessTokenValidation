package IdentityServer4_AccessTokenValidation

import (
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/IdentityModel/oidc"
	"github.com/karlseguin/ccache"
	"time"
)

type referenceHandler struct {
	cache *ccache.Cache //LRU Cache
	cacheTTL time.Duration
}

func New(cacheMaxSize int64,cacheTTL time.Duration) *referenceHandler{
	cache:= ccache.New(ccache.Configure().MaxSize(cacheMaxSize))
	return &referenceHandler{cache:cache,cacheTTL:cacheTTL}
}

func (h *referenceHandler) ParseReference(token string) map[string]string{
	item := h.cache.Get(token)//todo user hash
	if item != nil {
		if !item.Expired() {
			return
		}
	}

	oidc.Post(oidc.DefaultKeyLoader.OIDCConf().IntrospectionEndpoint,&oidc.TokenIntrospectionRequest{token})
	if !err && h.cacheTTL>0{
		h.cache.Set(token, ,h.cacheTTL)
	}
}