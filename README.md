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

### <a>项目目录</a>
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
