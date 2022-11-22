package http

import (
	"fintech-api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthorizeJWT : to authorize JWT Token
func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userName := ctx.Param("username")
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No Authorization header found"})
			return
		}
		token := strings.Split(authHeader, " ")[1]
		if err, ok := utils.ValidateToken(token, userName); !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Sprintf("Authorization error: %v", err)})
		}
	}
}

// Authorize : determines if current user has been authorized to take an action on an object.
//func (a *AuthedUser) Authorize(obj string, act string, enforcer *casbin.Enforcer) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// Get current user/subject
//		sub := a.User
//		if sub == "" {
//			c.AbortWithStatusJSON(401, gin.H{"msg": "User hasn't logged in yet"})
//			return
//		}
//
//		// Load policy from Database
//		err := enforcer.LoadPolicy()
//		if err != nil {
//			c.AbortWithStatusJSON(500, gin.H{"msg": "Failed to load policy from DB"})
//			return
//		}
//
//		// Casbin enforces policy
//		ok, err := enforcer.Enforce(sub, obj, act)
//
//		if err != nil {
//			c.AbortWithStatusJSON(500, gin.H{"msg": "Error occurred when authorizing user"})
//			return
//		}
//
//		if !ok {
//			c.AbortWithStatusJSON(403, gin.H{"msg": "You are not authorized"})
//			return
//		}
//		c.Next()
//	}
//}
