package dataworks_public

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

// GetDataServiceApi invokes the dataworks_public.GetDataServiceApi API synchronously
func (client *Client) GetDataServiceApi(request *GetDataServiceApiRequest) (response *GetDataServiceApiResponse, err error) {
	response = CreateGetDataServiceApiResponse()
	err = client.DoAction(request, response)
	return
}

// GetDataServiceApiWithChan invokes the dataworks_public.GetDataServiceApi API asynchronously
func (client *Client) GetDataServiceApiWithChan(request *GetDataServiceApiRequest) (<-chan *GetDataServiceApiResponse, <-chan error) {
	responseChan := make(chan *GetDataServiceApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDataServiceApi(request)
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

// GetDataServiceApiWithCallback invokes the dataworks_public.GetDataServiceApi API asynchronously
func (client *Client) GetDataServiceApiWithCallback(request *GetDataServiceApiRequest, callback func(response *GetDataServiceApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDataServiceApiResponse
		var err error
		defer close(result)
		response, err = client.GetDataServiceApi(request)
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

// GetDataServiceApiRequest is the request struct for api GetDataServiceApi
type GetDataServiceApiRequest struct {
	*requests.RpcRequest
	TenantId  requests.Integer `position:"Body" name:"TenantId"`
	ProjectId requests.Integer `position:"Body" name:"ProjectId"`
	ApiId     requests.Integer `position:"Body" name:"ApiId"`
}

// GetDataServiceApiResponse is the response struct for api GetDataServiceApi
type GetDataServiceApiResponse struct {
	*responses.BaseResponse
	ErrorCode      string                  `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string                  `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int                     `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string                  `json:"RequestId" xml:"RequestId"`
	Success        bool                    `json:"Success" xml:"Success"`
	Data           DataInGetDataServiceApi `json:"Data" xml:"Data"`
}

// CreateGetDataServiceApiRequest creates a request to invoke GetDataServiceApi API
func CreateGetDataServiceApiRequest() (request *GetDataServiceApiRequest) {
	request = &GetDataServiceApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetDataServiceApi", "", "")
	request.Method = requests.POST
	return
}

// CreateGetDataServiceApiResponse creates a response to parse from GetDataServiceApi response
func CreateGetDataServiceApiResponse() (response *GetDataServiceApiResponse) {
	response = &GetDataServiceApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}