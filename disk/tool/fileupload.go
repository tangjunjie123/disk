package tool

import (
	"bytes"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
)

func FileuploadTxoss(r *http.Request) string {
	u, _ := url.Parse("https://1-1256941246.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("AKIDxlLl199wBGJB1dGm6HKE0fGm8LVbNvFS"), // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: os.Getenv("Ogt0xiKpBOSL6f7Yw68xnUyFLFE7mMvO"),     // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	file, header, _ := r.FormFile("file")

	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	name := "disk/" + UUID() + path.Ext(header.Filename)
	// 1.通过字符串上传对象
	var err error
	// 2.通过本地文件上传对象
	_, err = c.Object.Put(context.Background(), name, file, nil)
	if err != nil {
		panic(err)
	}
	return "https://1-1256941246.cos.ap-nanjing.myqcloud.com" + name
}

func InitPart(Ext string) (string, string, error) {
	u, _ := url.Parse("https://1-1256941246.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDxlLl199wBGJB1dGm6HKE0fGm8LVbNvFS",
			SecretKey: "Ogt0xiKpBOSL6f7Yw68xnUyFLFE7mMvO",
		},
	})
	key := "disk/" + UUID() + Ext
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		return "", "", err
	}
	return key, v.UploadID, nil
}
func PartUpload(r *http.Request) (string, error) {
	u, _ := url.Parse("https://1-1256941246.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDxlLl199wBGJB1dGm6HKE0fGm8LVbNvFS",
			SecretKey: "Ogt0xiKpBOSL6f7Yw68xnUyFLFE7mMvO",
		},
	})
	key := r.PostForm.Get("key")
	UploadID := r.PostForm.Get("upload_id")
	partNumber, err := strconv.Atoi(r.PostForm.Get("part_number"))
	if err != nil {
		return "", err
	}
	f, _, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, f)

	// opt可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, partNumber, bytes.NewReader(buf.Bytes()), nil,
	)
	if err != nil {
		return "", err
	}
	return strings.Trim(resp.Header.Get("ETag"), "\""), nil
}
func CosPartUploadComplete(key, uploadId string, co []cos.Object) error {
	u, _ := url.Parse("https://1-1256941246.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  ("AKIDxlLl199wBGJB1dGm6HKE0fGm8LVbNvFS"),
			SecretKey: ("Ogt0xiKpBOSL6f7Yw68xnUyFLFE7mMvO"),
		},
	})

	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, co...)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, uploadId, opt,
	)
	return err
}
