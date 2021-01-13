# go-gin-example
接入服务示例

---

### 创建服务方式：
#### 配置文件：
```yaml
# 是否开启nacos配置中心
NacosConfiguration: true
# 是否开启nacos服务注册于发现
NacosDiscovery: true
# nacos地址
NacosServerIps: '127.0.0.1'
# nacos端口号
NacosServerPort: 8848
# nacos命名空间
NacosClientNamespaceId: 'f3e0c037-7fe1-452f-8f37-16b3810846b5'
# 请求Nacos服务端的超时时间（ms）
NacosClientTimeoutMs: 5000
# nacos配置文件名称
NacosDataId: 'go-framework.yml'
# nacos配置组名称
NacosGroup: 'go-framework'
# 项目启动环境：DEBUG-开发，TEST-测试，RELEASE-生产
ProgramEnv: 'DEBUG'
# 日志输出路径(本地配置优先级最高)
LogPath: ''
# 日志级别(本地配置优先级最高)
LogLevel: ''
# 版本 不能有"."
Version: 'v1-0-0'
```

---

#### main函数创建
```go
import (
	"fmt"
	"github.com/MassAdobe/go-gin-example/router"
	"github.com/MassAdobe/go-gin/logs"
	"github.com/MassAdobe/go-gin/nacos"
	"github.com/MassAdobe/go-gin/start"
	"net/http"
	"os"
	"strconv"
	"time"
)

/**
 * @Description: 启动项
**/
func main() {
	rtr := router.Routers() // 配置gin启动
	server := &http.Server{ // 创建服务
		Addr:           ":" + strconv.Itoa(int(nacos.InitConfiguration.Serve.Port)),
		Handler:        rtr,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logs.Lg.SysInfo("启动", logs.Desc(fmt.Sprintf("启动端口: %d", nacos.InitConfiguration.Serve.Port)))
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed { // 监听并启动服务
			logs.Lg.SysError("启动失败", err)
			os.Exit(1)
		}
	}()
	start.GracefulShutdown(server) // 优雅停服
}
```

---

#### 必须创建router文件：
```go
import (
	"github.com/MassAdobe/go-gin-example/controller"
	"github.com/MassAdobe/go-gin/filter"
	"github.com/MassAdobe/go-gin/goContext"
	"github.com/MassAdobe/go-gin/nacos"
	"github.com/MassAdobe/go-gin/routers"
	"github.com/gin-gonic/gin"
)

/**
 * @Description: 配置路由组
 * 主目录下创建router文件夹，创建routers文件，主要设置restFul接口路径和methods
**/
func Routers() (rtr *gin.Engine) {
	rtr = routers.NewRouter()
	// 登录
	login := rtr.Group(nacos.RequestPath("login")).Use(filter.SetTraceAndStep())
	{
		login.POST("/signIn", goContext.Handle(controller.SignIn))                       // 登录
		login.GET("/getUser", goContext.Handle(controller.GetUser))                      // 获取用户信息
		login.GET("/getUserExternal", goContext.Handle(controller.GetUserExternal))      // 获取用户额外信息
		login.POST("/postUserExternal", goContext.Handle(controller.PostUserExternal))   // 获取用户额外信息
		login.PUT("/putUserExternal", goContext.Handle(controller.PutUserExternal))      // 获取用户额外信息
		login.DELETE("/deleteUserExternal", goContext.Handle(controller.DeleteExternal)) // 获取用户额外信息
		login.GET("/getUserInfo", goContext.Handle(controller.GetUserInfo))              // 根据ID获取用户信息
		login.GET("/getUserRoleInfo", goContext.Handle(controller.GetUserRoleInfo))      // 根据ID获取用户角色信息
		login.POST("/addUser", goContext.Handle(controller.AddUser))                     // 新增用户
		login.PUT("/updateUser", goContext.Handle(controller.UpdateUser))                // 更新用户
		login.DELETE("/deleteUser", goContext.Handle(controller.DeleteUser))             // 删除用户
	}
	return
}
```

---

#### 创建全局业务错误处理：
```go
import "github.com/MassAdobe/go-gin/errs"

/**
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
 * @Description: 初始化
**/
func init() {
	errs.AddErrs(CodeDescMap)
}
```

---

