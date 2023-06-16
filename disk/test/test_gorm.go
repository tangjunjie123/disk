package main

import (
	"disk/sql"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func main() {
	viper.SetConfigName("application") // 配置文件名 (不带扩展格式)
	viper.SetConfigType("yaml")        // 如果你的配置文件没有写扩展名，那么这里需要声明你的配置文件属于什么格式
	viper.AddConfigPath("./config/")   // 配置文件的路径

	err := viper.ReadInConfig() //找到并读取配置文件
	if err != nil {             // 捕获读取中遇到的error
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: lg})
	if err != nil {
		panic("failed to connect mysql.")
	}
	db.AutoMigrate(sql.ShareBasic{})
}
