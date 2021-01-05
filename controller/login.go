/**
 * @Time : 2020/12/18 11:56 上午
 * @Author : MassAdobe
 * @Description: login
**/
package controller

import (
	"errors"
	"github.com/MassAdobe/go-gin-example/external/goFramework"
	"github.com/MassAdobe/go-gin-example/params"
	"github.com/MassAdobe/go-gin-example/service"
	"github.com/MassAdobe/go-gin/logs"
	"github.com/MassAdobe/go-gin/nacos"
	"github.com/MassAdobe/go-gin/systemUtils"
	"github.com/MassAdobe/go-gin/validated"
	"github.com/gin-gonic/gin"
	"net/url"
	"strconv"
)

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 1:42 下午
 * @Description: 其他注入实体类
**/
var (
	testNacos *params.SignInNacos
)

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 1:42 下午
 * @Description: 注入实体类
**/
func init() {
	// 注入其他实体类
	//testNacos = &params.SignInNacos{}
	testNacos = new(params.SignInNacos)
	nacos.InsertSelfProfile(testNacos)
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/18 1:07 下午
 * @Description: 登录
**/
func SignIn(c *gin.Context) {
	signInParam := new(params.SignInParam)
	validated.BindAndCheck(c, signInParam)
	logs.Lg.Debug("登录", c)
	logs.Lg.Info("登录", c, logs.Desc("abc"))
	logs.Lg.Error("登录", errors.New("login error"), c)
	//panic(errs.NewError(wrong.ErrLoginCode))
	// 返回信息
	validated.SuccRes(c, &params.SignInRtn{
		UserName:        signInParam.UserName,
		PassWord:        signInParam.PassWord,
		Timestamp:       signInParam.Timestamp,
		NacosTestInt:    testNacos.NacosTestInt,
		NacosTestBool:   testNacos.NacosTestStruct.NacosTestBool,
		NacosTestString: testNacos.NacosTestStruct.NacosTestString,
	})
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/18 2:51 下午
 * @Description: 获取用户信息
**/
func GetUser(c *gin.Context) {
	getUserParam := new(params.GetUserParam)
	getUserParam.PageNum, _ = strconv.Atoi(c.Query("page_num"))        // 获取参数
	getUserParam.PageSize, _ = strconv.Atoi(c.Query("page_size"))      // 获取参数
	getUserParam.UserId, _ = strconv.Atoi(c.Query("user_id"))          // 获取参数
	getUserParam.UserName, _ = url.QueryUnescape(c.Query("user_name")) // 获取参数
	validated.CheckParams(getUserParam)                                // 检查入参
	external := goFramework.GetUserExternal(getUserParam.UserId, c)
	validated.SuccRes(c, &params.GetUserRtn{
		UserId:   getUserParam.UserId,
		UserName: getUserParam.UserName,
		PageNum:  getUserParam.PageNum,
		PageSize: getUserParam.PageSize,
		UserType: external.UserType,
		UserSex:  external.UserSex,
	})
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/21 9:30 上午
 * @Description: 获取用户额外信息
**/
func GetUserExternal(c *gin.Context) {
	logs.Lg.Info("获取用户额外信息(GET)", c)
	getUserExternalParam := new(params.GetUserExternalParam)
	getUserExternalParam.UserId, _ = strconv.Atoi(c.Query("user_id")) // 获取参数
	validated.CheckParams(getUserExternalParam)                       // 检查入参
	external := goFramework.GetUserExternal(getUserExternalParam.UserId, c)
	validated.SuccRes(c, &params.GetUserExternalRtn{
		UserType: external.UserType,
		UserSex:  external.UserSex,
	})
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/21 9:30 上午
 * @Description: 获取用户额外信息
**/
func PostUserExternal(c *gin.Context) {
	logs.Lg.Info("获取用户额外信息(POST)", c)
	postUserParam := new(params.PostUserExternalParam)
	validated.BindAndCheck(c, postUserParam)
	external := goFramework.PostUserExternal(postUserParam.UserId, c)
	validated.SuccRes(c, &params.PostUserExternalRtn{
		UserType: external.UserType,
		UserSex:  external.UserSex,
	})
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/21 9:30 上午
 * @Description: 获取用户额外信息
**/
func PutUserExternal(c *gin.Context) {
	logs.Lg.Info("获取用户额外信息(PUT)", c)
	putUserExternalParam := new(params.PutUserExternalParam)
	putUserExternalParam.UserId, _ = strconv.Atoi(c.Query("user_id")) // 获取参数
	validated.CheckParams(putUserExternalParam)                       // 检查入参
	external := goFramework.PutUserExternal(putUserExternalParam.UserId, c)
	validated.SuccRes(c, &params.PutUserExternalRtn{
		UserType: external.UserType,
		UserSex:  external.UserSex,
	})
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/21 9:30 上午
 * @Description: 获取用户额外信息
**/
func DeleteExternal(c *gin.Context) {
	logs.Lg.Info("获取用户额外信息(DELETE)", c)
	deleteUserExternalParam := new(params.DeleteUserExternalParam)
	deleteUserExternalParam.UserId, _ = strconv.Atoi(c.Query("user_id")) // 获取参数
	validated.CheckParams(deleteUserExternalParam)                       // 检查入参
	external := goFramework.DeleteUserExternal(deleteUserExternalParam.UserId, c)
	validated.SuccRes(c, &params.DeleteUserExternalRtn{
		UserType: external.UserType,
		UserSex:  external.UserSex,
	})
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 1:17 下午
 * @Description: 根据ID获取用户信息
**/
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("user_id"))
	loginService := &service.Login{C: c} // 实例化Service
	rtn := new(params.GetUserInfoRtn)
	systemUtils.CopyProperty(rtn, loginService.GetUserInfo(id))
	validated.SuccRes(c, rtn)
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 1:18 下午
 * @Description: 根据ID获取用户角色信息
**/
func GetUserRoleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("user_id"))
	loginService := &service.Login{C: c} // 实例化Service
	rtn := new(params.GetUserRoleInfoRtn)
	systemUtils.CopyProperty(rtn, loginService.GetUserRoleInfo(id))
	validated.SuccRes(c, rtn)
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 6:01 下午
 * @Description: 新增用户
**/
func AddUser(c *gin.Context) {
	addUser := new(params.AddUserParam)
	validated.BindAndCheck(c, addUser)
	loginService := &service.Login{C: c} // 实例化Service
	validated.SuccRes(c, loginService.AddUser(addUser))
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 6:21 下午
 * @Description: 更新用户
**/
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("user_id"))
	userName, _ := url.QueryUnescape(c.Query("user_name"))
	loginService := &service.Login{C: c} // 实例化Service
	loginService.UpdateUser(id, userName)
	validated.SuccRes(c, nil)
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 6:26 下午
 * @Description: 删除用户
**/
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("user_id"))
	loginService := &service.Login{C: c} // 实例化Service
	loginService.DeleteUser(id)
	validated.SuccRes(c, nil)
}