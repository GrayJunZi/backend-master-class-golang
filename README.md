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
docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:12-alpine
```

进入容器
```bash
docker exec -it postgres12 psql -U postgres
```

退出容器
```bash
\q
```