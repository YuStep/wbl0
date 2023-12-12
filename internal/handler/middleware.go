package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	orderId = "id"
)

func getUserId(c *gin.Context) (string, error) {
	id, ok := c.Get(orderId)
	if !ok {
		return "", errors.New("user id not found")
	}

	idString, ok := id.(string)
	if !ok {
		return "", errors.New("user id is of invalid type")
	}

	return idString, nil
}
