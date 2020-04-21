package Handlers

import (
	"GoMicro-Project/TestService"
	"github.com/gin-gonic/gin"
)

func GetProdsList(ginCtx *gin.Context) {
	ginCtx.JSON(200, TestService.NewProdList(5))
}

func GetUser(ginCtx *gin.Context) {
	ginCtx.String(200, "users api")
}
