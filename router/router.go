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
		// user 表操作
		Router.POST("add", api.CreateUser)
		Router.POST("delete", api.DeleteUser)
		Router.POST("list", api.ListUser)
		Router.POST("view", api.ViewUser)
	}
}
