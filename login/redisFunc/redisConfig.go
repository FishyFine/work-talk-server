package redisFunc

type RedisConf struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func NewRedisConf() *RedisConf {
	redisConf := new(RedisConf)
	redisConf.Init()
	return redisConf
}

func (r *RedisConf) Init() {
	r.Host = "47.99.166.254"
	r.Port = "6379"
	r.Password = ""
	r.DB = 0
}
