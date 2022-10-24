package delivery

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	userRoute := UserRoute{}
	eventRoute := EventRoute{}

	//Render Routes
	//home
	r.GET("", eventRoute.Create) //if invalid jwt in cookies, -> redirect to /user/signin
	//dashboard
	r.GET("/dashboard", eventRoute.Create)

	//Event////////////////////////////////////
	event := r.Group("/event")
	//get created event
	cbm := event.Group("/createdbyme")
	cbm.GET("", eventRoute.Create)
	//get created event detail
	cbm.GET("/:id", eventRoute.Create)
	//participant
	cbm.GET("/participant/:id", eventRoute.Create)
	//detail
	event.GET("/:id", eventRoute.Create)
	//create
	event.GET("/new", eventRoute.Create)
	event.POST("/new", eventRoute.Create)
	//edit
	event.GET("/edit/:id", eventRoute.Create)
	event.PATCH("/edit/:id", eventRoute.Create)
	//sendmail
	event.GET("/:id/sendEmail", eventRoute.Create)
	event.POST("/:id/sendEmail", eventRoute.Create)

	//User/////////////////////////////////////
	user := r.Group("/user")
	//detail
	user.GET("/:id", userRoute.RenderGetByID)
	//signup
	user.GET("/signup", userRoute.RenderSignUp)
	user.POST("/signup", userRoute.RenderSignUp)
	//signin
	user.GET("/signin", userRoute.RenderSignIn)
	user.POST("/signin", userRoute.RenderSignIn)
	//edit
	user.GET("/edit/:id", userRoute.RenderEdit)
	user.PATCH("/edit/:id", userRoute.RenderEdit)

	//Register///////////////////////////////////
	register := r.Group("/register")
	//create
	register.POST("/:idEvent", eventRoute.Create)
	register.GET("/:idEvent", eventRoute.Create)

	//Presence//////////////////////////////////
	presence := r.Group("/presence")
	presence.POST("/:idEvent", eventRoute.Create)
	presence.GET("/:idEvent", eventRoute.Create)

	//API Routes (only if need AJAX)
	//participant
	participant := r.Group("/participant")
	//reject / approve registration
	participantReg := participant.Group("/reg/:id")
	participantReg.GET("/accept", eventRoute.Create)
	participantReg.GET("/reject", eventRoute.Create)
	//reject / approve presence
	participantPresence := participant.Group("/presence/:id")
	participantPresence.GET("/accept", eventRoute.Create)
	participantPresence.GET("/reject", eventRoute.Create)

	//event
	event.DELETE("/:id", eventRoute.Create)

	//like
	like := r.Group("/like")
	like.POST("/:idEvent", eventRoute.Create)
	like.DELETE("/:id", eventRoute.Create)

	//subscription
	subs := r.Group("/subs")
	subs.POST("/:idPublisher", eventRoute.Create)
	subs.DELETE("/:id", eventRoute.Create)
}
