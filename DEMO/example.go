package main

import (
	AccessTokenValidation "github.com/getfloret/IdentityServer4.AccessTokenValidation"
	"github.com/getfloret/IdentityServer4.AccessTokenValidation/options"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(AccessTokenValidation.IdentityServerAuthentication(func(options *options.IdentityServerAuthenticationOptions)(){
		options.Authority = "https://u.highyouth.com"
		options.DiscoveryDocumentRefreshInterval = 30 * time.Second
	}))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":5200") // listen and serve on 0.0.0.0:8080
}