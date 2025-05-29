package routes

import (

	"github.com/nethsaraPrabash/potral-backend/src/controller"


	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(r *gin.Engine, db*gorm.DB) {
	
	apiRoutes := r.Group("/api")
	{
		userRoutes := apiRoutes.Group("/user")
		{
			userRoutes.POST("/register", controller.RegisterUser)

			userRoutes.POST("/login", controller.LoginUser)
		}
	}
}