#### 创建controller：
```go
import (
	"errors"
	"github.com/MassAdobe/go-gin-example/external/goFramework"
	"github.com/MassAdobe/go-gin-example/params"
	"github.com/MassAdobe/go-gin-example/service"
	"github.com/MassAdobe/go-gin/goContext"
	"github.com/MassAdobe/go-gin/logs"
	"github.com/MassAdobe/go-gin/nacos"
	"github.com/MassAdobe/go-gin/systemUtils"
	"github.com/MassAdobe/go-gin/validated"
	"net/url"
	"strconv"
	"time"
)

/**
 * @Description: 其他注入实体类
**/
var (
	testNacos *params.SignInNacos
)

/**
 * @Description: 注入实体类
**/
func init() {
	// 注入其他实体类
	//testNacos = &params.SignInNacos{} // ① 第一种实例化方式
	testNacos = new(params.SignInNacos) // ② 第二种实例化方式
	nacos.InsertSelfProfile(testNacos)  // 放入监听容器中
}

/**
 * @Description: 登录(POST)
**/
func SignIn(c *goContext.Context) {
	signInParam := new(params.SignInParam)
	validated.BindAndCheck(c, signInParam) // 获取传参
	c.Debug("登录")
	c.Info("登录", logs.Desc("abc"))
	c.Error("登录", errors.New("login error"))
	// panic(errs.NewError(wrong.ErrLoginCode)) // 使用系统panic方式
	time.Sleep(time.Second * 10) // 测试请求超时
	// 返回信息
	c.SuccRes(&params.SignInRtn{
		UserName:        signInParam.UserName,
		PassWord:        signInParam.PassWord,
		Timestamp:       signInParam.Timestamp,
		NacosTestInt:    testNacos.NacosTestInt,
		NacosTestBool:   testNacos.NacosTestStruct.NacosTestBool,
		NacosTestString: testNacos.NacosTestStruct.NacosTestString,
	})
}

/**
* @Description: 获取用户信息(GET)
**/
func GetUser(c *goContext.Context) {
	getUserParam := new(params.GetUserParam)
	getUserParam.PageNum, _ = strconv.Atoi(c.GinContext.Query("page_num"))        // 获取参数
	getUserParam.PageSize, _ = strconv.Atoi(c.GinContext.Query("page_size"))      // 获取参数
	getUserParam.UserId, _ = strconv.Atoi(c.GinContext.Query("user_id"))          // 获取参数
	getUserParam.UserName, _ = url.QueryUnescape(c.GinContext.Query("user_name")) // 获取参数
	validated.CheckParams(getUserParam)                                           // 检查入参
	c.SuccRes(&params.GetUserRtn{
		UserId:   getUserParam.UserId,
		UserName: getUserParam.UserName,
		PageNum:  getUserParam.PageNum,
		PageSize: getUserParam.PageSize,
	})
}

/**
* @Description: 获取用户额外信息(GET-FEIGN)
**/
func GetUserExternal(c *goContext.Context) {
	c.Info("获取用户额外信息(GET)")
	getUserExternalParam := new(params.GetUserExternalParam)
	getUserExternalParam.UserId, _ = strconv.Atoi(c.GinContext.Query("user_id")) // 获取参数
	validated.CheckParams(getUserExternalParam)                                  // 检查入参
	external := goFramework.GetUserExternal(getUserExternalParam.UserId, c)      // 调用GET的feign方法
	c.SuccRes(&params.GetUserExternalRtn{
		UserType: external.UserType,
		UserSex:  external.UserSex,
	})
}

/**
* @Description: 获取用户额外信息(POST-FEIGN)
**/
func PostUserExternal(c *goContext.Context) {
	logs.Lg.Info("获取用户额外信息(POST)", c)
	postUserParam := new(params.PostUserExternalParam)
	validated.BindAndCheck(c, postUserParam)
	external := goFramework.PostUserExternal(postUserParam.UserId, c) // 调用POST的feign方法
	c.SuccRes(&params.PostUserExternalRtn{
		UserType: external.UserType,
		UserSex:  external.UserSex,
	})
}

/**
* @Description: 获取用户额外信息(PUT-FEIGN)
**/
func PutUserExternal(c *goContext.Context) {
	logs.Lg.Info("获取用户额外信息(PUT)", c)
	putUserExternalParam := new(params.PutUserExternalParam)
	putUserExternalParam.UserId, _ = strconv.Atoi(c.GinContext.Query("user_id")) // 获取参数
	validated.CheckParams(putUserExternalParam)                                  // 检查入参
	external := goFramework.PutUserExternal(putUserExternalParam.UserId, c)      // 调用PUT的feign方法
	c.SuccRes(&params.PutUserExternalRtn{
		UserType: external.UserType,
		UserSex:  external.UserSex,
	})
}

/**
* @Description: 获取用户额外信息(DELETE-FEIGN)
**/
func DeleteExternal(c *goContext.Context) {
	logs.Lg.Info("获取用户额外信息(DELETE)", c)
	deleteUserExternalParam := new(params.DeleteUserExternalParam)
	deleteUserExternalParam.UserId, _ = strconv.Atoi(c.GinContext.Query("user_id")) // 获取参数
	validated.CheckParams(deleteUserExternalParam)                                  // 检查入参
	external := goFramework.DeleteUserExternal(deleteUserExternalParam.UserId, c)   // 调用DELETE的feign方法
	c.SuccRes(&params.DeleteUserExternalRtn{
		UserType: external.UserType,
		UserSex:  external.UserSex,
	})
}

/**
* @Description: 根据ID获取用户信息
**/
func GetUserInfo(c *goContext.Context) {
	id, _ := strconv.Atoi(c.GinContext.Query("user_id"))
	loginService := &service.Login{C: c} // 实例化Service
	rtn := new(params.GetUserInfoRtn)
	systemUtils.CopyProperty(rtn, loginService.GetUserInfo(id)) // 调用方法
	c.SuccRes(rtn)
}

/**
* @Description: 根据ID获取用户角色信息
**/
func GetUserRoleInfo(c *goContext.Context) {
	id, _ := strconv.Atoi(c.GinContext.Query("user_id"))
	loginService := &service.Login{C: c} // 实例化Service
	rtn := new(params.GetUserRoleInfoRtn)
	systemUtils.CopyProperty(rtn, loginService.GetUserRoleInfo(id)) // 调用方法
	c.SuccRes(rtn)
}

/**
* @Description: 新增用户
**/
func AddUser(c *goContext.Context) {
	addUser := new(params.AddUserParam)
	validated.BindAndCheck(c, addUser)
	loginService := &service.Login{C: c} // 实例化Service
	c.SuccRes(loginService.AddUser(addUser)) // 调用方法
}

/**
* @Description: 更新用户
**/
func UpdateUser(c *goContext.Context) {
	id, _ := strconv.Atoi(c.GinContext.Query("user_id"))
	userName, _ := url.QueryUnescape(c.GinContext.Query("user_name"))
	loginService := &service.Login{C: c} // 实例化Service
	loginService.UpdateUser(id, userName) // 调用方法
	c.SuccRes(nil)
}

/**
* @Description: 删除用户
**/
func DeleteUser(c *goContext.Context) {
	id, _ := strconv.Atoi(c.GinContext.Query("user_id"))
	loginService := &service.Login{C: c} // 实例化Service
	loginService.DeleteUser(id) // 调用方法
	c.SuccRes(nil)
}
```

