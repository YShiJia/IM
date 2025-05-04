package model

// RedisConfig redis conf
type RedisConfig struct {
	Host     string // host
	Port     int    // port
	Password string // 密码
	Index    int    // redis db index
	PoolSize int    // redis  连接池大小
}

// EtcdConfig etcd conf
type EtcdConfig struct {
	Host     string // host
	Port     int    //port
	Username string // 用户名
	Password string // 密码
}

// KafkaConfig kafka的配置项
type KafkaConfig struct {
	Broker      string // broker ip:host
	Topic       string // topic 主题
	Replication int    // 副本数量
	Partition   int    // 分区（切片）数量
}

// MysqlConfig mysql的配置项
type MysqlConfig struct {
	Host     string // host
	Port     int    // port
	Username string // 用户名
	Password string // 密码
	Database string // 数据库
	Config   string // 额外配置
	MaxConns int    // 连接池最大连接数
}

type MinioConfig struct {
	Host            string // host
	Port            int    // port
	AccessKeyID     string // 访问ID
	SecretAccessKey string // 访问密钥
	UseSSL          bool   // 是否使用SSL
}

type AuthConfig struct {
	AccessSecret string // 访问token加密秘钥
	AccessExpire int64  // 访问token过期时间
}
