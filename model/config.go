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

type MysqlConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Config   string
	MaxConns int
}

type MinioConfig struct {
	Host            string
	Port            int
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

type AuthConfig struct {
	AccessSecret string
	AccessExpire int64
}
