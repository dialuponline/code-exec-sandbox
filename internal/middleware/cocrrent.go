package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/langgenius/dify-sandbox/internal/types"
	"github.com/langgenius/dify-sandbox/internal/utils/log"
)

func MaxWorker(max int) gin.HandlerFunc {
	log.Info("setting max workers to %d", max)
	sem := make(chan struct{}, max)

	return func(c *gin.Context) {
		sem <- struct{}{}
		defer func() {
			<-sem
		}