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
    npm install --unsafe-perm && \
    npm install vuex-persistedstate --save
RUN npm run build
