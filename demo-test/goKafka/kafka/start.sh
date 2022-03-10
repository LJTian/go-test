#!/bin/sh

#这是kafka启动脚本
echo "这是kafka启动脚本"

#后台启动 zookeeper--server
cmdExecPath='/usr/local/bin/kafka_2.13-3.1.0/'

cmd1=${cmdExecPath}bin/zookeeper-server-start.sh
cfg1=${cmdExecPath}config/zookeeper.properties
logFile1=zookeeper.log

echo "手动执行一下命令"
echo "nohup ${cmd1} ${cfg1} >${logFile1} 2>&1 &"

#后台启动 kafka
cmd2=${cmdExecPath}bin/kafka-server-start.sh
cfg2=${cmdExecPath}config/server.properties
logFile2=kafka.log

echo "nohup ${cmd2} ${cfg2} >${logFile2} 2>&1 &"
