/**
 * @Time : 2020/12/31 1:25 下午
 * @Author : MassAdobe
 * @Description: dao
**/
package dao

import (
	"github.com/MassAdobe/go-gin-example/database/entity"
	"github.com/gin-gonic/gin"
)

type TUserRoleDao struct {
	Table *entity.TUserRole
	C     *gin.Context
}
