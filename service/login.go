/**
 * @Time : 2020/12/31 1:20 下午
 * @Author : MassAdobe
 * @Description: service
**/
package service

import (
	"github.com/MassAdobe/go-gin-example/database/dao"
	"github.com/MassAdobe/go-gin-example/database/entity"
	"github.com/MassAdobe/go-gin-example/database/joinDao"
	"github.com/MassAdobe/go-gin-example/params"
	"github.com/MassAdobe/go-gin/logs"
	"github.com/gin-gonic/gin"
)

type Login struct {
	C *gin.Context
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 1:36 下午
 * @Description: 根据ID获取用户信息
**/
func (this *Login) GetUserInfo(userId int) (user *entity.TUser) {
	tUserDao := &dao.TUserDao{C: this.C} // 实例化Dao
	user = tUserDao.GetUserInfo(userId)
	logs.Lg.Debug("根据ID获取用户信息-Service", this.C, logs.Desc(user))
	return
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 5:44 下午
 * @Description: 根据ID获取用户角色信息
**/
func (this *Login) GetUserRoleInfo(userId int) (userRole *entity.UserRoleEntity) {
	userRoleDao := &joinDao.UserRoleDao{C: this.C} // 实例化Dao
	userRole = userRoleDao.GetUserAndRoleInfoByUserId(userId)
	logs.Lg.Debug("根据ID获取用户角色信息-Service", this.C, logs.Desc(userRole))
	return
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 6:04 下午
 * @Description: 新增用户
**/
func (this *Login) AddUser(user *params.AddUserParam) int {
	tUserDao := &dao.TUserDao{C: this.C} // 实例化Dao
	id := tUserDao.AddUser(user)
	logs.Lg.Debug("新增用户-Service", this.C, logs.Desc(id))
	return id
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 6:22 下午
 * @Description: 更新用户
**/
func (this *Login) UpdateUser(id int, username string) {
	tUserDao := &dao.TUserDao{C: this.C} // 实例化Dao
	tUserDao.UpdateUser(id, username)
	logs.Lg.Debug("更新用户-Service", this.C)
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 6:22 下午
 * @Description: 删除用户
**/
func (this *Login) DeleteUser(id int) {
	tUserDao := &dao.TUserDao{C: this.C} // 实例化Dao
	tUserDao.DeleteUser(id)
	logs.Lg.Debug("删除用户-Service", this.C)
}
