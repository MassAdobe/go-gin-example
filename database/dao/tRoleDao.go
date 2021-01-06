/**
 * @Time : 2020/12/31 1:23 下午
 * @Author : MassAdobe
 * @Description: dao
**/
package dao

import (
	"github.com/MassAdobe/go-gin-example/database/entity"
	"github.com/MassAdobe/go-gin/goContext"
)

type TRoleDao struct {
	Table *entity.TRole
	C     *goContext.Context
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 1:25 下午
 * @Description: 根据角色ID获取角色信息
**/
func (*TRoleDao) GetRoleInfoById(roleId int) {

}
