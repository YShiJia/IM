package model

type ServeInfo struct {
	Name         string
	IP           string
	HttpPort     int
	GrpcPort     int
	Type         ServeType
	RecvMsgQueue KafkaConfig
}