---

#### 创建service
```go
import (
	"github.com/MassAdobe/go-gin-example/database/dao"
	"github.com/MassAdobe/go-gin-example/database/entity"
	"github.com/MassAdobe/go-gin-example/database/joinDao"
	"github.com/MassAdobe/go-gin-example/params"
	"github.com/MassAdobe/go-gin/goContext"
	"github.com/MassAdobe/go-gin/logs"
)

type Login struct {
	C *goContext.Context
}

/**
 * @Description: 根据ID获取用户信息
**/
func (this *Login) GetUserInfo(userId int) (user *entity.TUser) {
	tUserDao := &dao.TUserDao{C: this.C} // 实例化Dao
	user = tUserDao.GetUserInfo(userId)
	this.C.Debug("根据ID获取用户信息-Service", logs.Desc(user))
	return
}

/**
 * @Description: 根据ID获取用户角色信息
**/
func (this *Login) GetUserRoleInfo(userId int) (userRole *entity.UserRoleEntity) {
	userRoleDao := &joinDao.UserRoleDao{C: this.C} // 实例化Dao
	userRole = userRoleDao.GetUserAndRoleInfoByUserId(userId)
	this.C.Debug("根据ID获取用户角色信息-Service", logs.Desc(userRole))
	return
}

/**
 * @Description: 新增用户
**/
func (this *Login) AddUser(user *params.AddUserParam) int {
	tUserDao := &dao.TUserDao{C: this.C} // 实例化Dao
	id := tUserDao.AddUser(user)
	this.C.Debug("新增用户-Service", logs.Desc(id))
	return id
}

/**
 * @Description: 更新用户
**/
func (this *Login) UpdateUser(id int, username string) {
	tUserDao := &dao.TUserDao{C: this.C} // 实例化Dao
	tUserDao.UpdateUser(id, username)
	this.C.Debug("更新用户-Service")
}

/**
 * @Description: 删除用户
**/
func (this *Login) DeleteUser(id int) {
	tUserDao := &dao.TUserDao{C: this.C} // 实例化Dao
	tUserDao.DeleteUser(id)
	this.C.Debug("删除用户-Service")
}

```

