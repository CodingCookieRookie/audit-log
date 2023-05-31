package handlers

import "github.com/gin-gonic/gin"

func Route(engine *gin.Engine) {

	{
		api := engine.Group("/api")
		api.GET("/token", GinHandlerWithError(HandleApiToken))
	}

	{
		user := engine.Group("/users")
		user.Use(UserAuth)
		user.GET("/events", GinHandlerWithError(HandleEventGet))
		user.POST("/events", GinHandlerWithError(HandleEventPost))
	}

}
