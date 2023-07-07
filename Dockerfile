FROM  golang:1.20-alpine3.18
COPY ./ /go/src/RedisClinetAdd_delete_update_look/
# 进行打包
RUN go env -w GO111MODULE=on && \
go env -w GOPROXY=https://goproxy.cn,direct && \
cd /go/src/RedisClinetAdd_delete_update_look/ && \
# go mod init go && \
go mod tidy && \
go build -o main

FROM alpine:3.18.0
COPY --from=0 /go/src/RedisClinetAdd_delete_update_look/* ./root/
RUN mkdir -p /root/logs/ && \
    chmod +x /root/* && \
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add  --no-cache tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo Asia/Shanghai > /etc/timezone && apk del tzdata
CMD ["/root/main"]