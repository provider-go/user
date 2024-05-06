package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/pkg/output"
	"github.com/provider-go/user/models"
)

func CreateUser(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)

	username := output.ParamToString(json["username"])
	password := output.ParamToString(json["password"])
	nickname := output.ParamToString(json["nickname"])
	sex := output.ParamToInt(json["sex"])
	avatar := output.ParamToString(json["avatar"])
	phone := output.ParamToString(json["phone"])
	email := output.ParamToString(json["email"])
	err := models.CreateUser(username, password, nickname, sex, avatar, phone, email)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}

func DeleteUser(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := output.ParamToInt32(json["id"])
	err := models.DeleteUser(id)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}

func ListUser(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	pageSize := output.ParamToInt(json["pageSize"])
	pageNum := output.ParamToInt(json["pageNum"])
	list, total, err := models.ListUser(pageSize, pageNum)

	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		res := make(map[string]interface{})
		res["records"] = list
		res["total"] = total
		output.ReturnSuccessResponse(ctx, res)
	}
}

func ViewUser(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := output.ParamToInt32(json["id"])
	row, err := models.ViewUser(id)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, row)
	}
}
