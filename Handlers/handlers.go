package Handlers

import (
	"GoMicro-Project/Core"
	"GoMicro-Project/TestService"
	"github.com/gin-gonic/gin"
)

func GetProdsList(ginCtx *gin.Context) {
	var msg Core.Message
	msg.Data = make(map[string]interface{})
	msg.Data["prod"] = TestService.NewProdList(5)
	ginCtx.JSON(200, msg)
}

func GetUser(ginCtx *gin.Context) {
	ginCtx.String(200, "users api")
}
