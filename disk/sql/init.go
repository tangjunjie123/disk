package sql

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	Ctx context.Context
	Db  *gorm.DB      = Gorm_init()
	Red *redis.Client = Redis_init()

	mg *mongo.Database = mongo_init()
)

func Viper_init() {
	viper.SetConfigName("application") // 配置文件名 (不带扩展格式)
	viper.SetConfigType("yaml")        // 如果你的配置文件没有写扩展名，那么这里需要声明你的配置文件属于什么格式
	viper.AddConfigPath("./config/")   // 配置文件的路径

	err := viper.ReadInConfig() //找到并读取配置文件
	if err != nil {             // 捕获读取中遇到的error
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

func Gorm_init() *gorm.DB {
	Viper_init()

	lg := logger2.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger2.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger2.Info,
			Colorful:      true,
		})
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		viper.Get("mysql.username"),
		viper.Get("mysql.password"),
		viper.Get("mysql.host"),
		viper.Get("mysql.port"),
		viper.Get("mysql.dbname"),
		viper.Get("mysql.timeout"))
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: lg})
	if err != nil {
		panic("failed to connect mysql.")
	}
	return d
}
func Redis_init() *redis.Client {
	Viper_init()
	Redis := redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.PoolSize"),
		MinIdleConns: viper.GetInt("redis.MinIdleConn"),
	})
	Ctx = context.Background()
	return Redis
}
func mongo_init() *mongo.Database {
	Viper_init()
	client, err := mongo.Connect(Ctx,
		options.Client().
			// 连接地址//mongodb://123.56.9.154:27017
			ApplyURI("mongodb://"+viper.GetString("mongodb.addr")).
			// 设置验证参数
			SetAuth(
				options.Credential{
					// 用户名
					Username: viper.GetString("mongodb.username"),
					// 密码
					Password: viper.GetString("mongodb.password"),
				}))
	if err != nil {
		fmt.Println(err)
	}
	return client.Database(viper.GetString("mongodb.database"))
}
