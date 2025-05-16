package handlers

import (
	"context"
	"net/http"

	"educabot.com/bookshop/services"
	"github.com/gin-gonic/gin"
)

type GetMetricsRequest struct {
	Author string `form:"author"`
}

type MetricsHandler struct {
	metricsService services.MetricsService
}

func NewMetricsHandler(metricsService services.MetricsService) MetricsHandler {
	return MetricsHandler{metricsService}
}

func (h MetricsHandler) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var query GetMetricsRequest
		err := ctx.ShouldBindQuery(&query)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid query parameters"})
			return
		}

		metrics, err := h.metricsService.GetMetrics(context.Background(), query.Author)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve metrics"})
			return
		}

		ctx.JSON(http.StatusOK, metrics)
	}
}
