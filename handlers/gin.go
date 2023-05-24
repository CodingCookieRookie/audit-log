package handlers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/CodingCookieRookie/audit-log/log"
	"github.com/CodingCookieRookie/audit-log/models"
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

func GinHandlerWithError(handler func(*gin.Context) (any, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//userEmail := ctx.Query("email")
		//eventType := ctx.Query("type")
		eventData, err := handler(ctx)
		//timeStamp := time.Now()

		event := &models.Event{
			//UserEmail: userEmail,
			//EventType: eventType,
			//Timestamp: timeStamp,
		}
		if err != nil {
			log.Errorf("error: %v", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewStackTracedError(err))
		} else {
			event.EventData = eventData
			ctx.AbortWithStatusJSON(http.StatusOK, event)
		}
	}
}
