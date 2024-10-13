<h1 align="center" style="border-bottom: none;">IM即时通讯系统</h1>
<h2 align="center" style="border-bottom: none;">基于微服务的高性能分布式IM即时通讯系统</h2>

### 											程序运行说明

---

## 后端

本项目建议在linux系统下运行，首先需要保证系统上有`golang`环境

```shell
# 克隆仓库
git clone https://github.com/YShiJia/IM.git
# 安装依赖
go mod tidy
```

第三方组件依赖：

- Mysql
- Redis
- Etcd
- Zookeeper
- Kafka
- Docker
- Nginx
- 阿里云容器托管
- ... ...

建议使用docker安装和部署相应的依赖，相关教程如下，也可以使用脚本进行一键部署（前提：已配置docker环境）

- [ ] IM系统 docker部署博客

后端分为多个服务：

> Edge Server 边缘服务

1. 监听`webscoket`端口， 与用户保持连接，便于即时收发消息数据
2. 每个edge节点都绑定三个消息队列：向状态服务节点传输接收到的消息、接收来着状态服务节点的指令消息、接收来自状态服务节点的待发送消息。使其在业务执行过程中与其他服务的调用完全解耦，后面会有详细讲解
3. Edge Server 将自身信息注册到Etcd中，以便让状态服务节点对其进行服务发现，并将需要发送的消息推送到Edge Server中

> Status RPC Server 状态RPC服务

1. IM系统的核心调度节点，使用go-zero进行开发，主要职责有：Edge节点的发现与消息推送，将消息路由到指定Edge节点，负载均衡，作为信息中心记录IM系统信息并进行实时调度
2. 内部使用基于Redis实现的一致性Hash算法，并开启后台线程监听和调度Edge节点的启动和停止
3. 统计所有Edge节点中所有用户的连接状态并进行实时更新
4. ... ...

> Status API Server 状态API服务

1. 作为外界获取IM系统状态信息的接口，使用go-zero进行开发，有以下功能：
2. 用户在与IM系统建立webScoket连接前需要从状态API服务节点获取对应的Edge节点地址，以便后续进行连接
3. 用户获取好友在线离线状态
4. ... ...

> USER RPC Server 用户RPC服务

1. 作为IM用户模块的业务处理节点，使用go-zero进行开发，有以下功能：
2. 用户注册，登录，下线等数据和状态更新，和查找删除冻结等业务逻辑
3. ... ...

> USER API Server 用户API服务

1. 使用go-zero进行开发,作为用户RPC服务的对外开放接口，后续可能会推出新功能
2.  ... ...

> Message RPC Server 消息RPC服务

1. 使用go-zero进行开发，主要进行IM系统消息的落盘，查找，同步等业务逻辑
2. ... ...

> Social RPC Server 社交RPC服务

1. 使用go-zero进行开发，主要进行IM系统社交关系如好友，群聊等数据相关的业务逻辑
2. ... ...

---

## 前端（待开发）