package model

// redis conf
type RedisConfig struct {
	Host     string
	Port     int
	Password string
	Index    int
	PoolSize int
}

// etcd conf
type EtcdConfig struct {
	Host string
	Port int
}

type KafkaConfig struct {
	Broker      string
	Topic       string
	TopicPrefix string
	Replication int
	Partition   int
}

type AuthConfig struct {
	AccessSecret string
	AccessExpire int64
}
