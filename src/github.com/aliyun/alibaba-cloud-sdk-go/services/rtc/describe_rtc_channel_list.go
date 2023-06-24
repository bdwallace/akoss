package rtc

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

// DescribeRtcChannelList invokes the rtc.DescribeRtcChannelList API synchronously
func (client *Client) DescribeRtcChannelList(request *DescribeRtcChannelListRequest) (response *DescribeRtcChannelListResponse, err error) {
	response = CreateDescribeRtcChannelListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRtcChannelListWithChan invokes the rtc.DescribeRtcChannelList API asynchronously
func (client *Client) DescribeRtcChannelListWithChan(request *DescribeRtcChannelListRequest) (<-chan *DescribeRtcChannelListResponse, <-chan error) {
	responseChan := make(chan *DescribeRtcChannelListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRtcChannelList(request)
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

// DescribeRtcChannelListWithCallback invokes the rtc.DescribeRtcChannelList API asynchronously
func (client *Client) DescribeRtcChannelListWithCallback(request *DescribeRtcChannelListRequest, callback func(response *DescribeRtcChannelListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRtcChannelListResponse
		var err error
		defer close(result)
		response, err = client.DescribeRtcChannelList(request)
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

// DescribeRtcChannelListRequest is the request struct for api DescribeRtcChannelList
type DescribeRtcChannelListRequest struct {
	*requests.RpcRequest
	SortType    string           `position:"Query" name:"SortType"`
	UserId      string           `position:"Query" name:"UserId"`
	ServiceArea string           `position:"Query" name:"ServiceArea"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	ShowLog     string           `position:"Query" name:"ShowLog"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	PageNo      requests.Integer `position:"Query" name:"PageNo"`
	AppId       string           `position:"Query" name:"AppId"`
	ChannelId   string           `position:"Query" name:"ChannelId"`
	TimePoint   string           `position:"Query" name:"TimePoint"`
}

// DescribeRtcChannelListResponse is the response struct for api DescribeRtcChannelList
type DescribeRtcChannelListResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	PageSize    int64       `json:"PageSize" xml:"PageSize"`
	PageNo      int64       `json:"PageNo" xml:"PageNo"`
	TotalCnt    int64       `json:"TotalCnt" xml:"TotalCnt"`
	ChannelList ChannelList `json:"ChannelList" xml:"ChannelList"`
}

// CreateDescribeRtcChannelListRequest creates a request to invoke DescribeRtcChannelList API
func CreateDescribeRtcChannelListRequest() (request *DescribeRtcChannelListRequest) {
	request = &DescribeRtcChannelListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "DescribeRtcChannelList", "rtc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRtcChannelListResponse creates a response to parse from DescribeRtcChannelList response
func CreateDescribeRtcChannelListResponse() (response *DescribeRtcChannelListResponse) {
	response = &DescribeRtcChannelListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
