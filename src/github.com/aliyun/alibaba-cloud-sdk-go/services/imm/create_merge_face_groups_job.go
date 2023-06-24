package imm

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

// CreateMergeFaceGroupsJob invokes the imm.CreateMergeFaceGroupsJob API synchronously
func (client *Client) CreateMergeFaceGroupsJob(request *CreateMergeFaceGroupsJobRequest) (response *CreateMergeFaceGroupsJobResponse, err error) {
	response = CreateCreateMergeFaceGroupsJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMergeFaceGroupsJobWithChan invokes the imm.CreateMergeFaceGroupsJob API asynchronously
func (client *Client) CreateMergeFaceGroupsJobWithChan(request *CreateMergeFaceGroupsJobRequest) (<-chan *CreateMergeFaceGroupsJobResponse, <-chan error) {
	responseChan := make(chan *CreateMergeFaceGroupsJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMergeFaceGroupsJob(request)
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

// CreateMergeFaceGroupsJobWithCallback invokes the imm.CreateMergeFaceGroupsJob API asynchronously
func (client *Client) CreateMergeFaceGroupsJobWithCallback(request *CreateMergeFaceGroupsJobRequest, callback func(response *CreateMergeFaceGroupsJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMergeFaceGroupsJobResponse
		var err error
		defer close(result)
		response, err = client.CreateMergeFaceGroupsJob(request)
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

// CreateMergeFaceGroupsJobRequest is the request struct for api CreateMergeFaceGroupsJob
type CreateMergeFaceGroupsJobRequest struct {
	*requests.RpcRequest
	Project         string `position:"Query" name:"Project"`
	NotifyEndpoint  string `position:"Query" name:"NotifyEndpoint"`
	GroupIdFrom     string `position:"Query" name:"GroupIdFrom"`
	NotifyTopicName string `position:"Query" name:"NotifyTopicName"`
	GroupIdTo       string `position:"Query" name:"GroupIdTo"`
	SetId           string `position:"Query" name:"SetId"`
}

// CreateMergeFaceGroupsJobResponse is the response struct for api CreateMergeFaceGroupsJob
type CreateMergeFaceGroupsJobResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	JobId       string `json:"JobId" xml:"JobId"`
	SetId       string `json:"SetId" xml:"SetId"`
	JobType     string `json:"JobType" xml:"JobType"`
	GroupIdTo   string `json:"GroupIdTo" xml:"GroupIdTo"`
	GroupIdFrom string `json:"GroupIdFrom" xml:"GroupIdFrom"`
}

// CreateCreateMergeFaceGroupsJobRequest creates a request to invoke CreateMergeFaceGroupsJob API
func CreateCreateMergeFaceGroupsJobRequest() (request *CreateMergeFaceGroupsJobRequest) {
	request = &CreateMergeFaceGroupsJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "CreateMergeFaceGroupsJob", "imm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateMergeFaceGroupsJobResponse creates a response to parse from CreateMergeFaceGroupsJob response
func CreateCreateMergeFaceGroupsJobResponse() (response *CreateMergeFaceGroupsJobResponse) {
	response = &CreateMergeFaceGroupsJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}