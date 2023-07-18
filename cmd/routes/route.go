package routes

import (
	"mark_emailchaser/pkg/api/handlers"
	"mark_emailchaser/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(app *gin.Engine) {

	jwtMiddleware := middleware.NewJWTMiddleware()

	userRoutes := app.Group("/user")
	userRoutes.POST("/signup", handlers.CreateUser)
	userRoutes.POST("/login", handlers.Login)
	userRoutes.Use(jwtMiddleware.Protect())
	{
		userRoutes.GET("/:id", handlers.GetProfile)
		userRoutes.DELETE("/:id", handlers.DeleteUser)
	}
}

func RegisterReservationRoutes(app *gin.Engine) {
	jwtMiddleware := middleware.NewJWTMiddleware()
	reservationRoutes := app.Group("/reservation")

	reservationRoutes.Use(jwtMiddleware.Protect())
	{
		reservationRoutes.POST("/create", handlers.CreateReservation)
		reservationRoutes.GET("/:id", handlers.GetReservation)
		reservationRoutes.DELETE("/:id", handlers.DeleteReservation)
		reservationRoutes.PUT("/update", handlers.UpdateReservation)
	}
}
