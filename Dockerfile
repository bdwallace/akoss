#docker 17.05+版本支持
#sudo docker build -t  akgo .
#sudo docker run --name akgo -p 8192:8192  --restart always  -d   akgo:latest 
FROM golang:1.12.4-alpine3.9 as golang
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update && \
    apk add  bash gcc musl-dev && \
    rm -rf /var/cache/apk/*   /tmp/*     
ADD src/ /data/akgo/src/
ADD control /data/akgo/control
WORKDIR /data/akgo/
RUN ./control build

FROM node:11.14.0-alpine as node 
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update --no-cache && \
    apk add  --no-cache --virtual .gyp && \ 
    rm -rf /var/cache/apk/*   /tmp/* 
ADD ./ /data/akgo/
WORKDIR /data/akgo/vue-akgo
RUN npm config set strict-ssl false && \
    npm config set unsafe_perm true && \
    npm install node-gyp && \
    npm install vue-wechat-title  && \
    npm install file-saver xlsx  && \
    npm install node-sass sass-loader --unsafe-perm  && \
    npm install --unsafe-perm 
RUN npm run build 

FROM alpine:3.9.3
ENV TZ='Asia/Shanghai' 
RUN TERM=linux && export TERM
USER root 
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates bash vim tzdata sudo curl wget openssh docker mariadb-client && \ 
    echo "Asia/Shanghai" > /etc/timezone && \
    cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    rm -rf /var/cache/apk/*   /tmp/*  && \ 
    mkdir -p /data/htdocs && \
    mkdir -p /data/logs && \
    ssh-keygen -q -N "" -f /root/.ssh/id_rsa && \
    #输出的key需要加入发布目标机的 ~/.ssh/authorized_keys
    cat ~/.ssh/id_rsa.pub  
WORKDIR /data/akgo
ADD control /data/akgo/control
COPY --from=golang /data/akgo/src/akgo /data/akgo/src/akgo
COPY --from=golang /data/akgo/src/conf /data/akgo/src/conf
COPY --from=golang /data/akgo/src/logs /data/akgo/src/logs
COPY --from=node /data/akgo/src/views /data/akgo/src/views
COPY --from=node /data/akgo/src/static /data/akgo/src/static
CMD ["./control","rundocker"]