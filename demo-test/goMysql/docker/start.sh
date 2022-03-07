#!/bin/bash

echo "这是一个docker方式启动mysql的脚本"

echo "参数1:  容器名称(默认mysql)"
echo "参数2:  端口号(默认13306)"
echo "参数3:  ROOT密码(默认123456)"

DockerName=$1
port=$2
RootWd=$3

if [ "1${DockerName}" = "1" ] ; then
	DockerName=mysql
fi

if [ "1${port}" = "1" ]; then
  port=13306
fi

if [ "1${RootWd}" = "1" ]; then
  RootWd=123456
fi

docker run --name=${DockerName} -it -p 13306:3306 -e MYSQL_ROOT_PASSWORD=${RootWd} -d mariadb
if [ $? -ne 0 ] ; then 
	echo "fail!!!"
	exit 1
fi       

DockerId=`docker ps | grep mysql | awk '{print $1}'`

## 修改内部数据库内部属性

## 用bash方式进入容器
echo "==============================================================="
echo "请手动执行一下操作[命令说明在脚本注释中]:"
echo "docker exec -it ${DockerId} bash"

## 登录mysql
echo "mysql -u root -p "
echo "设置的默认密码为:[123456]"
## 授权
echo "grant all privileges on *.*  to 'root'@'%';"
echo "flush privileges;"

# mysql8以后需要 修改密码验证方式否则navicat连接mysql可能会报错
echo "ALTER USER 'root'@'%' IDENTIFIED BY '123456' PASSWORD EXPIRE NEVER;"
echo "exit"
echo "exit"
echo "结束"
echo "==============================================================="
