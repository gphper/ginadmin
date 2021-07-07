# GinAdmin
这个项目是以Gin框架为基础搭建的后台管理平台，虽然很多人都认为go是用来开发高性能服务端项目的，但是也难免有要做web管理端的需求，总不能再使用别的语言来开发吧。所以整合出了GinAdmin项目，请大家多提意见指正！

## 依赖
* golang > 1.8
## 依赖
* Gin
* BootStrap
* LayUi
* WebUpload

## 使用文档
- [开始使用](#开始使用)
- [docker-compose构建环境](#docker-compose)
- [项目目录](#结构)
- [分页](#分页)
- [日志](#日志)
- [数据库](#数据库)
- [定时任务](#定时任务)
- [配置文件](#配置文件)
- [模板页面](#模板页面)
- [用户权限](#用户权限)

### :small_blue_diamond:<a name="开始使用">开始使用</a>

1. git 克隆地址 

   ```
   git clone https://github.com/gphper/ginadmin.git
   ```

2. 下载依赖包

   ```go
   go mod download
   ```

3. 配置 `conf/config.ini`文件

   ```
   [mysql]
   username=root
   password=123456
   database=db_beego
   host=127.0.0.1
   port=3306
   max_open_conn=50
   max_idle_conn=20
   [session]
   session_name=gosession_id
   [base]
   port=:8091
   ```

4. 运行 `go run main.go`访问地址 http://localhost:端口地址/admin/login。默认账户：admin  密码：111111

### :small_blue_diamond:<a name="docker-compose">docker-compose构建环境</a>

1. 替换conf目录下的配置项

   ```ini
   [mysql]
   username=docker
   password=123456
   database=docker_mysql
   host=localmysql
   port=3306
   max_open_conn=50
   max_idle_conn=20
   [session]
   session_name=gosession_id
   [base]
   host=0.0.0.0
   port=20010
   fill_data=true
   ```

2. 执行命令 docker-compose up

### :small_blue_diamond:<a name="结构">项目目录</a>

```
|--api  // Api接口控制器
|--cli  // 命令行角色
|--comment // 封装的公共方法
|--conf // 配置文件
|--controllers // Admin控制器存在目录
|--logs // 日志存放目录
|--middleware //中间件
|--models //Gorm中的model类
|--router //自定义路由目录
|--statics //css js等静态文件目录
|--uploadfile //上传文件目录
|--views //视图模板目录
```

### :small_blue_diamond:<a name="分页">分页</a>

1.  使用 `comment/util.go` 里面的 `PageOperation` 进行分页
    ```go
    adminDb := models.Db.Table("admin_users").Select("nickname","username").Where("uid != ?", 1)
    adminUserData := comment.PageOperation(c, adminDb, 1, &adminUserList)
    ```
2.  在html中使用
    ```go
    {{ .adminUserData.PageHtml }}
    ```

### :small_blue_diamond:<a name="日志">日志</a>
1.  自定义日志 在 `comment/loggers` 目录下新建logger
    ```
    参考 userlog.go 文件
    ```
2.  调用自定义的的logger写日志
    ```go
    loggers.UserLogger.Info("无法获取网址",
    zap.String("url", "http://www.baidu.com"),
    zap.Int("attempt", 3),
    zap.Duration("backoff", time.Second),)
    ```

### :small_blue_diamond:<a name="数据库">数据库</a>

1. models下定义的文件均需要实现 `TableName() string`  方法，并将实现该结构体的指针写入到 `GetModels` 方法中

   ```go
   func GetModels() []interface{} {
   	return []interface{}{
   		&AdminUsers{},
   		&AdminGroup{},
   	}
   }
   ```

2. 数据库迁移,在 `cli\cmd`  执行命令行工具

   ```go
   go run ginadmin-cli.go db migrate
   ```

3. 数据填充，需在相应目录下实现 `FillData()` 方法执行如下命令

   ```go
   go run ginadmin-cli.go db seed
   ```

### :small_blue_diamond:<a name="定时任务">定时任务</a>

-    在 `comment/cron/cron.go`  添加定时执行任务

### :small_blue_diamond:<a name="配置文件">配置文件</a>

1. 现在 `conf/conf.go` 添加配置项的 struct 类型，例如

   ```go
   type AppConf struct {
   	BaseConf `ini:"base"`
   }
   type BaseConf struct {
   	Port string `ini:"port"`
   }
   ```

2. 在 `conf/conf.ini` 添加配置信息

   ```
   [base]
   port=:8091
   ```

3. 在代码中调用配置文件的信息

   ```go
   conf.App.BaseConf.Port
   ```

### :small_blue_diamond:<a name="模板页面">模板页面</a>

- 所有的后台模板都写到 `views/template` 目录下面，并且分目录存储，调用时按照 `目录/模板名称` 的方式调用


### :small_blue_diamond:<a name="用户权限">用户权限</a>

- 菜单权限定义到 `comment/menu/menu.go` 文件下，定义完之后在用户组管理里面编辑权限

- 在控制器中可用从 `gin.context` 获取权限

  ```go
  privs,_ := c.Get("userPrivs")
  ```

- template 中判断权限的函数 `judgeContainPriv` 定义在 `comment/template/default.go` 文件下

  ```go
  "judgeContainPriv": func(privMap map[string]interface{},priv string)bool {
  	//判断权限是all的全通过
  	_,o :=privMap["all"]
  	if o {
  		return true
  	}
  	_,ok := privMap[priv]
  	return ok
  },
  ```

  





