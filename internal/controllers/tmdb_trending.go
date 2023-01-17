package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTrending(context *gin.Context) {
	var result interface{}

	context.JSON(http.StatusOK, result)
}
