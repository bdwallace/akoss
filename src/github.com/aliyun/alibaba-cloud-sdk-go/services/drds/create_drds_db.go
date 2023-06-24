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

// CreateDrdsDB invokes the drds.CreateDrdsDB API synchronously
func (client *Client) CreateDrdsDB(request *CreateDrdsDBRequest) (response *CreateDrdsDBResponse, err error) {
	response = CreateCreateDrdsDBResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDrdsDBWithChan invokes the drds.CreateDrdsDB API asynchronously
func (client *Client) CreateDrdsDBWithChan(request *CreateDrdsDBRequest) (<-chan *CreateDrdsDBResponse, <-chan error) {
	responseChan := make(chan *CreateDrdsDBResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDrdsDB(request)
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

// CreateDrdsDBWithCallback invokes the drds.CreateDrdsDB API asynchronously
func (client *Client) CreateDrdsDBWithCallback(request *CreateDrdsDBRequest, callback func(response *CreateDrdsDBResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDrdsDBResponse
		var err error
		defer close(result)
		response, err = client.CreateDrdsDB(request)
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

// CreateDrdsDBRequest is the request struct for api CreateDrdsDB
type CreateDrdsDBRequest struct {
	*requests.RpcRequest
	Encode         string `position:"Query" name:"Encode"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	Password       string `position:"Query" name:"Password"`
	DbName         string `position:"Query" name:"DbName"`
	RdsInstances   string `position:"Query" name:"RdsInstances"`
}

// CreateDrdsDBResponse is the response struct for api CreateDrdsDB
type CreateDrdsDBResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateDrdsDBRequest creates a request to invoke CreateDrdsDB API
func CreateCreateDrdsDBRequest() (request *CreateDrdsDBRequest) {
	request = &CreateDrdsDBRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2017-10-16", "CreateDrdsDB", "Drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDrdsDBResponse creates a response to parse from CreateDrdsDB response
func CreateCreateDrdsDBResponse() (response *CreateDrdsDBResponse) {
	response = &CreateDrdsDBResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}