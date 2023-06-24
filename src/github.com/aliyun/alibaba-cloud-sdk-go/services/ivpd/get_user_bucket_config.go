package ivpd

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetUserBucketConfig invokes the ivpd.GetUserBucketConfig API synchronously
// api document: https://help.aliyun.com/api/ivpd/getuserbucketconfig.html
func (client *Client) GetUserBucketConfig(request *GetUserBucketConfigRequest) (response *GetUserBucketConfigResponse, err error) {
	response = CreateGetUserBucketConfigResponse()
	err = client.DoAction(request, response)
	return
}

// GetUserBucketConfigWithChan invokes the ivpd.GetUserBucketConfig API asynchronously
// api document: https://help.aliyun.com/api/ivpd/getuserbucketconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetUserBucketConfigWithChan(request *GetUserBucketConfigRequest) (<-chan *GetUserBucketConfigResponse, <-chan error) {
	responseChan := make(chan *GetUserBucketConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUserBucketConfig(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetUserBucketConfigWithCallback invokes the ivpd.GetUserBucketConfig API asynchronously
// api document: https://help.aliyun.com/api/ivpd/getuserbucketconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetUserBucketConfigWithCallback(request *GetUserBucketConfigRequest, callback func(response *GetUserBucketConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUserBucketConfigResponse
		var err error
		defer close(result)
		response, err = client.GetUserBucketConfig(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// GetUserBucketConfigRequest is the request struct for api GetUserBucketConfig
type GetUserBucketConfigRequest struct {
	*requests.RpcRequest
}

// GetUserBucketConfigResponse is the response struct for api GetUserBucketConfig
type GetUserBucketConfigResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Code      string     `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateGetUserBucketConfigRequest creates a request to invoke GetUserBucketConfig API
func CreateGetUserBucketConfigRequest() (request *GetUserBucketConfigRequest) {
	request = &GetUserBucketConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivpd", "2019-06-25", "GetUserBucketConfig", "ivpd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetUserBucketConfigResponse creates a response to parse from GetUserBucketConfig response
func CreateGetUserBucketConfigResponse() (response *GetUserBucketConfigResponse) {
	response = &GetUserBucketConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}