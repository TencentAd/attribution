#!/bin/bash

# 编译
root=$(cd "$(dirname "$0")" && cd .. && pwd)

cd "${root}" || { exit 1; }
mkdir -p .build
cd .build || { exit 1; }
go build -o attribution_server "${root}/attribution_server.go"

# 启动服务



# 发送请求


# 停止服务

