package delivery

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	userRoute := UserRoute{}
	eventRoute := EventRoute{}
	regRoute := RegisterRoute{}
	subsRoute := SubscriptionRoute{}

	//Render Routes
	//home
	// r.GET("", eventRoute.RenderCreate) //if invalid jwt in cookies, -> redirect to /user/signin
	// //dashboard (visible publisher/consumer)
	// r.GET("/dashboard", eventRoute.RenderCreate)

	event := r.Group("/event")
	user := r.Group("/user")

	//USER
	//sign up (visible public)
	user.GET("/signup", userRoute.RenderSignUp)
	user.POST("/signup", userRoute.RenderSignUp)
	//sign in (visible public)
	user.GET("/signin", userRoute.RenderSignIn)
	user.POST("/signin", userRoute.RenderSignIn)
	//view user detail by id (visible public)
	user.GET("/:user-id", userRoute.RenderGetByID)
	//edit user (visible publisher/consumer)
	user.GET("/edit", userRoute.RenderEdit)
	user.PATCH("/edit", userRoute.RenderEdit)

	//EVENT
	//get all (visible public)
	event.GET("", eventRoute.RenderGet)
	//create new event (visible publisher)
	event.GET("/new", eventRoute.RenderCreate)
	event.POST("/new", eventRoute.RenderCreate)
	//edit event by id (visible publisher)
	event.GET("/edit/:event-id", eventRoute.RenderEdit)
	event.PATCH("/edit/:event-id", eventRoute.RenderEdit)
	//view event detail by id (visible public)
	event.GET("/:event-id", eventRoute.RenderGetByID)
	//AJAX (visible publisher)
	event.DELETE("/:event-id", eventRoute.AjaxDelete)

	//REGISTER
	register := event.Group("/register")
	//view register by event-id (visible publisher)
	register.GET("/event-id/:event-id", regRoute.RenderGetByEventID)
	//create register (visible consumer)
	register.GET("/new/:event-id", regRoute.RenderCreate)
	register.POST("/new/:event-id", regRoute.RenderCreate)
	//view register (visible publisher)
	register.GET("/:register-id", regRoute.RenderGetByID)

	//SUBSCRIPTION
	subs := r.Group("/subscription")
	//get subscription by userID (visible consumer)
	subs.GET("", subsRoute.GetByConsumerID)
	//AJAX (visible consumer)
	subs.POST("/:publisher-id", subsRoute.AjaxSubs)
	subs.DELETE("/:subscription-id", subsRoute.AjaxUnSubs)
}
