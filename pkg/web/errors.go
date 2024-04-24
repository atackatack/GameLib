package web

import "github.com/gin-gonic/gin"

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func ExistResponse() gin.H {
	return gin.H{"status": "exist"}
}
