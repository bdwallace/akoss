package live

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

// DescribeLiveStreamsPublishList invokes the live.DescribeLiveStreamsPublishList API synchronously
func (client *Client) DescribeLiveStreamsPublishList(request *DescribeLiveStreamsPublishListRequest) (response *DescribeLiveStreamsPublishListResponse, err error) {
	response = CreateDescribeLiveStreamsPublishListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveStreamsPublishListWithChan invokes the live.DescribeLiveStreamsPublishList API asynchronously
func (client *Client) DescribeLiveStreamsPublishListWithChan(request *DescribeLiveStreamsPublishListRequest) (<-chan *DescribeLiveStreamsPublishListResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamsPublishListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamsPublishList(request)
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

// DescribeLiveStreamsPublishListWithCallback invokes the live.DescribeLiveStreamsPublishList API asynchronously
func (client *Client) DescribeLiveStreamsPublishListWithCallback(request *DescribeLiveStreamsPublishListRequest, callback func(response *DescribeLiveStreamsPublishListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamsPublishListResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamsPublishList(request)
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

// DescribeLiveStreamsPublishListRequest is the request struct for api DescribeLiveStreamsPublishList
type DescribeLiveStreamsPublishListRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	AppName    string           `position:"Query" name:"AppName"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	StreamName string           `position:"Query" name:"StreamName"`
	QueryType  string           `position:"Query" name:"QueryType"`
	StreamType string           `position:"Query" name:"StreamType"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OrderBy    string           `position:"Query" name:"OrderBy"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveStreamsPublishListResponse is the response struct for api DescribeLiveStreamsPublishList
type DescribeLiveStreamsPublishListResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	PageNum     int         `json:"PageNum" xml:"PageNum"`
	PageSize    int         `json:"PageSize" xml:"PageSize"`
	TotalNum    int         `json:"TotalNum" xml:"TotalNum"`
	TotalPage   int         `json:"TotalPage" xml:"TotalPage"`
	PublishInfo PublishInfo `json:"PublishInfo" xml:"PublishInfo"`
}

// CreateDescribeLiveStreamsPublishListRequest creates a request to invoke DescribeLiveStreamsPublishList API
func CreateDescribeLiveStreamsPublishListRequest() (request *DescribeLiveStreamsPublishListRequest) {
	request = &DescribeLiveStreamsPublishListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamsPublishList", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveStreamsPublishListResponse creates a response to parse from DescribeLiveStreamsPublishList response
func CreateDescribeLiveStreamsPublishListResponse() (response *DescribeLiveStreamsPublishListResponse) {
	response = &DescribeLiveStreamsPublishListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
