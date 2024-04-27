package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/pkg/ioput"
	"github.com/provider-go/user/models"
)

func CreateUser(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)

	username := ioput.ParamToString(json["username"])
	password := ioput.ParamToString(json["password"])
	nickname := ioput.ParamToString(json["nickname"])
	sex := ioput.ParamToInt(json["sex"])
	avatar := ioput.ParamToString(json["avatar"])
	phone := ioput.ParamToString(json["phone"])
	email := ioput.ParamToString(json["email"])
	err := models.CreateUser(username, password, nickname, sex, avatar, phone, email)
	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		ioput.ReturnSuccessResponse(ctx, "success")
	}
}

func DeleteUser(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := ioput.ParamToInt32(json["id"])
	err := models.DeleteUser(id)
	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		ioput.ReturnSuccessResponse(ctx, "success")
	}
}

func ListUser(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	pageSize := ioput.ParamToInt(json["pageSize"])
	pageNum := ioput.ParamToInt(json["pageNum"])
	list, total, err := models.ListUser(pageSize, pageNum)

	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		res := make(map[string]interface{})
		res["records"] = list
		res["total"] = total
		ioput.ReturnSuccessResponse(ctx, res)
	}
}

func ViewUser(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := ioput.ParamToInt32(json["id"])
	row, err := models.ViewUser(id)
	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		ioput.ReturnSuccessResponse(ctx, row)
	}
}
