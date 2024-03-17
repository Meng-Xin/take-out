# 苍穹外卖Golang实现
> 如果你不知道在Golang中如何进行Web开发，那么你或许可以参考该项目进行实践，
> 该项目已经提供了一个初始的项目架构其思想。
> 在这里你或许可以收获以下知识或经验例如：
> 1. 如何在Go Web开发中选择合适的设计模式并设计架构。
> 2. 如何对Gorm中的`Hook`、`Transaction`、`动态SQL`封装,以及了解`Context`在Gorm的中应用场景。
> 3. 如何设计并使用`RouteGroup`解决复杂多变的中间件加载场景问题。
> 4. 以及一些常规的开发经验……

**需求文档：**
[苍穹外卖 (apifox.com)](https://apifox.com/apidoc/shared-93dd7a4f-adbc-4d2b-b6f5-24976908bc1c)

**学习交流群**
> `企鹅群`：828448599

## QuickStart
1.  切换到 web 前端目录,启动nginx.exe
2.  启动Redis
3.  启动Golang服务端
4.  访问 http:localhost

## 功能模块介绍

### 管理端

餐饮企业内部员工使用。 主要功能有：

| 模块      | 描述                                                         |
| --------- | ------------------------------------------------------------ |
| 登录/退出 | 内部员工必须登录后，才可以访问系统管理后台                   |
| 员工管理  | 管理员可以在系统后台对员工信息进行管理，包含查询、新增、编辑、禁用等功能 |
| 分类管理  | 主要对当前餐厅经营的 菜品分类 或 套餐分类 进行管理维护， 包含查询、新增、修改、删除等 |
| 菜品管理  | 主要维护各个分类下的菜品信息，包含查询、新增、修改、删除、启售、停售等功能 |
| 套餐管理  | 主要维护当前餐厅中的套餐信息，包含查询、新增、修改、删除、启售、停售等功能 |
| 订单管理  | 主要维护用户在移动端下的订单信息，包含查询、取消、派送、完成，以及订单报表下载等功能 |
| 数据统计  | 主要完成对餐厅的各类数据统计，如营业额、用户数量、订单等     |

### 用户端

| 模块        | 描述                                                         |
| ----------- | ------------------------------------------------------------ |
| 登录/退出   | 对接微信小程序开放API接口实现微信授权登录。                  |
| 点餐-菜单   | 可在点餐界面选择 菜品分类\|套餐分类，并根据当前选择的分类加载其中的菜品信息，供用户查询选择 |
| 点餐-购物车 | 用户选中的菜品就会加入用户的购物车，主要包含 查询购物车、加入购物车、删除购物车、清空购物车等功能 |
| 订单支付    | 用户选完菜品/套餐后，可以对购物车菜品进行结算支付，这时就需要进行订单的支付 |
| 个人信息    | 用户个人信息界面提供历史订单和收货地址管理，可查看历史订单信息或使用再来一单功能，收货地址可设置默认地址、新增地址、修改地址、删除地址等功能。 |

**技术栈介绍**

+ `Gin`：Gin 是一个用 Go (Golang) 编写的轻量级 HTTP Web 框架，使用责任链模式对中间件加载进行，并且内部封装Sync.Pool、RouterGroup等多种强大的内部组件，是一个较为流行的框架，但开放度太高，容易导致一个人一个开发风格。` go get github.com/gin-gonic/gin`
+ Gorm：是使用较多的一个Object Relational Mapping：对象关系映射。`go get gorm.io/gorm`
+ go-redis：Golang中操作Redis的库 `go get github.com/go-redis/redis`
+ go-jwt： Golang中使用Jwt认证的库 `go get github.com/golang-jwt/jwt`
+ GoCron: Golang中使用的定时任务库：`go get github.com/go-co-op/gocron` 

## 1.Quick Start

**拉取项目**

```shell
# 克隆项目到本地
$ git clone https://github.com/Meng-Xin/sky-take-out-go.git
# 执行 /script下的sky.sql脚本，创建数据库基础数据。
$ sky.sql
```

**下载依赖**

```shell
# 切换到工作目录
cd /takeout/

# 下载依赖 
go mod tidy
# 如果下载缓慢你需要去配置镜像源
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
```

**运行环境**

+ MySQL
+ Redis

**运行服务器**

```shell
# 默认启动使用dev配置
$ go run main.go 
# 指定release配置文件启动
$ go run main.go --env=release
```
## 2.DockerCompose 启动
```shell
# 下载项目
$ git clone https://github.com/Meng-Xin/sky-take-out-go.git
# 创建配置文件需要的共享卷
$ mkdir /home/running/takeout/config
$ mkdir /home/running/takeout/logs
# 切换到运行目录
$ cd /takeout
# 拷贝配置文件到共享卷中
$ cp ./config/*.yaml /home/running/takeout/config/
# 运行docker-compose
$ docker-compose up -d
```
## 项目架构图

```xquery
client/ #WEB客户端
common/ #存放通用内容的包
 |---e #存放自定义错误、错误code、code对应消息。
  |---code.go
  |---error.go
  |---msg.go
 |---enum #存放自定义枚举、常量、变量
 |---utils#工具包、例如jwt、limit限流、Email邮件、泛型工具函数
 |---result.go#自定义通用数据返回格式
config/ #项目配置文件
 |---application-dev.yaml
 |---application-release.yaml
 |---config.go #配置文件解析类
global/ #全局包，存放例如：GormDB、RedisClient、AllConfig……
 |---global.go 
initialize/ #初始化包内部主要是需要初始化构建的组件
 |---enter.go
 |---gorm.go
 |---redis.go
 |---router.go
internal/ #内部包，这里面主要实现Controller、Service、Repository层的操作。
 |---api/
 |---model/
 |---repository/
 |---router/
 |service/
logger/ #日志包，用来管理日志
 |---log.go 
middle/ #中间件包，主要该项目需要使用的中间件、例如身份、权限、限流、等拦截器功能。
 |---jwt_middle.go
script/ #脚本包，主要做一些初始化脚本工作，例如MySQL数据初始化脚本、DevOps发布脚本等。

go.mod #goalng的项目依赖文件，类似于java的maven
main.go #入口函数，项目启动从main函数开始。
 
```
---
**流程实现图**
![流程实现图](http://xiaoxiangzhu.oss-cn-beijing.aliyuncs.com/doc/%E9%A1%B9%E7%9B%AE%E6%A8%A1%E5%9E%8B%E5%9B%BE.png)
---
**项目架构图**
![项目架构图](http://xiaoxiangzhu.oss-cn-beijing.aliyuncs.com/doc/%E6%9E%B6%E6%9E%84%E5%9B%BE.png)
---
**参考实现**
![工作台](http://xiaoxiangzhu.oss-cn-beijing.aliyuncs.com/doc/%E5%B7%A5%E4%BD%9C%E5%8F%B0.png)
![数据统计](http://xiaoxiangzhu.oss-cn-beijing.aliyuncs.com/doc/%E6%95%B0%E6%8D%AE%E7%BB%9F%E8%AE%A1.png)
![订单管理](http://xiaoxiangzhu.oss-cn-beijing.aliyuncs.com/doc/%E8%AE%A2%E5%8D%95%E7%AE%A1%E7%90%86.png)
![套餐管理](http://xiaoxiangzhu.oss-cn-beijing.aliyuncs.com/doc/%E5%A5%97%E9%A4%90%E7%AE%A1%E7%90%86.png)
![移动端购物车](http://xiaoxiangzhu.oss-cn-beijing.aliyuncs.com/doc/%E7%A7%BB%E5%8A%A8%E7%AB%AF%E8%B4%AD%E7%89%A9%E8%BD%A6.png)
![移动端历史订单](http://xiaoxiangzhu.oss-cn-beijing.aliyuncs.com/doc/%E5%8E%86%E5%8F%B2%E8%AE%A2%E5%8D%95.png)