---

#### 创建Dao类
##### 基于gorm框架
##### 程序
```go
/**
 * @Description: 接口实体类
**/
type TUserDao struct {
	Table *entity.TUser
	C     *context.Context
}

/**
 * @Description: 根据ID获取用户信息
**/
func (this *TUserDao) GetUserInfo(userId int) (user *entity.TUser) {
	user = new(entity.TUser)
	// 读库使用
	if find := db.Read.Table(this.Table.TableName()).Where("deleted = ? and id = ?", constants.UNDELETED, userId).Find(&user);
		find.Error != nil && find.Error != gorm.ErrRecordNotFound {
		this.C.Error("根据ID获取用户信息", find.Error)
		panic(errs.NewError(wrong.ErrFindUserInfoCode))
	}
	this.C.Debug("根据ID获取用户信息-Dao", logs.Desc(user))
	return
}

/**
 * @Description: 新增用户
**/
func (this *TUserDao) AddUser(user *params.AddUserParam) int {
	userCreate := new(entity.TUser)
	systemUtils.CopyProperty(userCreate, user)
	// 写库使用
	userCreate.CreatedBy, userCreate.CreatedTm, userCreate.UpdatedBy, userCreate.UpdatedTm = 0, time.Now(), 0, time.Now()
	if create := db.Write.Table(this.Table.TableName()).Create(&userCreate); create.RowsAffected == 0 || create.Error != nil {
		this.C.Error("新增用户", create.Error)
		panic(errs.NewError(wrong.ErrAddUserCode))
	}
	this.C.Debug("新增用户-Dao", logs.Desc(userCreate))
	return userCreate.ID
}

/**
 * @Description: 更新用户
**/
func (this *TUserDao) UpdateUser(id int, username string) {
	user := &entity.TUser{UserName: username, UpdatedBy: 1, UpdatedTm: time.Now()}
    // 写库使用
	if update := db.Write.Table(this.Table.TableName()).Where("id = ?", id).Update(&user); update.Error != nil {
		this.C.Error("更新用户", update.Error)
		panic(errs.NewError(wrong.ErrAddUserCode))
	}
	this.C.Debug("更新用户-Dao")
}

/**
 * @Description: 删除用户
**/
func (this *TUserDao) DeleteUser(id int) {
	user := &entity.TUser{Deleted: constants.DELETED, UpdatedBy: 1, UpdatedTm: time.Now()}
    // 写库使用
	if update := db.Write.Table(this.Table.TableName()).Where("id = ?", id).Update(&user); update.Error != nil {
		this.C.Error("删除用户", update.Error)
		panic(errs.NewError(wrong.ErrAddUserCode))
	}
	this.C.Debug("删除用户-Dao")
}
```
##### 实体类
```go
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
```

#### join语句
##### 程序
```go
type UserRoleDao struct {
	C *context.Context
}

/**
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
		this.C.Error("根据用户ID获取用户和角色信息", err)
		panic(errs.NewError(wrong.ErrFindUserInfoCode))
	}
	this.C.Debug("根据用户ID获取用户和角色信息-Dao", logs.Desc(userRole))
	return
}
```
##### 实体类
```go
/**
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
```

---

#### 接口出入参
```go
/**
 * @Description: SignIn入参 validate为校验参数的使用方式
**/
type SignInParam struct {
	UserName  string `json:"user_name" validate:"required" comment:"用户名"`
	PassWord  string `json:"pass_word" validate:"required" comment:"密码"`
	Timestamp int64  `json:"timestamp" validate:"required" comment:"时间戳"`
}

/**
 * @Description: nacos配置文件参数
 * 对应配置：
 * nacos-test-int: 100
 * nacos_test_struct:
 *   nacos-test-string: 'string word'
 *   nacos-test-bool: false
**/
type SignInNacos struct {
	NacosTestInt    int `json:"nacos_test_int" yaml:"nacos-test-int"`
	NacosTestStruct struct {
		NacosTestString string `json:"nacos_test_string" yaml:"nacos-test-string"`
		NacosTestBool   bool   `json:"nacos_test_bool" yaml:"nacos-test-bool"`
	} `json:"nacos_test_struct" yaml:"nacos_test_struct"`
}

/**
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
```

