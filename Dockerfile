FROM golang:1.18-alpine AS build-env
MAINTAINER danbai

#修改镜像源为国内
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk update
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOPATH="/go"
#安装所需工具
RUN apk add gcc g++ make upx git
#配置时区为中国
RUN apk add tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

#拉取代码
RUN mkdir /build
ADD ./ /build
#构建
WORKDIR /build
RUN go build -ldflags '-w -s' -o web_demo
RUN upx web_demo


FROM alpine:latest
#运行环境
LABEL maintainer="danbai@88.com"
LABEL description="web_demo build image file"
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk update
RUN apk --no-cache add tzdata ca-certificates libc6-compat libgcc libstdc++ apache2-utils vim
#时区
ENV TZ=Asia/Shanghai

#配置时区为中国
RUN apk add tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

RUN mkdir /app
RUN mkdir /app/configs
WORKDIR /app
COPY --from=build-env /build/web_demo /app/web_demo
COPY --from=build-env /build/assets /app/assets
RUN chmod +x /app/web_demo
CMD ["/app/web_demo","-env=pro"]