<<<<<<< HEAD
# akoss

This is Main Branch
=======
## Docker 快速启动
``` shell
# 创建akgo-mysql
sudo docker run --name akgo-mysql -h akgo-mysql  -p 3306:3306  -v /data/akgo-mysql:/var/lib/mysql -v /etc/localtime:/etc/localtime -e MYSQL_ROOT_PASSWORD=123456  --restart always -d mysql:5.7.24 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci

# 创建akgo使用的mysql数据库ak_go
mysql -uroot -p123456 -h127.0.0.1
mysql > CREATE DATABASE ak_go;

# 使用docker启动akgo
docker run --name akgo -h akgo-pre -p 80:8192  --restart always  -d  --link akgo-mysql:akgo-mysql harbor.one2.cc/devops/akgo:$version

```
## Docker 后续更新
``` shell
# 使用update.sh更新脚本，脚本内容如下，就在服务器的家目录下。
[opsadmin@pre-akgo ~]$ cat update.sh
#!/bin/bash

version=$1
echo $version

docker pull harbor.one2.cc/devops/akgo:$version && \
docker rm -fv akgo; \
docker run --name akgo -h akgo-pre -p 80:8192  --restart always  -d  --link akgo-mysql:akgo-mysql harbor.one2.cc/devops/akgo:$version && \
echo "---- $version ok!"

# 执行更新
[opsadmin@pre-akgo ~]$ ./update.sh v0.7.9
```

### Docker 镜像制作
``` shell
# 使用命令手工docker bulid
sudo docker build -t  harbor.one2.cc/devops/akgo:$version .
sudo docker login -u docker -p r0JT67Ud6Ybt harbor.one2.cc
sudo docker push harbor.one2.cc/devops/akgo:$version

# 使用脚本命令docker bulid，并上传，脚本在项目git的根目录下
./build.sh $version
```


## 源码编译安装
### 编译环境
- golang >= 1.8+
- nodejs >= 4.0.0（编译过程中需要可以连公网下载依赖包）

### 源码编译

``` shell
# 克隆项目
git clone https://gitee.com/dev-ops/gopub.git

# 编译前端,npm较慢可使用cnpm
cd vue-gopub
npm install
npm run build
#调试启动，开发使用模式
npm run dev

# 编译后端
# 修改配置 数据库配置文件在 src/conf/app.conf
# 编译，control需要给可执行权限,并修改go安装目录 export GOROOT=/usr/local/go
./control build
#调试启动，开发使用模式
./control beerun

```

## 其它
#### 将admin密码重置为123456
```
mysql -uroot -p123456 -h127.0.0.1 gopub
update user set password_hash = "$2y$13$8q0MfKpnghuqCL.3FAAjiOkA8kBFNCW.ECUlqWp1zTpMHs9e5xn6u" where username = "admin";
```
>>>>>>> master
