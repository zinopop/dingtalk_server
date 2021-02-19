## 生产环境的dockerfile

# FROM loads/alpine:3.8

# LABEL maintainer="yuzheng"

# ###############################################################################
# #                                INSTALLATION
# ###############################################################################

# # 设置固定的项目路径
# ENV WORKDIR /var/www/dingtalk_server

# # 添加应用可执行文件，并设置执行权限
# ADD ./bin/linux_amd64/main   $WORKDIR/main
# RUN chmod +x $WORKDIR/main

# # 添加I18N多语言文件、静态文件、配置文件、模板文件
# ADD i18n     $WORKDIR/i18n
# ADD public   $WORKDIR/public
# ADD config   $WORKDIR/config
# ADD template $WORKDIR/template

# ###############################################################################
# #                                   START
# ###############################################################################
# WORKDIR $WORKDIR
# CMD ./main
# ###############################################################################


# 开发环境的dockerfile
FROM golang:1.15 as builder
# 配置代理
ENV GOPROXY https://goproxy.cn
# 设置go缓存
ENV GO111MODULE=on

# 设置编码格式
ENV LANG en_US.UTF-8
ENV LANGUAGE en_US:en
ENV LC_ALL en_US.UTF-8
ENV TZ=Asia/Shanghai

#编译gf-cli工具
WORKDIR /go/gf-cli
ADD gf-cli .
RUN go mod download
RUN go build -o gf main.go

# go缓存
WORKDIR /go/cache
ADD go.mod .
#ADD go.sum .
RUN go mod download
# 项目工作路径
WORKDIR /go/release

ADD . .

ENTRYPOINT /go/gf-cli/gf run main.go



