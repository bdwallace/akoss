package mts

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

// DeletePipeline invokes the mts.DeletePipeline API synchronously
func (client *Client) DeletePipeline(request *DeletePipelineRequest) (response *DeletePipelineResponse, err error) {
	response = CreateDeletePipelineResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePipelineWithChan invokes the mts.DeletePipeline API asynchronously
func (client *Client) DeletePipelineWithChan(request *DeletePipelineRequest) (<-chan *DeletePipelineResponse, <-chan error) {
	responseChan := make(chan *DeletePipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePipeline(request)
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

// DeletePipelineWithCallback invokes the mts.DeletePipeline API asynchronously
func (client *Client) DeletePipelineWithCallback(request *DeletePipelineRequest, callback func(response *DeletePipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePipelineResponse
		var err error
		defer close(result)
		response, err = client.DeletePipeline(request)
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

// DeletePipelineRequest is the request struct for api DeletePipeline
type DeletePipelineRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
}

// DeletePipelineResponse is the response struct for api DeletePipeline
type DeletePipelineResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	PipelineId string `json:"PipelineId" xml:"PipelineId"`
}

// CreateDeletePipelineRequest creates a request to invoke DeletePipeline API
func CreateDeletePipelineRequest() (request *DeletePipelineRequest) {
	request = &DeletePipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "DeletePipeline", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeletePipelineResponse creates a response to parse from DeletePipeline response
func CreateDeletePipelineResponse() (response *DeletePipelineResponse) {
	response = &DeletePipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
