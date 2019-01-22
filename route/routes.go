package route

import (
	"dropout_s_back/controller"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Init ルーティング
func Init(conn *gorm.DB) *gin.Engine {
	r := gin.Default()

	ctrler := controller.NewController(conn)

	accounts := gin.Accounts{
		"admin": "password",
	}
	authorized := r.Group("/", gin.BasicAuth(accounts))
	api := authorized.Group("/api")
	{
		api.GET("/ble/get", ctrler.GetBle)
		api.GET("/message/get", ctrler.GetMessage)
		api.POST("/user/signup", ctrler.SignUp)
		api.POST("/message/post", ctrler.PostMessage)
	}

	return r
}
