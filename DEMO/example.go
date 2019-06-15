package main

import (
	IS4 "github.com/getfloret/IdentityServer4.AccessTokenValidation"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/options"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(IS4.IdentityServerAuthentication(func(options *options.IdentityServerAuthenticationOptions)(){
		options.Authority = "https://u.highyouth.com"
		options.DiscoveryDocumentRefreshInterval = 24 * time.Hour
		options.ApiName = "floret"
		options.ApiSecret = "getfloret.com"
	}))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":5200") // listen and serve on 0.0.0.0:8080
}