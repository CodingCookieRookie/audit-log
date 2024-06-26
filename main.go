package main

import (
	_ "github.com/CodingCookieRookie/audit-log/constants"
	"github.com/CodingCookieRookie/audit-log/handlers"
	"github.com/CodingCookieRookie/audit-log/my_sql"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	my_sql.Init()
	handlers.Route(engine)
	engine.Run(":3000")
}
