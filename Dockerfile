FROM golang:1.21.1 as builder

# 设置环境变量
ENV HOME /app
ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /app
COPY . .

RUN go build -v -a -o go-webshot ./main.go


FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk --no-cache add ca-certificates && \
    apk update && \
    apk add --no-cache xvfb-run chromium chromium-chromedriver && \
    ln -sf ${LOCAL_PKG}/bin/* /usr/local/bin/

# 设置工作目录
WORKDIR /bin/

# 将上个容器编译的二进制文件复制到 工作目录
COPY --from=builder /app/go-webshot .
COPY SourceHanSansSC-Normal.otf /usr/share/fonts/

ENTRYPOINT ["/bin/go-webshot"]