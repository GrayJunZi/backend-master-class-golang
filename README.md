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
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
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

## 四、根据SOL生成CRUD代码

- Create - 在数据库中插入新记录
- Read - 在数据库中选择或搜索记录
- Update - 更改数据库中记录的某些字段
- Delete - 从数据库中删除记录

### database/sql

- 优点是 速度快并且代码简单(Straightforward)
- 缺点是 必须手动将SQL字段映射到变量上，容易出错。

### gorm

- 已经封装好CRUD的代码，只需声明模型。
- 必须学习如何使用gorm提供的函数编写查询。
- 当流量很高(high load)时运行很慢。

### sqlx

- 查询几乎和标准库一样快以及容易使用
- 字段映射是通过查询文本或结构标签完成的

### sqlc

- 速度快且容易使用
- 自动生成代码
- 缺点是只完全支持 PostgreSQL, MySQL处于实验阶段

安装 `sqlc`
```bash
sudo snap install sqlc
```

初始化 `sqlc`
```bash
sqlc init
```

生成代码
```bash
sqlc generate
```

初始化项目
```bash
go mod init github.com/grayjunzi/backend-master-class-golang
go mod tidy
```

## 五、数据库增删改查单元测试

安装依赖包 
```bash
go get github.com/lib/pq
go get github.com/stretchr/testify
```

## 六、实现数据库事务

### 什么是数据库事务？

- 单一的工作单元。
- 通常由多个数据库操作组成。

### 为什么需要事务呢？

1. 提供可靠和一致的工作单元，即使在系统出现故障的情况下。
2. 在并发访问数据库的程序之间提供隔离

### ACID 属性(Property)

数据库事务必须满足ACID属性

- Atomicity(原子性) - 要么所有操作都成功完成，要么事务失败，数据库保持不变。
- Consitency(一致性) - 数据库状态在事务执行之后必须有效。必须满足所有约束。
- Isolation(隔离性) - 并发事务不能相互影响。
- Durability(持久性) - 成功事务写入的数据必须记录在持久性存储中。

## 七、数据库事务锁

### 处理死锁

## 八、避免数据库死锁

## 九、事务隔离级别

### 读现象(Read Phenomena)

- `脏读(Dirty Read)` - 事务读取其他并发未提交事务写入的数据。
- `不可重复读(Non Repeatable Read)` - 一个事务两次读取同一行并看到不同的值，因为它已被其他已提交的事务修改
- `幻读(Phantom Read)` - 事务重新执行用于查找满足条件的行的查询，并且由于其他已提交事务的更改而看到一组不同的行。
- `序列化异常(Serialization Anomaly)` - 一组并发提交的事务的结果是不可能实现的，如果我们试图以任何顺序依次运行它们而不重叠

### 标准隔离级别

American National Standards Institute - ANSI

1. 读取未提交 - 可以看到未提交事务写入的数据。
2. 读取已提交 - 只看到已提交事务写入的数据。
3. 可重复读取 - 相同的读取查询总是返回相同的结果。
4. 可序列化的 - 如果按某种顺序而不是并发地执行事务则可以实现相同的结果。

### MySQL

查看事务隔离级别
```sql
select @@transaction_isolation;
```

查看全局事务隔离界别
```sql
select @@global.transaction_isolation;
```

修改事务隔离级别
```sql
set session transaction isolation level read uncommitted;
```

开启事务
```sql
start transaction;
```

提交事务
```sql
commit;
```

回滚事务
```sq;
rollback;
```

### Postgres

查看事务隔离级别
```bash
show transaction isolation level;
```

修改事务隔离级别
```bash
set transaction isolation level read uncommitted;
```


### 重试机制

处理可能存在错误、超时或死锁。

