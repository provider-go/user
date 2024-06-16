package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/pkg/output"
	"github.com/provider-go/user/models"
)

func CreateKey(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)

	did := output.ParamToString(json["did"])
	pubkey := output.ParamToString(json["pubkey"])
	err := models.CreateUserKey(did, pubkey)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}

func UpdateKey(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	did := output.ParamToString(json["did"])
	pubkey := output.ParamToString(json["pubkey"])
	err := models.UpdateUserKey(did, pubkey)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}

func DeleteKey(ctx *gin.Context) {
	did := output.ParamToString(ctx.Query("did"))
	err := models.DeleteUserKey(did)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, "success")
	}
}

func ListKey(ctx *gin.Context) {
	pageSize := output.ParamToInt(ctx.Query("pageSize"))
	pageNum := output.ParamToInt(ctx.Query("pageNum"))
	list, total, err := models.ListUserKey(pageSize, pageNum)

	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		res := make(map[string]interface{})
		res["records"] = list
		res["total"] = total
		output.ReturnSuccessResponse(ctx, res)
	}
}

func ViewKey(ctx *gin.Context) {
	did := output.ParamToString(ctx.Query("did"))
	row, err := models.ViewUserKey(did)
	if err != nil {
		output.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		output.ReturnSuccessResponse(ctx, row)
	}
}
