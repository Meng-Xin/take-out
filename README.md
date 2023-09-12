# 苍穹外卖Golang实现
> 如果你是一个Java程序员不知道在Golang中如何进行Web开发，那么你或许可以参考该项目进行实践。
> 该项目已经提供了一个初始的项目架构其思想和Java中的面向接口开发一致，但Golang中并没有SpringBoot那样
> 强大的功能例如便捷的BeanUtils.Copy……。
>

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
[README.md](README.md)
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

## 项目架构图

```xquery
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
![内部架构图](https://cdn.learnku.com/uploads/images/202308/04/97585/feB1ylGNUp.png!large)
---
**项目架构图**
![img.png](img.png)