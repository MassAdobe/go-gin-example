/**
 * @Time : 2020/12/18 6:06 下午
 * @Author : MassAdobe
 * @Description: goFramework
**/
package goFramework

import (
	"encoding/json"
	"github.com/MassAdobe/go-gin/errs"
	"github.com/MassAdobe/go-gin/http"
	"github.com/gin-gonic/gin"
)

const (
	GO_FRAMEWORK_PROVIDER_SERVER_NAME = "go-framework-provider"
	GO_FRAMEWORK_PROVIDER_GROUP_NAME  = "go-framework"
)

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/18 6:33 下午
 * @Description: 获取用户额外信息
**/
func GetUserExternal(userId int, c *gin.Context) *UserExternalRtn {
	feign := &http.FeignRequest{
		Body:       UserExternalParam{UserId: userId},              // 请求参数，可以为空
		ServerName: GO_FRAMEWORK_PROVIDER_SERVER_NAME,              // 服务名，不能为空
		GroupName:  GO_FRAMEWORK_PROVIDER_GROUP_NAME,               // 组别名，不能为空
		Url:        "/go-framework-provider/feign/getUserExternal", // 调用URL(二级路径)
		C:          c,                                              // 当前请求的上下文
	}
	if get, err := feign.FeignGet(); err != nil {
		panic(errs.NewError(errs.ErrGetRequestCode, err))
	} else {
		rtn := new(UserExternalRtn)
		if err := json.Unmarshal(get, &rtn); err != nil {
			panic(errs.NewError(errs.ErrJsonCode, err))
		}
		return rtn
	}
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/18 6:33 下午
 * @Description: 获取用户额外信息
**/
func PostUserExternal(userId int, c *gin.Context) *UserExternalRtn {
	feign := &http.FeignRequest{
		Body:       UserExternalParam{UserId: userId},               // 请求参数，可以为空
		ServerName: GO_FRAMEWORK_PROVIDER_SERVER_NAME,               // 服务名，不能为空
		GroupName:  GO_FRAMEWORK_PROVIDER_GROUP_NAME,                // 组别名，不能为空
		Url:        "/go-framework-provider/feign/postUserExternal", // 调用URL(二级路径)
		C:          c,                                               // 当前请求的上下文
	}
	if get, err := feign.FeignPost(); err != nil {
		panic(errs.NewError(errs.ErrPostRequestCode, err))
	} else {
		rtn := new(UserExternalRtn)
		if err := json.Unmarshal(get, &rtn); err != nil {
			panic(errs.NewError(errs.ErrJsonCode, err))
		}
		return rtn
	}
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/18 6:33 下午
 * @Description: 获取用户额外信息
**/
func PutUserExternal(userId int, c *gin.Context) *UserExternalRtn {
	feign := &http.FeignRequest{
		Body:       UserExternalParam{UserId: userId},              // 请求参数，可以为空
		ServerName: GO_FRAMEWORK_PROVIDER_SERVER_NAME,              // 服务名，不能为空
		GroupName:  GO_FRAMEWORK_PROVIDER_GROUP_NAME,               // 组别名，不能为空
		Url:        "/go-framework-provider/feign/putUserExternal", // 调用URL(二级路径)
		C:          c,                                              // 当前请求的上下文
	}
	if get, err := feign.FeignPut(); err != nil {
		panic(errs.NewError(errs.ErrPutRequestCode, err))
	} else {
		rtn := new(UserExternalRtn)
		if err := json.Unmarshal(get, &rtn); err != nil {
			panic(errs.NewError(errs.ErrJsonCode, err))
		}
		return rtn
	}
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/18 6:33 下午
 * @Description: 获取用户额外信息
**/
func DeleteUserExternal(userId int, c *gin.Context) *UserExternalRtn {
	feign := &http.FeignRequest{
		Body:       UserExternalParam{UserId: userId},                 // 请求参数，可以为空
		ServerName: GO_FRAMEWORK_PROVIDER_SERVER_NAME,                 // 服务名，不能为空
		GroupName:  GO_FRAMEWORK_PROVIDER_GROUP_NAME,                  // 组别名，不能为空
		Url:        "/go-framework-provider/feign/deleteUserExternal", // 调用URL(二级路径)
		C:          c,                                                 // 当前请求的上下文
	}
	if get, err := feign.FeignDelete(); err != nil {
		panic(errs.NewError(errs.ErrDeleteRequestCode, err))
	} else {
		rtn := new(UserExternalRtn)
		if err := json.Unmarshal(get, &rtn); err != nil {
			panic(errs.NewError(errs.ErrJsonCode, err))
		}
		return rtn
	}
}
