package Middlewares

import (
	"GoMicro-Project/ServiceImpl"
	"github.com/gin-gonic/gin"
)

func InitMiddleware(prodService ServiceImpl.ProdService1) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Keys = make(map[string]interface{})
		ctx.Keys["prodService"] = prodService
		ctx.Next()
	}
}
