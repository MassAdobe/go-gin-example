/**
 * @Time : 2020/12/18 6:17 下午
 * @Author : MassAdobe
 * @Description: goFramework
**/
package goFramework

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/18 6:18 下午
 * @Description: UserExternal入参
**/
type UserExternalParam struct {
	UserId int `json:"user_id" validate:"required" comment:"用户ID"`
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/18 6:18 下午
 * @Description: UserExternal出参
**/
type UserExternalRtn struct {
	UserType string `json:"user_type"`
	UserSex  string `json:"user_sex"`
}
