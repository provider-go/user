package router

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/user/api"
)

type Group struct {
	Router
}

var GroupApp = new(Group)

type Router struct{}

func (s *Router) InitRouter(Router *gin.RouterGroup) {
	{
		// user_infos 表操作
		Router.POST("info", api.CreateInfo)
		Router.PUT("info", api.UpdateInfo)
		Router.DELETE("info", api.DeleteInfo)
		Router.GET("infos", api.ListInfo)
		Router.GET("info", api.ViewInfo)
		// user_keys 表操作
		Router.POST("key", api.CreateKey)
		Router.PUT("key", api.UpdateKey)
		Router.DELETE("key", api.DeleteKey)
		Router.GET("keys", api.ListKey)
		Router.GET("key", api.ViewKey)
	}
}
