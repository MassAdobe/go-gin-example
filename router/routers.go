/**
 * @Time : 2020-04-26 19:57
 * @Author : MassAdobe
 * @Description: router
**/
package router

import (
	"github.com/MassAdobe/go-gin-example/controller"
	"github.com/MassAdobe/go-gin/context"
	"github.com/MassAdobe/go-gin/filter"
	"github.com/MassAdobe/go-gin/nacos"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: MassAdobe
 * @TIME: 2020-04-26 21:08
 * @Description: 配置路由组
**/
func Routers() (rtr *gin.Engine) {
	rtr = context.NewRouter()
	// 登录
	login := rtr.Group(nacos.RequestPath("login")).Use(filter.SetTraceAndStep())
	{
		login.POST("/signIn", context.Handle(controller.SignIn))                       // 登录
		login.GET("/getUser", context.Handle(controller.GetUser))                      // 获取用户信息
		login.GET("/getUserExternal", context.Handle(controller.GetUserExternal))      // 获取用户额外信息
		login.POST("/postUserExternal", context.Handle(controller.PostUserExternal))   // 获取用户额外信息
		login.PUT("/putUserExternal", context.Handle(controller.PutUserExternal))      // 获取用户额外信息
		login.DELETE("/deleteUserExternal", context.Handle(controller.DeleteExternal)) // 获取用户额外信息
		login.GET("/getUserInfo", context.Handle(controller.GetUserInfo))              // 根据ID获取用户信息
		login.GET("/getUserRoleInfo", context.Handle(controller.GetUserRoleInfo))      // 根据ID获取用户角色信息
		login.POST("/addUser", context.Handle(controller.AddUser))                     // 新增用户
		login.PUT("/updateUser", context.Handle(controller.UpdateUser))                // 更新用户
		login.DELETE("/deleteUser", context.Handle(controller.DeleteUser))             // 删除用户
	}
	return
}
