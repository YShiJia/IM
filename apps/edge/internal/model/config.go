package model

import (
	imModel "github.com/YShiJia/IM/model"
)

// kafka conf
type EdgeKafkaConfig struct {
	SendMessageQueue imModel.KafkaConfig
	RecvMessageQueue imModel.KafkaConfig
}
