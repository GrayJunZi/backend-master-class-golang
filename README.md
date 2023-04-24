# Eternity

Backend Master Class [Golang + Postgres + Kubernetes + gRPC]
---
了解有关后端Web开发的一切:Golang,Postgres,Redis,GRPC,Docker,Kubernetes,AWS,CI/CD

## 一、介绍

创建一个简单银行项目，包含以下功能。

- 创建并管理帐号
- 记录余额变更
- 金额交易事务

数据库设计：
- 使用 [dbdiagram.io](https://dbdiagram.io) 设计 SQL 数据库架构。
- 将架构另存为 PDF 或 PNG 图标。
- 生成SQL脚本。

## 二、安装 Docker 和 Postgres

拉取 `postgres` 镜像
```bash
docker pull postgres:12-alpine
```

创建容器
```bash
docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:12-alpine
```

进入容器
```bash
docker exec -it postgres12 psql -U root
```

退出容器
```bash
\q
```

## 三、Golang迁移数据库

### 安装 go

安装 `go`
```bash
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz
```

修改 $HOME/.profile 文件追加如下命令
```bash
export PATH=$PATH:/usr/local/go
```

### 安装 golang-migrate

安装 `golang-migrate`
```bash
mkdir /usr/local/golang-migrate
tar -C /usr/local/golang-migrate -xzf migrate.linux-arm64.tar.gz
```

修改 $HOME/.profile 文件追加如下命令
```bash
export PATH=$PATH:/usr/local/golang-migrate
```

查看 `golang-migrate` 版本
```bash
migrate version
```

### 迁移数据库

生成迁移文件
```bash
migrate create -ext sql -dir db/migration -seq init_schema
```

`migrate up` - 执行脚本更新数据库。
`migrate down` - 从数据库中撤回更改。

查看正在运行中的容器
```bash
docker ps
```

停止正在运行的容器
```bash
docker stop postgres12
```

查看所有容器
```bash
docker ps -a
```

进入容器
```bash
docker exec -it postgres12 /bin/sh
```

创建数据库
```bash
createdb --username=root --owner=root simple_bank
psql simple_bank
\q
```

迁移数据库
```bash
migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up
```

安装 `make`
```bash
sudo apt-get install make
```

查看 `make` 版本
```bash
make --version
```

执行 `Makefile` 文件
```bash
# 向上迁移
make migrateup

# 向下迁移
make migratedown
```