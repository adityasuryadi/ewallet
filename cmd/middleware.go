package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/adityasuryadi/ewallet/helpers"
	"github.com/gin-gonic/gin"
)

func (d *Dependency) MiddlewareValidateAuth(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		ctx.JSON(400, "token needed")
		ctx.Abort()
		return
	}

	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		ctx.Abort()
		return
	}

	// Check if the header has the Bearer prefix
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
		ctx.Abort()
		return
	}

	// Extract the token
	token := parts[1]
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
		ctx.Abort()
		return
	}

	_, err := d.UserRepository.GetUserSessionByToken(ctx, token)
	if err != nil {
		ctx.JSON(401, "unauthorized")
		ctx.Abort()
		return
	}

	claim, err := helpers.ValidateToken(ctx, token)
	fmt.Println("claim", err)
	if err != nil {
		ctx.JSON(401, "unauthorized")
		ctx.Abort()
		return
	}

	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	ctx.Set("claim", claim)
	ctx.Set("token", token)
	ctx.Next()
}

func (d *Dependency) MiddlewareRefreshToken(ctx *gin.Context) {

	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		log.Println("authorization empty")
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	_, err := d.UserRepository.GetUserSessionByRefreshToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println("failed to get user session on DB: ", err)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	claim, err := helpers.ValidateToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println(err)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("jwt token is expired: ", claim.ExpiresAt)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	ctx.Set("token", claim)

	ctx.Next()
}
