package handlers

import "github.com/gin-gonic/gin"

func Route(engine *gin.Engine) {
	event := engine.Group("/event")
	// add jwt token
	//event.GET("", GinHandlerWithError(HandleEventGet))
	event.POST("", GinHandlerWithError(HandleEventPost))
}
