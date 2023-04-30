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

## 十、Github Actions

### Workflow

- 是一个自动过程。
- 由一个以上的Job组成。
- 由事件触发、计划或手动。
- 将 `.yml` 添加到存储库。

```yml
name: build-and-test

on:
  push:
    branches: [ master ]
  schedule:
    - cron: '*/15 * * * *'

jobs:
  build:
    runs-on: ubuntu-latest
```

### Runner

- 是运行Job的服务
- 一次运行一个Job。
- Github托管或自托管。
- 向Github报告进度、日志和结果。

```yml
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Check out code
        uses: actions/checkout@v2
      - name: Build server
        run: ./build_server.sh
    test:
      needs: build
      runs-on: ubuntu-latest
      steps:
        - run: ./test_server.sh
```

### Job

- 是在同一个Runner上执行的一组步骤。
- 正常作业并行运行。
- 连续运行相关作业。

```yml
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Check out code
        uses: actions/checkout@v2
      - name: Build server
        run: ./build_server.sh
    test:
      needs: build
      runs-on: ubuntu-latest
      steps:
        - run: ./test_server.sh
```

### Step

- 是一个单独的任务。
- 在作业中连续运行。
- 包含多个 Actions

### Action

- 是一个独立的命令。
- 在一个步骤内连续运行。
- 可重复使用。

```yml
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Check out code
        uses: actions/checkout@v2
      - name: Build server
        run: ./build_server.sh
```

## 十一、使用 Gin 实现 HTTP API

### 流行的Web框架

- Gin
- Beego
- Echo
- Revel
- Martini
- Fiber
- Buffalo

### 流行 HTTP Router 库

- FastHttp
- Gorilla Mux
- HttpRouter
- Chi


### 安装 `gin`

```bash
go get -u github.com/gin-gonic/gin
```

## 十二、使用 Viper 从文件或环境变量中加载配置

### 为什么使用文件加载配置？

轻松指定本地开发和测试的默认配置。

### 为什么使用环境变量加载配置？

使用docker容器部署时轻松覆盖默认配置。

### 为什么使用Viper？

`Viper` 是一个非常受环境的包，它可以从配置文件中查找、浇在和解析配置文件。

- 支持多种类型的文件例如 `JSON`、`TOML`、`YAML`、`ENV`、`INI`。
- 支持从远程系统中读取配置信息例如 `ETCD`、`Consul`。
- 支持监视配置文件中的更改，并通知应用程序。
- 使用`Viper`保存我们对文件所作的任何配置修改。

### 安装viper

```bash
go get github.com/spf13/viper
```

## 十三、用于测试Go中的 HTTP API 并实现100%覆盖的MockDB

### 为什么使用 mock 数据库？

- 独立测试 - 隔离测试数据以避免冲突。
- 测试更快 - 减少与数据库对话的大量时间。
- 100%覆盖率 - 轻松设置边缘案例:意外错误

首先，它可以帮助我们更轻松的编写独立的测试。
其次，我们的测试将运行的更快。不必花时间与数据库进行交互并等待返回结果。所有的操作都将再内存中执行，并在同一进程中执行。
第三个也是非常重要的原因是：它允许我们编写实现100%覆盖率的测试。

### 用模拟数据库测试我们的API是否足够好?

是的!我们真正的DB存储已经过测试。

模拟数据库和真实数据库应该实现相同的接口。

### 如何模拟数据？

- 使用假数据库 - 内存(Memory), 实现一个假版本的数据库:在内存中存储数据。 
- 使用数据库存根 - GoMock, 生成和构建返回硬编码值的存根。

### 使用 gomock

安装 `mockgen`
```bash
go get github.com/golang/mock/mockgen@v1.6.0
```

查看 `mockgen` 路径
```bash
which mockgen
```

查看 `mockgen` 帮助
```bash
mockgen -help
```

```bash
mockgen --package mockdb -destination db/mock/store.go github.com/grayjunzi/backend-master-class-golang/db/sqlc Store
```

## 十四、自定义参数验证其