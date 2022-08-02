#!/bin/bash
echo "查看protoc版:" ` protoc --version`
echo  "先删除旧的pb文件"
rm -rf pb/*
echo "准备生成pb文件"
protoc -I./protobuf --gofast_out=plugins=grpc:./pb  ./protobuf/*.proto
echo "程序执行完成"
exit


