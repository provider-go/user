package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/pkg/output"
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
	did := output.ParamToString(json["did"])
	username := output.ParamToString(json["username"])
	nickname := output.ParamToString(json["nickname"])
	sex := output.ParamToInt(json["sex"])
	avatar := output.ParamToString(json["avatar"])
	phone := output.ParamToString(json["phone"])
	email := output.ParamToString(json["email"])
	err := models.UpdateUserInfo(did, username, nickname, sex, avatar, phone, email)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}

func DeleteInfo(ctx *gin.Context) {
	did := output.ParamToString(ctx.Query("did"))
	err := models.DeleteUserInfo(did)
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
	did := output.ParamToString(ctx.Query("did"))
	row, err := models.ViewUserInfo(did)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, row)
	}
}
