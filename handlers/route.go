package handlers

import "github.com/gin-gonic/gin"

func Route(engine *gin.Engine) {

	// add jwt token
	//event.GET("", GinHandlerWithError(HandleEventGet))
	//user := engine.Group("user")

	engine.GET("/token", GinHandlerWithError(HandleApiToken))

	//	event := engine.Group("")
	engine.Use(UserAuth)
	engine.POST("/event", GinHandlerWithError(HandleEventPost))
}
