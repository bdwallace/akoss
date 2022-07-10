#!/bin/bash

version=$1

if [ -e $version ];then
    echo "no version!"
    exit 1
fi

# 52.76.119.1:7149/devops/gopub  正式环境
# 18.139.82.154:7149/devops/gopub 旧预发布
#192.168.75.11:7149/devops/gopub
docker_repo="
harbor.one2.cc/devops/akgo
"

docker build . -t harbor.one2.cc/devops/akgo:$version

#docker login -u docker -p harbordev.one2.cc
#docker push $repo:$version

for repo in $docker_repo;do
    echo $repo:$version
    docker login -u docker -p H6MUai+gsbMJSTCA harbor.one2.cc
    docker tag harbor.one2.cc/devops/akgo:$version $repo:$version
    docker push $repo:$version
done
