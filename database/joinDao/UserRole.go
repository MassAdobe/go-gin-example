/**
 * @Time : 2020/12/31 5:33 下午
 * @Author : MassAdobe
 * @Description: joinDao
**/
package joinDao

import (
	"github.com/MassAdobe/go-gin-example/database/entity"
	"github.com/MassAdobe/go-gin-example/wrong"
	"github.com/MassAdobe/go-gin/db"
	"github.com/MassAdobe/go-gin/errs"
	"github.com/MassAdobe/go-gin/logs"
	"github.com/gin-gonic/gin"
)

type UserRoleDao struct {
	C *gin.Context
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/31 1:26 下午
 * @Description: 根据用户ID获取用户和角色信息
**/
func (this *UserRoleDao) GetUserAndRoleInfoByUserId(userId int) (userRole *entity.UserRoleEntity) {
	sql := `
select a.id               as user_id,
       a.real_name        as real_name,
       a.user_name        as user_name,
       a.pass_word        as pass_word,
       a.salt             as salt,
       c.id               as role_id,
       c.role_name        as role_name,
       c.role_description as role_description
from t_user a
         left join t_user_role b on a.id = b.user_id and b.deleted = '0'
         left join t_role c on b.role_id = c.id and c.deleted = '0'
where a.deleted = '0' and a.id = ?;
`
	userRole = new(entity.UserRoleEntity)
	if err := db.Read.Raw(sql, userId).Scan(&userRole).Error; err != nil {
		logs.Lg.Error("根据用户ID获取用户和角色信息", err, this.C)
		panic(errs.NewError(wrong.ErrFindUserInfoCode))
	}
	logs.Lg.Debug("根据用户ID获取用户和角色信息-Dao", logs.Desc(userRole), this.C)
	return
}
