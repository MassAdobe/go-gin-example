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
	"github.com/MassAdobe/go-gin/goContext"
	"github.com/MassAdobe/go-gin/logs"
	"github.com/MassAdobe/go-gin/nacos"
	"github.com/MassAdobe/go-gin/systemUtils"
	"github.com/MassAdobe/go-gin/validated"
	"net/url"
	"strconv"
	"time"
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
func SignIn(c *goContext.Context) {
	signInParam := new(params.SignInParam)
	validated.BindAndCheck(c, signInParam)
	c.Debug("登录")
	c.Info("登录", logs.Desc("abc"))
	c.Error("登录", errors.New("login error"))
	// panic(errs.NewError(wrong.ErrLoginCode))
	time.Sleep(time.Second * 40)
	// 返回信息
	c.SuccRes(&params.SignInRtn{
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
func GetUser(c *goContext.Context) {
	getUserParam := new(params.GetUserParam)
	getUserParam.PageNum, _ = strconv.Atoi(c.GinContext.Query("page_num"))        // 获取参数
	getUserParam.PageSize, _ = strconv.Atoi(c.GinContext.Query("page_size"))      // 获取参数
	getUserParam.UserId, _ = strconv.Atoi(c.GinContext.Query("user_id"))          // 获取参数
	getUserParam.UserName, _ = url.QueryUnescape(c.GinContext.Query("user_name")) // 获取参数
	validated.CheckParams(getUserParam)                                           // 检查入参
	c.SuccRes(&params.GetUserRtn{
		UserId:   getUserParam.UserId,
		UserName: getUserParam.UserName,
		PageNum:  getUserParam.PageNum,
		PageSize: getUserParam.PageSize,
	})
}

/**
* @Author: MassAdobe
* @TIME: 2020/12/21 9:30 上午
* @Description: 获取用户额外信息
**/
func GetUserExternal(c *goContext.Context) {
	c.Info("获取用户额外信息(GET)")
	getUserExternalParam := new(params.GetUserExternalParam)
	getUserExternalParam.UserId, _ = strconv.Atoi(c.GinContext.Query("user_id")) // 获取参数
	validated.CheckParams(getUserExternalParam)                                  // 检查入参
	external := goFramework.GetUserExternal(getUserExternalParam.UserId, c)
	c.SuccRes(&params.GetUserExternalRtn{
		UserType: external.UserType,
		UserSex:  external.UserSex,
	})
}

/**
* @Author: MassAdobe
* @TIME: 2020/12/21 9:30 上午
* @Description: 获取用户额外信息
**/
func PostUserExternal(c *goContext.Context) {
	logs.Lg.Info("获取用户额外信息(POST)", c)
	postUserParam := new(params.PostUserExternalParam)
	validated.BindAndCheck(c, postUserParam)
	external := goFramework.PostUserExternal(postUserParam.UserId, c)
	c.SuccRes(&params.PostUserExternalRtn{
		UserType: external.UserType,
		UserSex:  external.UserSex,
	})
}

/**
* @Author: MassAdobe
* @TIME: 2020/12/21 9:30 上午
* @Description: 获取用户额外信息
**/
func PutUserExternal(c *goContext.Context) {
	logs.Lg.Info("获取用户额外信息(PUT)", c)
	putUserExternalParam := new(params.PutUserExternalParam)
	putUserExternalParam.UserId, _ = strconv.Atoi(c.GinContext.Query("user_id")) // 获取参数
	validated.CheckParams(putUserExternalParam)                                  // 检查入参
	external := goFramework.PutUserExternal(putUserExternalParam.UserId, c)
	c.SuccRes(&params.PutUserExternalRtn{
		UserType: external.UserType,
		UserSex:  external.UserSex,
	})
}

/**
* @Author: MassAdobe
* @TIME: 2020/12/21 9:30 上午
* @Description: 获取用户额外信息
**/
func DeleteExternal(c *goContext.Context) {
	logs.Lg.Info("获取用户额外信息(DELETE)", c)
	deleteUserExternalParam := new(params.DeleteUserExternalParam)
	deleteUserExternalParam.UserId, _ = strconv.Atoi(c.GinContext.Query("user_id")) // 获取参数
	validated.CheckParams(deleteUserExternalParam)                                  // 检查入参
	external := goFramework.DeleteUserExternal(deleteUserExternalParam.UserId, c)
	c.SuccRes(&params.DeleteUserExternalRtn{
		UserType: external.UserType,
		UserSex:  external.UserSex,
	})
}

/**
* @Author: MassAdobe
* @TIME: 2020/12/31 1:17 下午
* @Description: 根据ID获取用户信息
**/
func GetUserInfo(c *goContext.Context) {
	id, _ := strconv.Atoi(c.GinContext.Query("user_id"))
	loginService := &service.Login{C: c} // 实例化Service
	rtn := new(params.GetUserInfoRtn)
	systemUtils.CopyProperty(rtn, loginService.GetUserInfo(id))
	c.SuccRes(rtn)
}

/**
* @Author: MassAdobe
* @TIME: 2020/12/31 1:18 下午
* @Description: 根据ID获取用户角色信息
**/
func GetUserRoleInfo(c *goContext.Context) {
	id, _ := strconv.Atoi(c.GinContext.Query("user_id"))
	loginService := &service.Login{C: c} // 实例化Service
	rtn := new(params.GetUserRoleInfoRtn)
	systemUtils.CopyProperty(rtn, loginService.GetUserRoleInfo(id))
	c.SuccRes(rtn)
}

/**
* @Author: MassAdobe
* @TIME: 2020/12/31 6:01 下午
* @Description: 新增用户
**/
func AddUser(c *goContext.Context) {
	addUser := new(params.AddUserParam)
	validated.BindAndCheck(c, addUser)
	loginService := &service.Login{C: c} // 实例化Service
	c.SuccRes(loginService.AddUser(addUser))
}

/**
* @Author: MassAdobe
* @TIME: 2020/12/31 6:21 下午
* @Description: 更新用户
**/
func UpdateUser(c *goContext.Context) {
	id, _ := strconv.Atoi(c.GinContext.Query("user_id"))
	userName, _ := url.QueryUnescape(c.GinContext.Query("user_name"))
	loginService := &service.Login{C: c} // 实例化Service
	loginService.UpdateUser(id, userName)
	c.SuccRes(nil)
}

/**
* @Author: MassAdobe
* @TIME: 2020/12/31 6:26 下午
* @Description: 删除用户
**/
func DeleteUser(c *goContext.Context) {
	id, _ := strconv.Atoi(c.GinContext.Query("user_id"))
	loginService := &service.Login{C: c} // 实例化Service
	loginService.DeleteUser(id)
	c.SuccRes(nil)
}

/**
 * @Author: MassAdobe
 * @TIME: 2021/1/7 3:09 下午
 * @Description: 测试幂等接口
**/
func TestIdempotent(c *goContext.Context) {
	testIdempotentParam := new(params.TestIdempotentParam)
	validated.BindAndCheck(c, testIdempotentParam)
	c.SuccRes(&params.TestIdempotentParamRtn{
		String: testIdempotentParam.TestString,
		Int:    testIdempotentParam.TestInt,
	})
}
