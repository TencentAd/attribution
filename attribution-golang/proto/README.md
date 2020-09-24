## Proto

proto编译出来的.pb.go文件会一并放到代码库，这里说明对应工具的安装和使用

## 编译

### step 1

安装protoc工具，推荐使用版本[3.7.1](https://github.com/protocolbuffers/protobuf/releases/tag/v3.7.1)

### step 2
安装protoc golang插件

```bash
GIT_TAG="v1.3.2"
go get -d -u github.com/golang/protobuf/protoc-gen-go
cd "$(go env GOPATH)"/src/github.com/golang/protobuf && git fetch --tag && git checkout $GIT_TAG
go install github.com/golang/protobuf/protoc-gen-go
```

### step 3
编译proto文件，如在项目根目录执行

```bash
protoc --go_out=plugins=grpc:. proto/click/click.proto
```

具体使用方法参考 [https://github.com/golang/protobuf](https://github.com/golang/protobuf