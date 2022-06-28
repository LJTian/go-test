#!/bin/bash
num=0
while [ true ]; do
    num=`expr $num + 1`
    sleep 1
    echo $num >> tian2
done