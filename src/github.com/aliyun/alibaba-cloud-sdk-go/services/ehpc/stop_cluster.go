package ehpc

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

// StopCluster invokes the ehpc.StopCluster API synchronously
// api document: https://help.aliyun.com/api/ehpc/stopcluster.html
func (client *Client) StopCluster(request *StopClusterRequest) (response *StopClusterResponse, err error) {
	response = CreateStopClusterResponse()
	err = client.DoAction(request, response)
	return
}

// StopClusterWithChan invokes the ehpc.StopCluster API asynchronously
// api document: https://help.aliyun.com/api/ehpc/stopcluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopClusterWithChan(request *StopClusterRequest) (<-chan *StopClusterResponse, <-chan error) {
	responseChan := make(chan *StopClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopCluster(request)
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

// StopClusterWithCallback invokes the ehpc.StopCluster API asynchronously
// api document: https://help.aliyun.com/api/ehpc/stopcluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopClusterWithCallback(request *StopClusterRequest, callback func(response *StopClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopClusterResponse
		var err error
		defer close(result)
		response, err = client.StopCluster(request)
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

// StopClusterRequest is the request struct for api StopCluster
type StopClusterRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

// StopClusterResponse is the response struct for api StopCluster
type StopClusterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

// CreateStopClusterRequest creates a request to invoke StopCluster API
func CreateStopClusterRequest() (request *StopClusterRequest) {
	request = &StopClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "StopCluster", "", "")
	request.Method = requests.GET
	return
}

// CreateStopClusterResponse creates a response to parse from StopCluster response
func CreateStopClusterResponse() (response *StopClusterResponse) {
	response = &StopClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
