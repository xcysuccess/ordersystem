package service

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Index TODO
func Index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "hello gin " + strings.ToLower(context.Request.Method) + " method",
	})
}