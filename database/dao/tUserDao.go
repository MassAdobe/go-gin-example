/**
 * @Time : 2020/12/31 1:22 下午
 * @Author : MassAdobe
 * @Description: dao
**/
package dao

import (
	"github.com/MassAdobe/go-gin-example/database/entity"
	"github.com/MassAdobe/go-gin-example/params"
	"github.com/MassAdobe/go-gin-example/wrong"
	"github.com/MassAdobe/go-gin/constants"
	"github.com/MassAdobe/go-gin/context"
	"github.com/MassAdobe/go-gin/db"
	"github.com/MassAdobe/go-gin/errs"
	"github.com/MassAdobe/go-gin/logs"
	"github.com/MassAdobe/go-gin/systemUtils"
	"github.com/jinzhu/gorm"
	"time"
)

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 1:23 下午
 * @Description: 接口实体类
**/
type TUserDao struct {
	Table *entity.TUser
	C     *context.Context
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 1:23 下午
 * @Description: 根据ID获取用户信息
**/
func (this *TUserDao) GetUserInfo(userId int) (user *entity.TUser) {
	user = new(entity.TUser)
	if find := db.Read.Table(this.Table.TableName()).Where("deleted = ? and id = ?", constants.NOT_DELETED, userId).Find(&user); find.Error != nil && find.Error != gorm.ErrRecordNotFound {
		logs.Lg.Error("根据ID获取用户信息", find.Error, this.C)
		panic(errs.NewError(wrong.ErrFindUserInfoCode))
	}
	logs.Lg.Debug("根据ID获取用户信息-Dao", logs.Desc(user), this.C)
	return
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 6:13 下午
 * @Description: 新增用户
**/
func (this *TUserDao) AddUser(user *params.AddUserParam) int {
	userCreate := new(entity.TUser)
	systemUtils.CopyProperty(userCreate, user)
	userCreate.CreatedBy, userCreate.CreatedTm, userCreate.UpdatedBy, userCreate.UpdatedTm = 0, time.Now(), 0, time.Now()
	if create := db.Write.Table(this.Table.TableName()).Create(&userCreate); create.RowsAffected == 0 || create.Error != nil {
		logs.Lg.Error("新增用户", create.Error, this.C)
		panic(errs.NewError(wrong.ErrAddUserCode))
	}
	logs.Lg.Debug("新增用户-Dao", logs.Desc(userCreate), this.C)
	return userCreate.ID
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 6:22 下午
 * @Description: 更新用户
**/
func (this *TUserDao) UpdateUser(id int, username string) {
	user := &entity.TUser{UserName: username, UpdatedBy: 1, UpdatedTm: time.Now()}
	if update := db.Write.Table(this.Table.TableName()).Where("id = ?", id).Update(&user); update.Error != nil {
		logs.Lg.Error("更新用户", update.Error, this.C)
		panic(errs.NewError(wrong.ErrAddUserCode))
	}
	logs.Lg.Debug("更新用户-Dao", this.C)
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 6:22 下午
 * @Description: 删除用户
**/
func (this *TUserDao) DeleteUser(id int) {
	user := &entity.TUser{Deleted: constants.NOT_DELETED, UpdatedBy: 1, UpdatedTm: time.Now()}
	if update := db.Write.Table(this.Table.TableName()).Where("id = ?", id).Update(&user); update.Error != nil {
		logs.Lg.Error("删除用户", update.Error, this.C)
		panic(errs.NewError(wrong.ErrAddUserCode))
	}
	logs.Lg.Debug("删除用户-Dao", this.C)
}
