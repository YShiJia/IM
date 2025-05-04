package model

// ServiceInfo 服务信息，存储在etcd中，用于服务发现
type ServiceInfo struct {
	Name             string      // 服务名称
	IP               string      // 服务在集群中的IP地址
	HttpPort         int         // 服务开放的HTTP端口
	GrpcPort         int         // 服务开放的GRPC端口
	Type             ServeType   // 服务实例类别，如edge，message
	RecvMessageQueue KafkaConfig // 本服务接收msg的kafka配置
}
