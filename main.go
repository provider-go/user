package user

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/pkg/types"
	"github.com/provider-go/user/global"
	"github.com/provider-go/user/router"
)

type Plugin struct{}

func CreatePlugin() *Plugin {
	return &Plugin{}
}

func CreatePluginAndDB(instance types.PluginNeedInstance) *Plugin {
	global.DB = instance.Mysql
	global.SecretKey = "SecretKey"
	return &Plugin{}
}

func (*Plugin) Register(group *gin.RouterGroup) {
	router.GroupApp.InitRouter(group)
}

func (*Plugin) RouterPath() string {
	return "user"
}
