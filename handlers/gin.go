package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorMessage struct {
	Error string `json:"error"`
}

func GinHandlerWithError(handler func(*gin.Context) (any, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := handler(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, &ErrorMessage{
				Error: err.Error(),
			})
		} else {
			ctx.AbortWithStatusJSON(http.StatusOK, data)
		}
		return
	}
}
