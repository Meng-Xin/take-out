package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"take-out/config"
	"take-out/logger"
)

var (
	Config *config.AllConfig // 全局Config
	Log    logger.ILog
	DB     *gorm.DB
	Redis  *redis.Client
)
