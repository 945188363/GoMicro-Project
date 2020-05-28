package Handlers

import (
	"GoMicro-Project/Core"
	"GoMicro-Project/TestService"
	"github.com/gin-gonic/gin"
	"time"
)

func GetProdsList(ginCtx *gin.Context) {
	time.Sleep(time.Duration(1) * time.Second)
	var msg Core.Message
	msg.Data = make(map[string]interface{})
	msg.Data["prod"] = TestService.NewProdList(5)
	ginCtx.JSON(200, msg)
}

func GetUser(ginCtx *gin.Context) {
	ginCtx.String(200, "users api")
}
