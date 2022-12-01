package global

import (
	"GinDemo_v1/config"
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 常量
const (
	ConfigFile = "./config.yaml" // 配置文件
)

// 变量
var (
	GvaConfig        config.ServerConfig // 全局配置
	GvaLogger        *zap.Logger         // 日志
	GvaMysqlClient   *gorm.DB            //Mysql客户端
	GvaMongoDbClient *mongo.Client       //MongoDb客户端
	GvaRedis         *redis.Client       //Redis客户端
	GvaElastic       *elastic.Client     // ES客户端
)
