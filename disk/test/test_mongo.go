package main

import (
	"context"
	"disk/sql"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func main() {
	client, err := mongo.Connect(sql.Ctx,
		options.Client().
			// 连接地址
			ApplyURI("mongodb://123.56.9.154:27017").
			// 设置验证参数
			SetAuth(
				options.Credential{
					// 用户名
					Username: "admin",
					// 密码
					Password: "Tjj@2002",
				}))
	if err != nil {
		fmt.Println(err)
	}
	err = client.Ping(sql.Ctx, readpref.Primary())
	if err != nil {
		log.Println(err)
		return
	}

	collection := client.Database("test").Collection("test")
	one, err := collection.InsertOne(context.Background(), &Demo{
		Id:   "1",
		Name: "lomtom",
	})
	fmt.Println(one)
	err = client.Disconnect(sql.Ctx)
	if err != nil {
		log.Println(err)
	}
}

type Demo struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}
