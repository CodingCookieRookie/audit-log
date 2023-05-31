package handlers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func NewStackTracedError(err error) *StackTracedError {
	return &StackTracedError{
		err:   err,
		stack: debug.Stack(),
	}
}

func NewStackTracedErrorf(format string, args ...interface{}) *StackTracedError {
	return &StackTracedError{
		err:   fmt.Errorf(format, args...),
		stack: debug.Stack(),
	}
}

type StackTracedError struct {
	err   error
	stack []byte
}

func (ste *StackTracedError) Error() string {
	return ste.err.Error() + "\n" + string(ste.stack)
}

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
