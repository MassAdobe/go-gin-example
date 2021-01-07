/**
 * @Time : 2020-04-26 19:57
 * @Author : MassAdobe
 * @Description: router
**/
package router

import (
	"github.com/MassAdobe/go-gin-example/controller"
	"github.com/MassAdobe/go-gin/filter"
	"github.com/MassAdobe/go-gin/goContext"
	"github.com/MassAdobe/go-gin/nacos"
	"github.com/MassAdobe/go-gin/routers"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: MassAdobe
 * @TIME: 2020-04-26 21:08
 * @Description: 配置路由组
**/
func Routers() (rtr *gin.Engine) {
	rtr = routers.NewRouter()
	// 登录
	login := rtr.Group(nacos.RequestPath("login")).Use(filter.SetTraceAndStep())
	{
		login.POST("/signIn", goContext.Handle(controller.SignIn))                       // 登录
		login.GET("/getUser", goContext.Handle(controller.GetUser))                      // 获取用户信息
		login.GET("/getUserExternal", goContext.Handle(controller.GetUserExternal))      // 获取用户额外信息
		login.POST("/postUserExternal", goContext.Handle(controller.PostUserExternal))   // 获取用户额外信息
		login.PUT("/putUserExternal", goContext.Handle(controller.PutUserExternal))      // 获取用户额外信息
		login.DELETE("/deleteUserExternal", goContext.Handle(controller.DeleteExternal)) // 获取用户额外信息
		login.GET("/getUserInfo", goContext.Handle(controller.GetUserInfo))              // 根据ID获取用户信息
		login.GET("/getUserRoleInfo", goContext.Handle(controller.GetUserRoleInfo))      // 根据ID获取用户角色信息
		login.POST("/addUser", goContext.Handle(controller.AddUser))                     // 新增用户
		login.PUT("/updateUser", goContext.Handle(controller.UpdateUser))                // 更新用户
		login.DELETE("/deleteUser", goContext.Handle(controller.DeleteUser))             // 删除用户
		login.POST("/testIdempotent", filter.GetReqUser(), filter.ValidIdempotent(),
			goContext.Handle(controller.TestIdempotent)) // 测试幂等接口
	}
	return
}
