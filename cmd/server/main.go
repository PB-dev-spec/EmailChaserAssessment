package main

import (
	"github.com/gin-gonic/gin"

	"mark_emailchaser/cmd/routes"
	"mark_emailchaser/pkg/common"
)

func main() {
	common.LoadEnvVariables()
	common.GetDBConnection()
	app := gin.Default()
	routes.RegisterUserRoutes(app)
	routes.RegisterReservationRoutes(app)
	app.Run()
}
