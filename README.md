# gin-tool-study

### 一、gin相关

参考学习项目: Gin-study-example

https://github.com/EDDYCJY/go-gin-example



中文学习文档

https://gin-gonic.com/zh-cn/docs/



go mod 管理

```
# 初始化go mod
go mod init
# 拉取依赖 依赖包会自动下载到$GOPATH/pkg/mod，多个项目可以共享缓存的mod
go mod download
# 整理依赖关系
go mod tidy
# 缓存到vendor目录 从mod中拷贝到项目的vendor目录下，这样IDE就可以识别了！
go mod vendor
```



### 二、项目相关

###### 1、启动项目：

go run main.go 

###### 2、项目的主要依赖：
Golang V1.17
- gin		github.com/gin-gonic/gin

- gorm     gorm.io/gorm     github.com/jinzhu/gorm

  gorm文档    https://gorm.io/zh_CN/docs/

- mysql

- redis

- ini          gopkg.in/ini.v1

- jwt-go

- crypto

- logrus    https://github.com/sirupsen/logrus

- qiniu-go-sdk

- dbresolver

###### 3、设置环境变量：
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)





### 三、随便记录

删除.DS_Store文件
 find . -name .DS_Store -print0 | xargs -0 git rm -f --ignore-unmatch


 QQ 邮箱
POP3 服务器地址：qq.com（端口：995）
SMTP 服务器地址：smtp.qq.com（端口：465/587）

163 邮箱：
POP3 服务器地址：pop.163.com（端口：110）
SMTP 服务器地址：smtp.163.com（端口：25）

126 邮箱：
POP3 服务器地址：pop.126.com（端口：110）
SMTP 服务器地址：smtp.126.com（端口：25）