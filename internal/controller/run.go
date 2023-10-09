package controller

import (
	"github.com/gin-gonic/gin"
	runner_types "github.com/langgenius/dify-sandbox/internal/core/runner/types"
	"github.com/langgenius/dify-sandbox/internal/service"
	"github.com/langgenius/dify-sandbox/internal/types"
)

func RunSandboxController(c *gin.Context) {
	BindRequest(c, func(req struct {
		Language      string `json:"langua