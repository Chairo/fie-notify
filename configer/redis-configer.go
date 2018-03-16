package configer

import (
	"encoding/json"

	rice "github.com/GeertJohan/go.rice"
	"github.com/go-redis/redis"
)

// 配置文件
type Config struct {
	Host     string `json:"host"`    //redis地址
	Password string `json:"pwd"`     //redis密码
	DB       int    `json:"db"`      //redis数据库
	Channel  string `json:"channel"` //通道
	Source   string `json:"source"`  //监控文件源
	File     string `json:"file"`    //
}

type RedisConfiger struct {
	Client *redis.Client
	Config *Config
}

// 设置连接
func (this *RedisConfiger) SetOption() {
	config := loadConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       config.DB,
	})
	this.Client = client
	this.Config = config
}

func (this *RedisConfiger) GetClient() *redis.Client {
	return this.Client
}

func (this *RedisConfiger) GetConfig() *Config {
	return this.Config
}

func loadConfig() *Config {
	config := &Config{}
	box := rice.MustFindBox(".")
	json.Unmarshal(box.MustBytes("./config.json"), config)
	return config
}

// 创建Redis链接对象
func NewRedisConfiger() *RedisConfiger {
	redis := &RedisConfiger{}
	redis.SetOption()
	return redis
}
