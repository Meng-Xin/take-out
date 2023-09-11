# 苍穹外卖Golang实现
> 如果你是一个Java程序员不知道在Golang中如何进行Web开发，那么你或许可以参考该项目进行实践。
> 该项目已经提供了一个初始的项目架构其思想和Java中的面向接口开发一致，但Golang中并没有SpringBoot那样
> 强大的功能例如便捷的BeanUtils.Copy……。
>

## 1.Quick Start

**拉取项目**

```shell
$ git clone https://github.com/Meng-Xin/sky-take-out-go.git
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

```go
go run main.go
```

