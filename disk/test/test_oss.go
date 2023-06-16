package main

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
)

func main() {
	u, _ := url.Parse("https://1-1256941246.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("AKIDxlLl199wBGJB1dGm6HKE0fGm8LVbNvFS"), // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: os.Getenv("Ogt0xiKpBOSL6f7Yw68xnUyFLFE7mMvO"),     // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	name := "disk/1.jpg"
	// 1.通过字符串上传对象
	var err error
	// 2.通过本地文件上传对象
	_, err = c.Object.PutFromFile(context.Background(), name, "test/img/1.jpg", nil)
	if err != nil {
		panic(err)
	}
}