---

#### feign接口
##### controller
```go
import (
	"github.com/MassAdobe/go-gin/logs"
	"github.com/MassAdobe/go-gin/validated"
	"github.com/gin-gonic/gin"
	"strconv"
)

/**
 * @Description: 获取用户额外信息
**/
func GetUserExternal(c *context.Context) {
	c.Info("获取用户额外信息(GET)")
	userExternalParam := new(UserExternalParam)
	userExternalParam.UserId, _ = strconv.Atoi(c.GinContext.Query("user_id")) // 用户ID
	validated.CheckParams(userExternalParam)                       // 检查入参
	// 返回信息
	c.SuccResFeign(&UserExternalRtn{
		UserType: "A",
		UserSex:  "男",
	})
}

/**
 * @Description: 获取用户额外信息
**/
func PostUserExternal(c *context.Context) {
	c.Info("获取用户额外信息(POST)")
	userExternalParam := new(UserExternalParam)
	validated.BindAndCheck(c, userExternalParam)
	// 返回信息
	c.SuccResFeign(&UserExternalRtn{
		UserType: "B",
		UserSex:  "女",
	})
}

/**
 * @Description: 获取用户额外信息
**/
func PutUserExternal(c *context.Context) {
	c.Info("获取用户额外信息(PUT)")
	userExternalParam := new(UserExternalParam)
	userExternalParam.UserId, _ = strconv.Atoi(c.GinContext.Query("user_id")) // 用户ID
	validated.CheckParams(userExternalParam)                       // 检查入参
	// 返回信息
	c.SuccResFeign(&UserExternalRtn{
		UserType: "C",
		UserSex:  "男",
	})
}

/**
 * @Description: 获取用户额外信息
**/
func DeleteUserExternal(c *context.Context) {
	c.Info("获取用户额外信息(DELETE)")
	userExternalParam := new(UserExternalParam)
	userExternalParam.UserId, _ = strconv.Atoi(c.GinContext.Query("user_id")) // 用户ID
	validated.CheckParams(userExternalParam)                       // 检查入参
	// 返回信息
	c.SuccResFeign(&UserExternalRtn{
		UserType: "D",
		UserSex:  "女",
	})
}
```

##### 参数
```go
/**
 * @Description: UserExternal入参
**/
type UserExternalParam struct {
	UserId int `json:"user_id" validate:"required" comment:"用户ID"`
}

/**
 * @Description: UserExternal出参
**/
type UserExternalRtn struct {
	UserType string `json:"user_type"`
	UserSex  string `json:"user_sex"`
}
```

##### routers
```go
import (
	"github.com/MassAdobe/go-gin-example/controller"
	"github.com/MassAdobe/go-gin/filter"
	"github.com/MassAdobe/go-gin/goContext"
	"github.com/MassAdobe/go-gin/nacos"
	"github.com/MassAdobe/go-gin/routers"
	"github.com/gin-gonic/gin"
)

/**
 * @Description: 配置路由组
**/
func Routers() (rtr *gin.Engine) {
	rtr = routers.NewRouter()
	// 登录
	login := rtr.Group(nacos.RequestPath("login")).Use(filter.SetTraceAndStep())
	{
		login.POST("/signIn", goContext.Handle(controller.SignIn))                       // 登录
		login.GET("/getUser", goContext.Handle(controller.GetUser))                      // 获取用户信息
		login.GET("/getUserExternal", goContext.Handle(controller.GetUserExternal))      // 获取用户额外信息
		login.POST("/postUserExternal", goContext.Handle(controller.PostUserExternal))   // 获取用户额外信息
		login.PUT("/putUserExternal", goContext.Handle(controller.PutUserExternal))      // 获取用户额外信息
		login.DELETE("/deleteUserExternal", goContext.Handle(controller.DeleteExternal)) // 获取用户额外信息
		login.GET("/getUserInfo", goContext.Handle(controller.GetUserInfo))              // 根据ID获取用户信息
		login.GET("/getUserRoleInfo", goContext.Handle(controller.GetUserRoleInfo))      // 根据ID获取用户角色信息
		login.POST("/addUser", goContext.Handle(controller.AddUser))                     // 新增用户
		login.PUT("/updateUser", goContext.Handle(controller.UpdateUser))                // 更新用户
		login.DELETE("/deleteUser", goContext.Handle(controller.DeleteUser))             // 删除用户
	}
	return
}
```