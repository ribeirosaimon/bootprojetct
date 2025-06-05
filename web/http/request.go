package http

import "github.com/gin-gonic/gin"

func GetBody[T any](c *gin.Context) (*T, error) {
	var value T
	if err := c.ShouldBind(value); err != nil {
		return nil, err
	}
	return &value, nil
}
