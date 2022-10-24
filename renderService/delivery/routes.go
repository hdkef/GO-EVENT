package delivery

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	userRoute := UserRoute{}

	//Render Routes
	//Event
	event := r.Group("/event")
	event.GET("", GetAllEvent)

	//User
	user := r.Group("/user")
	user.GET("/:id", userRoute.RenderGetByID)
	user.POST("/signup", userRoute.RenderSignUp)
	user.POST("/signin", userRoute.RenderSignIn)
	user.PATCH("/:id", userRoute.RenderEdit)

	//API Routes (only if need AJAX)
}
