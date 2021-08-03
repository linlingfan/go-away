# Kratos 框架

- 介绍：Kratos 是 B 站基于 Golang 实现的一个开源的面向微服务的框架. 使用 Kratos 可以很方便地构建一个规范的服务.

- 开源 GitHub 地址: https://github.com/go-kratos/kratos

- 文档地址: https://go-kratos.dev/docs/

### 如何创建一个kratos项目

- 安装
```bazaar
go get github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

- 文档

```bazaar
# create project template
kratos new helloworld

cd helloworld
# download modules
go mod download

# generate Proto template
kratos proto add api/helloworld/helloworld.proto
# generate Proto source code
kratos proto client api/helloworld/helloworld.proto
# generate server template
kratos proto server api/helloworld/helloworld.proto -t internal/service

# generate all proto source code, wire, etc.
go generate ./...

```

- 编译执行

```bazaar

mkdir bin
# compile
go build -o ./bin/ ./...
# run
./bin/helloworld -conf ./configs

```

