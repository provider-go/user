package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/pkg/logger"
	"github.com/provider-go/pkg/output"
	"github.com/provider-go/user/global"
	"github.com/provider-go/user/middleware"
	"github.com/provider-go/user/models"
)

func CreateInfo(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)

	did := output.ParamToString(json["did"])
	username := output.ParamToString(json["username"])
	nickname := output.ParamToString(json["nickname"])
	sex := output.ParamToInt(json["sex"])
	avatar := output.ParamToString(json["avatar"])
	phone := output.ParamToString(json["phone"])
	email := output.ParamToString(json["email"])
	err := models.CreateUserInfo(did, username, nickname, sex, avatar, phone, email)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}

func UpdateInfo(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := output.ParamToInt(json["id"])
	did := output.ParamToString(json["did"])
	username := output.ParamToString(json["username"])
	nickname := output.ParamToString(json["nickname"])
	sex := output.ParamToInt(json["sex"])
	avatar := output.ParamToString(json["avatar"])
	phone := output.ParamToString(json["phone"])
	email := output.ParamToString(json["email"])
	err := models.UpdateUserInfo(id, did, username, nickname, sex, avatar, phone, email)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}

func DeleteInfo(ctx *gin.Context) {
	id := output.ParamToInt(ctx.Query("id"))
	err := models.DeleteUserInfo(id)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}

func ListInfo(ctx *gin.Context) {
	pageSize := output.ParamToInt(ctx.Query("pageSize"))
	pageNum := output.ParamToInt(ctx.Query("pageNum"))
	list, total, err := models.ListUserInfo(pageSize, pageNum)

	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		res := make(map[string]interface{})
		res["records"] = list
		res["total"] = total
		output.ReturnSuccessResponse(ctx, res)
	}
}

func ViewInfo(ctx *gin.Context) {
	id := output.ParamToInt(ctx.Query("id"))
	row, err := models.ViewUserInfo(id)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, row)
	}
}

func LoginByPlugin(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	pluginToken := output.ParamToString(json["pluginToken"])
	claims := middleware.InitJwt(global.SecretKey).ParseToken(pluginToken)
	did, err := claims.GetSubject()
	if err != nil || len(did) < 2 {
		logger.Error("LoginByPlugin", "step", "GetSubject", "err", err)
		output.ReturnErrorResponse(ctx, 9999, "token解析错误~")
		return
	}
	// 对比数据库记录
	item, err := models.ViewUserInfoByDID(did)
	if err != nil {
		if err.Error() == "ErrRecordNotFound" {
			output.ReturnErrorResponse(ctx, 9999, "用户不存在~")
			return
		}
		logger.Error("LoginByPlugin", "step", "ViewManagerUserByPlugin", "err", err)
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
		return
	}
	// 生成token
	token := middleware.InitJwt(global.SecretKey).GenerateToken(item.Username)
	output.ReturnSuccessResponse(ctx, token)
}
