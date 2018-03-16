package configer

import "github.com/go-redis/redis"

type Configer interface {
	SetOption()
	GetConfig() *Config
	GetClient() *redis.Client
}
