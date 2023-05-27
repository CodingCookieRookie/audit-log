package handlers

import "github.com/gin-gonic/gin"

func Route(engine *gin.Engine) {
	event := engine.Group("")
	// add jwt token
	//event.GET("", GinHandlerWithError(HandleEventGet))
	event.POST("/event", GinHandlerWithError(HandleEventPost))
	event.POST("/user", GinHandlerWithError(HandleUserPost))
}
