package handler

import (
	"github.com/YuStep/wbl0/internal/models"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

type getAllOrdersResponse struct {
	Data  []*models.Order `json:"data"`
	Count int             `json:"count"`
}

type getOrderByIdResponse struct {
	Order *models.Order `json:"data"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
