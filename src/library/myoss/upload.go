package myoss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func Upload(remote string, local string) (err error) {
	var Oss BaseOss
	Oss.OssInit()

	client, err := oss.New(Oss.Endpoint, Oss.AccessKeyId, Oss.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 获取存储空间。
	bucket, err := client.Bucket(Oss.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 上传本地文件。
	err = bucket.PutObjectFromFile(remote, local)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	return
}

