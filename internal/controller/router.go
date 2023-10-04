package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/langgenius/dify-sandbox/internal/middleware"
	"github.com/langgenius/dify-sandbox/internal/static"
	"net/http"
)

func Setup(Router *gin.Engine) {
	PublicGroup := Router.Group("")
	Privat