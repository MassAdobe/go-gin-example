/**
 * @Time : 2020/12/18 10:11 上午
 * @Author : MassAdobe
 * @Description: entity
**/
package entity

import (
	"time"
)

// TRole 角色表
type TRole struct {
	ID              int       `gorm:"primary_key;column:id;type:int(11);not null"`        // 角色ID
	RoleName        string    `gorm:"column:role_name;type:varchar(32);not null"`         // 角色名称
	RoleDescription string    `gorm:"column:role_description;type:varchar(255);not null"` // 角色描述
	Enabled         string    `gorm:"column:enabled;type:char(1);not null"`               // 是否有效(0-有效；1-无效)
	Deleted         string    `gorm:"column:deleted;type:char(1);not null"`               // 是否删除(0-未删除,1-已删除)
	CreatedTm       time.Time `gorm:"column:created_tm;type:timestamp;not null"`          // 创建时间
	CreatedBy       int       `gorm:"column:created_by;type:int(11);not null"`            // 创建人ID(0:sys)
	UpdatedTm       time.Time `gorm:"column:updated_tm;type:timestamp;not null"`          // 更新时间
	UpdatedBy       int       `gorm:"column:updated_by;type:int(11);not null"`            // 更新人ID(0:sys)
}

// TableName get sql table name.获取数据库表名
func (m *TRole) TableName() string {
	return "t_role"
}

// TUser 用户表
type TUser struct {
	ID        int       `gorm:"primary_key;column:id;type:int(11);not null"` // 用户ID
	RealName  string    `gorm:"column:real_name;type:varchar(32);not null"`  // 真实姓名
	UserName  string    `gorm:"column:user_name;type:varchar(64);not null"`  // 用户名
	PassWord  string    `gorm:"column:pass_word;type:varchar(128);not null"` // 密码
	Salt      string    `gorm:"column:salt;type:varchar(64);not null"`       // 用户盐值
	Enabled   string    `gorm:"column:enabled;type:char(1);not null"`        // 是否有效(0-有效；1-无效)
	Deleted   string    `gorm:"column:deleted;type:char(1);not null"`        // 是否删除(0-未删除,1-已删除)
	CreatedTm time.Time `gorm:"column:created_tm;type:timestamp;not null"`   // 创建时间
	CreatedBy int       `gorm:"column:created_by;type:int(11);not null"`     // 创建人ID(0:sys)
	UpdatedTm time.Time `gorm:"column:updated_tm;type:timestamp;not null"`   // 更新时间
	UpdatedBy int       `gorm:"column:updated_by;type:int(11);not null"`     // 更新人ID(0:sys)
}

// TableName get sql table name.获取数据库表名
func (m *TUser) TableName() string {
	return "t_user"
}

// TUserRole 用户角色关联表
type TUserRole struct {
	ID        int       `gorm:"primary_key;column:id;type:int(11);not null"` // 用户角色关联ID
	UserID    int       `gorm:"column:user_id;type:int(11);not null"`        // 用户ID
	RoleID    int       `gorm:"column:role_id;type:int(11);not null"`        // 角色ID
	Deleted   string    `gorm:"column:deleted;type:char(1);not null"`        // 是否删除(0-未删除,1-已删除)
	CreatedTm time.Time `gorm:"column:created_tm;type:timestamp;not null"`   // 创建时间
	CreatedBy int       `gorm:"column:created_by;type:int(11);not null"`     // 创建人ID(0:sys)
	UpdatedTm time.Time `gorm:"column:updated_tm;type:timestamp;not null"`   // 更新时间
	UpdatedBy int       `gorm:"column:updated_by;type:int(11);not null"`     // 更新人ID(0:sys)
}

// TableName get sql table name.获取数据库表名
func (m *TUserRole) TableName() string {
	return "t_user_role"
}
