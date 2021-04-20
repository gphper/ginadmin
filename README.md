# GinAdmin
这个项目是以Gin框架为基础搭建的后台管理模块<br>

## 依赖
* golang > 1.8
## 依赖
* Gin
* BootStrap
* LayUi
* WebUpload

## 功能列表
- [x] **用户登录**
- [x] **权限管理**
- [x] **用户组管理**
- [x] **文件上传**
- [x] **优雅关闭重启**

### 使用文档
- [项目目录](#结构)
- [分页](#分页)
- [日志](#日志)
- [数据库](#数据库)
- [定时任务](#定时任务)
- [配置文件](#配置文件)
- [用户权限](#用户权限)

### <a name="结构">项目目录</a>
```
|--api  // Api接口控制器
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

### <a name="分页">分页</a>

1.  使用 `comment/util.go` 里面的 `PageOperation` 进行分页
    ```go
    adminDb := models.Db.Table("admin_users").Select("nickname","username").Where("uid != ?", 1)
    adminUserData := comment.PageOperation(c, adminDb, 1, &adminUserList)
    ```
2.  在html中使用
    ```go
    {{ .adminUserData.PageHtml }}
    ```

### <a name="日志">日志</a>
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

### <a name="数据库">数据库</a>

1. 数据库迁移，将定义好的model填充写到下面的 `AutoMigrate` 方法中

   ```go
   Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&AdminUsers{},&AdminGroup{})
   ```

2. 数据填充，将数据写入到 `models\default.go` 下面的 `FillData` 中

### <a name="定时任务">定时任务</a>

-    在 `comment/cron/cron.go`  添加定时执行任务

### <a name="配置文件">配置文件</a>

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

   











