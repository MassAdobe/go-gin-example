/**
 * @Time : 2020/12/28 6:27 下午
 * @Author : MassAdobe
 * @Description: errors
**/
package wrong

import "github.com/MassAdobe/go-gin/errs"

/**
 * @Author: MassAdobe
 * @TIME: 2020-04-26 21:05
 * @Description: 错误封装
**/
const (
	//自定义错误码
	ErrLoginCode = 1000 + iota // 登录错误
	ErrFindUserInfoCode
	ErrFindUserRoleInfoCode
	ErrAddUserCode
	ErrUpdateUserCode
	ErrDeleteUserCode

	//自定义错误描述
	ErrLoginDesc            = "登录错误(用户名密码错误或不存在相关用户)"
	ErrFindUserInfoDesc     = "获取用户信息失败"
	ErrFindUserRoleInfoDesc = "获取用户角色信息失败"
	ErrAddUserDesc          = "新增用户失败"
	ErrUpdateUserDesc       = "更新用户失败"
	ErrDeleteUserDesc       = "删除用户失败"
)

/**
 * @Author: MassAdobe
 * @TIME: 2020-04-26 21:06
 * @Description: 错误参数体
**/
var CodeDescMap = map[int]string{
	// 自定义错误
	ErrLoginCode:            ErrLoginDesc,
	ErrFindUserInfoCode:     ErrFindUserInfoDesc,
	ErrFindUserRoleInfoCode: ErrFindUserRoleInfoDesc,
	ErrAddUserCode:          ErrAddUserDesc,
	ErrUpdateUserCode:       ErrUpdateUserDesc,
	ErrDeleteUserCode:       ErrDeleteUserDesc,
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/28 6:28 下午
 * @Description: 初始化
**/
func init() {
	errs.AddErrs(CodeDescMap)
}
