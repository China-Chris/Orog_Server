

# OROG Project
OROG.ai backend core logic implement


## 1. 运行前配置
你需要在一下几个地方将配置
app/etc/app-api.yaml
rpcx/account/etc/account.yaml
rpcx/chains/sol/etc/sol.yaml
rpcx/ws/etc/ws.yaml

将这些如
节点、数据库、redis、kafuka都换成自己的即可


## 2.运行数据库

数据库sql文件在docs文件下


## 3.运行项目

先在本地运行etcd

按照需要运行 Makefile里的 run系列命令

如 make run_account 运行Account 服务

make run_app 运行 Api服务


前端开源仓库
https://github.com/FCY316/OROG_WEB
