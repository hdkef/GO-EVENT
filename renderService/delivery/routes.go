package delivery

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	userRoute := UserRoute{}
	eventRoute := EventRoute{}

	//Render Routes
	//home
	r.GET("", eventRoute.RenderCreate) //if invalid jwt in cookies, -> redirect to /user/signin
	//dashboard
	r.GET("/dashboard", eventRoute.RenderCreate)

	//Event////////////////////////////////////
	event := r.Group("/event")
	//get created event
	cbm := event.Group("/createdbyme")
	cbm.GET("", eventRoute.RenderCreate)
	//get created event detail
	cbm.GET("/:id", eventRoute.RenderCreate)
	//participant
	cbm.GET("/participant/:id", eventRoute.RenderCreate)
	//detail
	event.GET("/:id", eventRoute.RenderGetByID)
	//create
	event.GET("/new", eventRoute.RenderCreate)
	event.POST("/new", eventRoute.RenderCreate)
	//edit
	event.GET("/edit/:id", eventRoute.RenderCreate)
	event.PATCH("/edit/:id", eventRoute.RenderEdit)
	//sendmail
	event.GET("/:id/sendEmail", eventRoute.RenderCreate)
	event.POST("/:id/sendEmail", eventRoute.RenderCreate)

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
	register.POST("/:idEvent", eventRoute.RenderCreate)
	register.GET("/:idEvent", eventRoute.RenderCreate)

	//Presence//////////////////////////////////
	presence := r.Group("/presence")
	presence.POST("/:idEvent", eventRoute.RenderCreate)
	presence.GET("/:idEvent", eventRoute.RenderCreate)

	//API Routes (only if need AJAX)
	//participant
	participant := r.Group("/participant")
	//reject / approve registration
	participantReg := participant.Group("/reg/:id")
	participantReg.GET("/accept", eventRoute.RenderCreate)
	participantReg.GET("/reject", eventRoute.RenderCreate)
	//reject / approve presence
	participantPresence := participant.Group("/presence/:id")
	participantPresence.GET("/accept", eventRoute.RenderCreate)
	participantPresence.GET("/reject", eventRoute.RenderCreate)

	//event
	event.DELETE("/:id", eventRoute.RenderCreate)

	//like
	like := r.Group("/like")
	like.POST("/:idEvent", eventRoute.RenderCreate)
	like.DELETE("/:id", eventRoute.RenderCreate)

	//subscription
	subs := r.Group("/subs")
	subs.POST("/:idPublisher", eventRoute.RenderCreate)
	subs.DELETE("/:id", eventRoute.RenderCreate)
}
