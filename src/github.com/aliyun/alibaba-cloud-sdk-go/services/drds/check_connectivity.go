package drds

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

// CheckConnectivity invokes the drds.CheckConnectivity API synchronously
func (client *Client) CheckConnectivity(request *CheckConnectivityRequest) (response *CheckConnectivityResponse, err error) {
	response = CreateCheckConnectivityResponse()
	err = client.DoAction(request, response)
	return
}

// CheckConnectivityWithChan invokes the drds.CheckConnectivity API asynchronously
func (client *Client) CheckConnectivityWithChan(request *CheckConnectivityRequest) (<-chan *CheckConnectivityResponse, <-chan error) {
	responseChan := make(chan *CheckConnectivityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckConnectivity(request)
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

// CheckConnectivityWithCallback invokes the drds.CheckConnectivity API asynchronously
func (client *Client) CheckConnectivityWithCallback(request *CheckConnectivityRequest, callback func(response *CheckConnectivityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckConnectivityResponse
		var err error
		defer close(result)
		response, err = client.CheckConnectivity(request)
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

// CheckConnectivityRequest is the request struct for api CheckConnectivity
type CheckConnectivityRequest struct {
	*requests.RpcRequest
	DbInfo string           `position:"Query" name:"DbInfo"`
	DbType requests.Integer `position:"Query" name:"DbType"`
}

// CheckConnectivityResponse is the response struct for api CheckConnectivity
type CheckConnectivityResponse struct {
	*responses.BaseResponse
	RequestId               string                  `json:"RequestId" xml:"RequestId"`
	Success                 bool                    `json:"Success" xml:"Success"`
	CheckConnectivityResult CheckConnectivityResult `json:"CheckConnectivityResult" xml:"CheckConnectivityResult"`
}

// CreateCheckConnectivityRequest creates a request to invoke CheckConnectivity API
func CreateCheckConnectivityRequest() (request *CheckConnectivityRequest) {
	request = &CheckConnectivityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "CheckConnectivity", "Drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckConnectivityResponse creates a response to parse from CheckConnectivity response
func CreateCheckConnectivityResponse() (response *CheckConnectivityResponse) {
	response = &CheckConnectivityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}