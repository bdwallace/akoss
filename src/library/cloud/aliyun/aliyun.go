package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

func Balance(access_key_id, access_secret string) (response *bssopenapi.QueryAccountBalanceResponse, err error) {
	// client, err := bssopenapi.NewClientWithAccessKey("cn-hangzhou", "LTAI4GGTbhZwzMee2rsxJJyH", "zEkTGAX3D5Iz9LPLV96wJYYhVSEKDc")
	client, err := bssopenapi.NewClientWithAccessKey("cn-hangzhou", access_key_id, access_secret)

	request := bssopenapi.CreateQueryAccountBalanceRequest()
	request.Scheme = "https"

	response, err = client.QueryAccountBalance(request)

	return
}
