package middleware

import (
	"investify/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// check for valid admin token
func JWTOwnerAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := util.ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		error := util.ValidateOwnerRoleJWT(context)
		if error != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Only Business Owner is allowed to perform this action"})
			context.Abort()
			return
		}
		context.Next()
	}
}

// check for valid customer token
func JWTAuthInvestor() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := util.ValidateJWT(context)

		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		error := util.ValidateInvestorRoleJWT(context)
		if error != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Only  Investors are allowed to perform this action"})
			context.Abort()
			return
		}
		context.Next()
	}
}

func JWTAuthAny() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := util.ValidateJWT((context))

		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}

		context.Next()
	}
}
