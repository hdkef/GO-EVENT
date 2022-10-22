package delivery

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var VERSION string

func init() {
	_ = godotenv.Load()
	VERSION = os.Getenv("APP_VER")
}

func Routes(r *gin.Engine) {

	//ApiVer
	ver := r.Group(fmt.Sprintf("/api/%s", VERSION))

	//Event
	event := ver.Group("/event")
	event.GET("", GetAllEvent)

	//User
	user := ver.Group("/user")
	user.POST("", SignUp)
}
