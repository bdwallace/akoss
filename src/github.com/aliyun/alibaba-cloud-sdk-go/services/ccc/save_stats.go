package ccc

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

// SaveStats invokes the ccc.SaveStats API synchronously
// api document: https://help.aliyun.com/api/ccc/savestats.html
func (client *Client) SaveStats(request *SaveStatsRequest) (response *SaveStatsResponse, err error) {
	response = CreateSaveStatsResponse()
	err = client.DoAction(request, response)
	return
}

// SaveStatsWithChan invokes the ccc.SaveStats API asynchronously
// api document: https://help.aliyun.com/api/ccc/savestats.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveStatsWithChan(request *SaveStatsRequest) (<-chan *SaveStatsResponse, <-chan error) {
	responseChan := make(chan *SaveStatsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveStats(request)
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

// SaveStatsWithCallback invokes the ccc.SaveStats API asynchronously
// api document: https://help.aliyun.com/api/ccc/savestats.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveStatsWithCallback(request *SaveStatsRequest, callback func(response *SaveStatsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveStatsResponse
		var err error
		defer close(result)
		response, err = client.SaveStats(request)
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

// SaveStatsRequest is the request struct for api SaveStats
type SaveStatsRequest struct {
	*requests.RpcRequest
	CallId        string           `position:"Query" name:"CallId"`
	RecordTime    requests.Integer `position:"Query" name:"RecordTime"`
	CallStartTime requests.Integer `position:"Query" name:"CallStartTime"`
	Uid           string           `position:"Query" name:"Uid"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	Stats         string           `position:"Query" name:"Stats"`
	TenantId      string           `position:"Query" name:"TenantId"`
	CalleeNumber  string           `position:"Query" name:"CalleeNumber"`
	CallerNumber  string           `position:"Query" name:"CallerNumber"`
}

// SaveStatsResponse is the response struct for api SaveStats
type SaveStatsResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RowCount       int64  `json:"RowCount" xml:"RowCount"`
}

// CreateSaveStatsRequest creates a request to invoke SaveStats API
func CreateSaveStatsRequest() (request *SaveStatsRequest) {
	request = &SaveStatsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "SaveStats", "", "")
	return
}

// CreateSaveStatsResponse creates a response to parse from SaveStats response
func CreateSaveStatsResponse() (response *SaveStatsResponse) {
	response = &SaveStatsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
