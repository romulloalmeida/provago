package web

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type response struct {
	Data interface{} `json:"data"`
}

func Failed(ctx *gin.Context, statusCode int, status, message string) {
	ctx.JSON(statusCode, errorResponse{
		Code:    statusCode,
		Status:  status,
		Message: message,
	})
}

func Success(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, response{Data: data})
}
