package cloudwf

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

// ListStaStatus invokes the cloudwf.ListStaStatus API synchronously
// api document: https://help.aliyun.com/api/cloudwf/liststastatus.html
func (client *Client) ListStaStatus(request *ListStaStatusRequest) (response *ListStaStatusResponse, err error) {
	response = CreateListStaStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ListStaStatusWithChan invokes the cloudwf.ListStaStatus API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/liststastatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListStaStatusWithChan(request *ListStaStatusRequest) (<-chan *ListStaStatusResponse, <-chan error) {
	responseChan := make(chan *ListStaStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListStaStatus(request)
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

// ListStaStatusWithCallback invokes the cloudwf.ListStaStatus API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/liststastatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListStaStatusWithCallback(request *ListStaStatusRequest, callback func(response *ListStaStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListStaStatusResponse
		var err error
		defer close(result)
		response, err = client.ListStaStatus(request)
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

// ListStaStatusRequest is the request struct for api ListStaStatus
type ListStaStatusRequest struct {
	*requests.RpcRequest
	OrderCol          string           `position:"Query" name:"OrderCol"`
	SearchGroupName   string           `position:"Query" name:"SearchGroupName"`
	SearchStatus      requests.Integer `position:"Query" name:"SearchStatus"`
	Length            requests.Integer `position:"Query" name:"Length"`
	SearchUsername    string           `position:"Query" name:"SearchUsername"`
	OrderDir          string           `position:"Query" name:"OrderDir"`
	SearchProtocal    string           `position:"Query" name:"SearchProtocal"`
	SearchSsid        string           `position:"Query" name:"SearchSsid"`
	SearchApName      string           `position:"Query" name:"SearchApName"`
	SearchIp          string           `position:"Query" name:"SearchIp"`
	PageIndex         requests.Integer `position:"Query" name:"PageIndex"`
	SearchMac         string           `position:"Query" name:"SearchMac"`
	SearchDescription string           `position:"Query" name:"SearchDescription"`
}

// ListStaStatusResponse is the response struct for api ListStaStatus
type ListStaStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateListStaStatusRequest creates a request to invoke ListStaStatus API
func CreateListStaStatusRequest() (request *ListStaStatusRequest) {
	request = &ListStaStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ListStaStatus", "cloudwf", "openAPI")
	return
}

// CreateListStaStatusResponse creates a response to parse from ListStaStatus response
func CreateListStaStatusResponse() (response *ListStaStatusResponse) {
	response = &ListStaStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
