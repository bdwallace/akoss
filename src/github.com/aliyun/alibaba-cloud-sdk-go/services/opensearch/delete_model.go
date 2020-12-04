package opensearch

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

// DeleteModel invokes the opensearch.DeleteModel API synchronously
func (client *Client) DeleteModel(request *DeleteModelRequest) (response *DeleteModelResponse, err error) {
	response = CreateDeleteModelResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteModelWithChan invokes the opensearch.DeleteModel API asynchronously
func (client *Client) DeleteModelWithChan(request *DeleteModelRequest) (<-chan *DeleteModelResponse, <-chan error) {
	responseChan := make(chan *DeleteModelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteModel(request)
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

// DeleteModelWithCallback invokes the opensearch.DeleteModel API asynchronously
func (client *Client) DeleteModelWithCallback(request *DeleteModelRequest, callback func(response *DeleteModelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteModelResponse
		var err error
		defer close(result)
		response, err = client.DeleteModel(request)
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

// DeleteModelRequest is the request struct for api DeleteModel
type DeleteModelRequest struct {
	*requests.RoaRequest
	ModelName        string `position:"Path" name:"modelName"`
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// DeleteModelResponse is the response struct for api DeleteModel
type DeleteModelResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    string `json:"result" xml:"result"`
}

// CreateDeleteModelRequest creates a request to invoke DeleteModel API
func CreateDeleteModelRequest() (request *DeleteModelRequest) {
	request = &DeleteModelRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "DeleteModel", "/v4/openapi/app-groups/[appGroupIdentity]/algorithm/models/[modelName]", "opensearch", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteModelResponse creates a response to parse from DeleteModel response
func CreateDeleteModelResponse() (response *DeleteModelResponse) {
	response = &DeleteModelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
