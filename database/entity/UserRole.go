/**
 * @Time : 2020/12/31 5:30 下午
 * @Author : MassAdobe
 * @Description: entity
**/
package entity

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 5:30 下午
 * @Description: 用户角色实体类
**/
type UserRoleEntity struct {
	UserId          int    `gorm:"column:user_id;type:int(11);not null"`               // 用户ID
	RealName        string `gorm:"column:real_name;type:varchar(32);not null"`         // 真实姓名
	UserName        string `gorm:"column:user_name;type:varchar(64);not null"`         // 用户名
	PassWord        string `gorm:"column:pass_word;type:varchar(128);not null"`        // 密码
	Salt            string `gorm:"column:salt;type:varchar(64);not null"`              // 用户盐值
	RoleId          int    `gorm:"column:role_id;type:int(11);not null"`               // 角色ID
	RoleName        string `gorm:"column:role_name;type:varchar(32);not null"`         // 角色名称
	RoleDescription string `gorm:"column:role_description;type:varchar(255);not null"` // 角色描述
}
