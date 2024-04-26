package service

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(app *gin.Engine) {
	pprof.Register(app)
	router := app.Group("/nekoerp")
	router.POST("/user/login", UserLogin)
	router.GET("/user", UserList)
	router.PUT("/user/block", BlockUser)
	router.PUT("/user/unblock", UnBlockUser)
	router.PUT("/user/edit", EditUser)
	router.POST("/user", AddUser)
	router.GET("/storage", ListStorage)
	router.GET("/di", ListDi)
	router.POST("/di", Di)
	router.POST("/tiao", Tiao)
	router.GET("/tiao", ListTiao)
	router.GET("/goods", ListGoods)
	router.POST("/goods", AddGoods)
}
