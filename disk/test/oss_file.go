package main

import (
	"bytes"
	"context"
	"disk/tool"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func fun1() (string, string, error) {
	viper := viper.Viper{}
	u, _ := url.Parse(viper.GetString("CosBucket"))
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  viper.GetString("Id"),
			SecretKey: viper.GetString("key"),
		},
	})
	key := "disk/" + tool.UUID() + "Ext"
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		return "", "", err
	}
	return key, v.UploadID, nil
}
func fun2() {
	u, _ := url.Parse("https://1-1256941246.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: os.Getenv("AKIDxlLl199wBGJB1dGm6HKE0fGm8LVbNvFS"), // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: os.Getenv("Ogt0xiKpBOSL6f7Yw68xnUyFLFE7mMvO"), // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	key := "disk/7.jpg"
	UploadID := "16864972954d26048722708904e0611293fa02fd74521516eb74207fc5abb012cee33a21eb"
	by, _ := os.ReadFile("./0.chunk")
	// opt 可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, 1, bytes.NewReader(by), nil,
	)
	if err != nil {
		panic(err)
	}
	PartETag := resp.Header.Get("ETag")
	fmt.Println(PartETag)
}
func fun3() {
	u, _ := url.Parse("https://1-1256941246.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("AKIDxlLl199wBGJB1dGm6HKE0fGm8LVbNvFS"), // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: os.Getenv("Ogt0xiKpBOSL6f7Yw68xnUyFLFE7mMvO"),     // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	// 1.通过响应体获取对象
	name := "disk/b5a44a96-09cc-11ee-93bb-8cc84be8ee13.apk"
	resp, err := c.Object.Get(context.Background(), name, nil)
	if err != nil {
		panic(err)
	}
	bs, _ := ioutil.ReadAll(resp.Body)

	resp.Body.Close()
	fmt.Printf("%s\n", string(bs))
	// 2.获取对象到本地文件
	_, err = c.Object.GetToFile(context.Background(), name, "exampleobject", nil)
	if err != nil {
		panic(err)
	}
}
func main() {
	fun3()
}
