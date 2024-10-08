package controllers

import (
	"bluebell/global"
	"bluebell/logic"
	"bluebell/models"
	"errors"

	"github.com/gin-gonic/gin"
)

// SignUpHandler 实现注册功能
func SignUpHandler(c *gin.Context) {
	// 获取参数并且参数校验
	ps := new(models.ParamSignUp)
	if err := c.ShouldBind(ps); err != nil {
		global.Log.Error("SignUp with invaild params,they are not json params", err.Error())
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 业务处理
	err := logic.SignUp(ps)
	if err != nil {
		global.Log.Error("SignUp with invaild params", err.Error())
		if err == errors.New("用户已存在,请换个用户名") {
			ResponseError(c, CodeUserExist)
			return
		} else {
			ResponseErrorWithData(c, CodeInvalidParam, err.Error())
		}
		return
	}
	// 返回响应
	ResponseSucceed(c, "succeed singup")
}

// LoginHandler 实现登录功能
func LoginHandler(c *gin.Context) {
	// 获取参数并且参数校验
	pl := new(models.ParamLogin)
	err := c.ShouldBind(pl)
	if err != nil {
		global.Log.Error("Login with invaild params,they are not json params", err.Error())
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 业务处理
	token, err := logic.Login(pl)
	if err != nil {
		global.Log.Error("Login with invaild params")
		if err == errors.New("用户名不存在") {
			ResponseError(c, CodeUserExist)
			return
		} else if err == errors.New("密码错误") {
			ResponseError(c, CodeInvalidPassword)
			return
		} else {
			ResponseErrorWithData(c, CodeInvalidParam, err.Error())
			return
		}
	}
	// 返回响应
	type data struct {
		Token    string 
		UserName string
	}
	d := &data{
		Token:    token,
		UserName: pl.Username,
	}
	ResponseSucceed(c, d)
}
