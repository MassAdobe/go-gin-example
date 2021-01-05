/**
 * @Time : 2020-04-26 19:57
 * @Author : MassAdobe
 * @Description: router
**/
package router

import (
	"github.com/MassAdobe/go-gin-example/controller"
	"github.com/MassAdobe/go-gin/errs"
	"github.com/MassAdobe/go-gin/filter"
	"github.com/MassAdobe/go-gin/nacos"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: MassAdobe
 * @TIME: 2020-04-26 21:08
 * @Description: 配置路由组
**/
func Routers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rtr := gin.New()
	rtr.NoMethod(errs.HandleNotFound) // 处理没有相关方法时的错误处理
	rtr.NoRoute(errs.HandleNotFound)  // 处理没有相关路由时的错误处理
	rtr.Use(errs.ErrHandler())        // 全局错误处理
	// 登录
	login := rtr.Group(nacos.RequestPath("login")).Use(filter.SetTraceAndStep())
	{
		login.POST("/signIn", controller.SignIn)                       // 登录
		login.GET("/getUser", controller.GetUser)                      // 获取用户信息
		login.GET("/getUserExternal", controller.GetUserExternal)      // 获取用户额外信息
		login.POST("/postUserExternal", controller.PostUserExternal)   // 获取用户额外信息
		login.PUT("/putUserExternal", controller.PutUserExternal)      // 获取用户额外信息
		login.DELETE("/deleteUserExternal", controller.DeleteExternal) // 获取用户额外信息
		login.GET("/getUserInfo", controller.GetUserInfo)              // 根据ID获取用户信息
		login.GET("/getUserRoleInfo", controller.GetUserRoleInfo)      // 根据ID获取用户角色信息
		login.POST("/addUser", controller.AddUser)                     // 新增用户
		login.PUT("/updateUser", controller.UpdateUser)                // 更新用户
		login.DELETE("/deleteUser", controller.DeleteUser)             // 删除用户
	}
	return rtr
}
