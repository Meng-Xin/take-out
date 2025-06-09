package global

import (
	logger "github.com/Meng-Xin/logger"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"take-out/config"
)

var (
	Config *config.AllConfig // 全局Config
	Log    logger.ILog
	DB     *gorm.DB
	Redis  *redis.Client
)
