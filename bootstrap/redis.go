package bootstrap

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"github.com/youkoulayley/api-collection/models"
)

var rds *redis.Client

// OpenRedis open the redis database
func OpenRedis(c *models.Conf) {
	DBRedis := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host + ":" + c.Redis.Port,
		Password: c.Redis.Password,
		DB:       0,
	})

	_, err := DBRedis.Ping().Result()
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Redis - Connected")
	}
}

// Redis Getter for redis var
func Redis() *redis.Client {
	return rds
}
