package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

const (
	success = 900
	failed  = 999
)

func responseSuccessWithData(data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &Response{
		Code:    success,
		Data:    data,
		Message: "",
	})
}

func responseFailedWithMessage(msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &Response{
		Code:    failed,
		Data:    nil,
		Message: msg,
	})
}
