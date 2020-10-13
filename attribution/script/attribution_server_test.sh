#!/bin/bash

# 编译
root=$(cd "$(dirname "$0")" && cd .. && pwd)

cd "${root}" || { exit 1; }
source script/common.sh

mkdir -p .build
cd .build || { exit 1; }
go build -o attribution_server "${root}/attribution_server.go"

# 启动服务
./attribution_server \
    -conv_parser_name=ams \
    -click_parser_name=ams \
    -attribution_result_storage=stdout \
    -server_address=:9081 \
    -v=50 \
    -logtostderr &

PID=$!

wait_part_listened 9081 5 || {
    echo "start server failed"
    kill -9 "$PID"
    exit 1
}

# 点击转发请求
curl -XGET 'http://localhost:9081/click?click_time=1602558787&click_id=abcdef&callback=YWRzX3NlcnZpY2UsMTU4NDUxMDI3OSwyNjg5MzNhMzc5MTM0YzBjMDQ4ZGZjMGQyNGYzMTk0NWYzMzJiOWNi&device_os_type=android&imei=d4b8f3898515056278ccf78a7a2cca2d&promoted_object_id=1001&muid=d4b8f3898515056278ccf78a7a2cca2d' |
    grep -q '"message":"success"' || {
    kill -9 "$PID"
    exit 2
}

# 发送转化请求
curl -XPOST 'http://localhost:9081/conv?conv_id=2001&app_id=1001' -d '
{
  "actions": [
    {
      "outer_action_id": "outer_action_identity",
      "action_time": 1602558788,
      "user_id": {
        "hash_imei": "d4b8f3898515056278ccf78a7a2cca2d",
        "hash_android_id": "",
        "oaid": "",
        "hash_oaid": ""
      },
      "action_type": "ACTIVATE_APP",
      "action_param": {
      }
    }
  ]
}
' | grep -q '"message":"success"' || {
    echo "start to handle conv"
    kill -9 "$PID"
    exit 3
}

# 退出程序
kill -9 "$PID"
exit 0
