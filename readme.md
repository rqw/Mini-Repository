该项目基于[Maven-Go](https://github.com/fanxcv/Maven-Go)扩展
### 自编译
本人使用的Go版本为: `1.19.3`
```shell
git clone --depth 1 https://gitee.com/renqiwei/mini-repository.git
cd mini-repository
# 使用make编译二进制文件
make
# 本地编译docker镜像
make docker
# 手动编译命令
go mod tidy
go build -o MiniRepos src/main.go
# 交叉编译
# CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o MavenGo src/main.go
chmod a+x MiniRepos
./MiniRepos -c config.yaml
```
### 启动参数
启动时, 可以使用-c指定配置文件路径, 默认加载同目录下的config.yaml
### 配置文件说明
```yaml
listen: 0.0.0.0 # 监听地址
port: 8880 # 监听端口
logging:
  path: /data/log # 文件日志保存地址, 默认为空, 即不写入文件
  level: debug # 日志级别
context: maven # 基础路径
localRepository: /data/data # 本地仓库地址
user: # 认证用户配置, 支持多个
  - name: user
    password: password
repository: # 仓库设置
  - id: public # 仓库ID
    name: public repository # 名字, 随意
    mode: 4 # 模式, 0 无效 2 仅可写 4 仅可读 6 可读写
    cache: true # 是否缓存镜像文件, 默认不缓存
    target: private # 数据目录, localRepository的相对路径, 默认取id值
    mirror: # 镜像地址, 会先尝试在本地加载, 如果加载失败, 会尝试从镜像依次读取
      - https://maven.aliyun.com/nexus/content/repositories/public
      - https://repo1.maven.org/maven2
  - id: private
    name: private repository
    mode: 2
```