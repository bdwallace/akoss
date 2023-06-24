package airec

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

// CreateMix invokes the airec.CreateMix API synchronously
// api document: https://help.aliyun.com/api/airec/createmix.html
func (client *Client) CreateMix(request *CreateMixRequest) (response *CreateMixResponse, err error) {
	response = CreateCreateMixResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMixWithChan invokes the airec.CreateMix API asynchronously
// api document: https://help.aliyun.com/api/airec/createmix.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMixWithChan(request *CreateMixRequest) (<-chan *CreateMixResponse, <-chan error) {
	responseChan := make(chan *CreateMixResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMix(request)
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

// CreateMixWithCallback invokes the airec.CreateMix API asynchronously
// api document: https://help.aliyun.com/api/airec/createmix.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMixWithCallback(request *CreateMixRequest, callback func(response *CreateMixResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMixResponse
		var err error
		defer close(result)
		response, err = client.CreateMix(request)
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

// CreateMixRequest is the request struct for api CreateMix
type CreateMixRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"InstanceId"`
}

// CreateMixResponse is the response struct for api CreateMix
type CreateMixResponse struct {
	*responses.BaseResponse
	RequestId string            `json:"RequestId" xml:"RequestId"`
	Code      string            `json:"Code" xml:"Code"`
	Message   string            `json:"Message" xml:"Message"`
	Result    ResultInCreateMix `json:"Result" xml:"Result"`
}

// CreateCreateMixRequest creates a request to invoke CreateMix API
func CreateCreateMixRequest() (request *CreateMixRequest) {
	request = &CreateMixRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2018-10-12", "CreateMix", "/openapi/instances/[InstanceId]/mixes", "airec", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateMixResponse creates a response to parse from CreateMix response
func CreateCreateMixResponse() (response *CreateMixResponse) {
	response = &CreateMixResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}