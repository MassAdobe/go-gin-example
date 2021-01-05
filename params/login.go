/**
 * @Time : 2020/12/18 11:57 上午
 * @Author : MassAdobe
 * @Description: login
**/
package params

import "time"

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/18 11:58 上午
 * @Description: SignIn入参
**/
type SignInParam struct {
	UserName  string `json:"user_name" validate:"required" comment:"用户名"`
	PassWord  string `json:"pass_word" validate:"required" comment:"密码"`
	Timestamp int64  `json:"timestamp" validate:"required" comment:"时间戳"`
}

type SignInNacos struct {
	NacosTestInt    int `json:"nacos_test_int" yaml:"nacos-test-int"`
	NacosTestStruct struct {
		NacosTestString string `json:"nacos_test_string" yaml:"nacos-test-string"`
		NacosTestBool   bool   `json:"nacos_test_bool" yaml:"nacos-test-bool"`
	} `json:"nacos_test_struct" yaml:"nacos_test_struct"`
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/18 1:07 下午
 * @Description: SignIn出参
**/
type SignInRtn struct {
	UserName        string `json:"user_name"`
	PassWord        string `json:"pass_word"`
	Timestamp       int64  `json:"timestamp"`
	NacosTestInt    int    `json:"nacos_test_int"`
	NacosTestString string `json:"nacos_test_string"`
	NacosTestBool   bool   `json:"nacos_test_bool"`
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/18 2:52 下午
 * @Description: GetUser入参
**/
type GetUserParam struct {
	PageNum  int    `validate:"required" comment:"翻页参数(当前页)"`
	PageSize int    `validate:"required" comment:"单页数据量"`
	UserId   int    `validate:"required" comment:"用户ID"`
	UserName string `comment:"用户名"`
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/18 2:53 下午
 * @Description: GetUser出参
**/
type GetUserRtn struct {
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	UserType string `json:"user_type"`
	UserSex  string `json:"user_sex"`
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/21 9:33 上午
 * @Description: 获取用户额外信息入参
**/
type GetUserExternalParam struct {
	UserId int `json:"user_id" validate:"required" comment:"用户ID"`
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/21 9:33 上午
 * @Description: 获取用户额外信息出参
**/
type GetUserExternalRtn struct {
	UserType string `json:"user_type"`
	UserSex  string `json:"user_sex"`
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/21 9:33 上午
 * @Description: 获取用户额外信息入参
**/
type PostUserExternalParam struct {
	UserId int `json:"user_id" validate:"required" comment:"用户ID"`
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/21 9:33 上午
 * @Description: 获取用户额外信息出参
**/
type PostUserExternalRtn struct {
	UserType string `json:"user_type"`
	UserSex  string `json:"user_sex"`
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/21 9:33 上午
 * @Description: 获取用户额外信息入参
**/
type PutUserExternalParam struct {
	UserId int `json:"user_id" validate:"required" comment:"用户ID"`
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/21 9:33 上午
 * @Description: 获取用户额外信息出参
**/
type PutUserExternalRtn struct {
	UserType string `json:"user_type"`
	UserSex  string `json:"user_sex"`
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/21 9:33 上午
 * @Description: 获取用户额外信息入参
**/
type DeleteUserExternalParam struct {
	UserId int `json:"user_id" validate:"required" comment:"用户ID"`
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/21 9:33 上午
 * @Description: 获取用户额外信息出参
**/
type DeleteUserExternalRtn struct {
	UserType string `json:"user_type"`
	UserSex  string `json:"user_sex"`
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 1:56 下午
 * @Description: 根据ID获取用户信息出参
**/
type GetUserInfoRtn struct {
	ID        int       `json:"id"`         // 用户ID
	RealName  string    `json:"real_name"`  // 真实姓名
	UserName  string    `json:"user_name"`  // 用户名
	PassWord  string    `json:"pass_word"`  // 密码
	Salt      string    `json:"salt"`       // 用户盐值
	Enabled   string    `json:"enabled"`    // 是否有效(0-有效；1-无效)
	Deleted   string    `json:"deleted"`    // 是否删除(0-未删除,1-已删除)
	CreatedTm time.Time `json:"created_tm"` // 创建时间
	CreatedBy int       `json:"created_by"` // 创建人ID(0:sys)
	UpdatedTm time.Time `json:"updated_tm"` // 更新时间
	UpdatedBy int       `json:"updated_by"` // 更新人ID(0:sys)
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 5:45 下午
 * @Description: 根据ID获取用户角色信息出参
**/
type GetUserRoleInfoRtn struct {
	UserId          int    `json:"user_id"`          // 用户ID
	RealName        string `json:"real_name"`        // 真实姓名
	UserName        string `json:"user_name"`        // 用户名
	PassWord        string `json:"pass_word"`        // 密码
	Salt            string `json:"salt"`             // 用户盐值
	RoleId          int    `json:"role_id"`          // 角色ID
	RoleName        string `json:"role_name"`        // 角色名称
	RoleDescription string `json:"role_description"` // 角色描述
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 5:58 下午
 * @Description: 新增用户入参
**/
type AddUserParam struct {
	RealName string `json:"real_name" validate:"required" comment:"真是姓名"`
	UserName string `json:"user_name" validate:"required" comment:"用户名"`
	PassWord string `json:"pass_word" validate:"required" comment:"密码"`
	Salt     string `json:"salt" validate:"required" comment:"盐值"`
}
