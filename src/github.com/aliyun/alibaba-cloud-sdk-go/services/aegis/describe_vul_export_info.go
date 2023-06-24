package aegis

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

// DescribeVulExportInfo invokes the aegis.DescribeVulExportInfo API synchronously
// api document: https://help.aliyun.com/api/aegis/describevulexportinfo.html
func (client *Client) DescribeVulExportInfo(request *DescribeVulExportInfoRequest) (response *DescribeVulExportInfoResponse, err error) {
	response = CreateDescribeVulExportInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVulExportInfoWithChan invokes the aegis.DescribeVulExportInfo API asynchronously
// api document: https://help.aliyun.com/api/aegis/describevulexportinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVulExportInfoWithChan(request *DescribeVulExportInfoRequest) (<-chan *DescribeVulExportInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeVulExportInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVulExportInfo(request)
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

// DescribeVulExportInfoWithCallback invokes the aegis.DescribeVulExportInfo API asynchronously
// api document: https://help.aliyun.com/api/aegis/describevulexportinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVulExportInfoWithCallback(request *DescribeVulExportInfoRequest, callback func(response *DescribeVulExportInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVulExportInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeVulExportInfo(request)
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

// DescribeVulExportInfoRequest is the request struct for api DescribeVulExportInfo
type DescribeVulExportInfoRequest struct {
	*requests.RpcRequest
	SourceIp string           `position:"Query" name:"SourceIp"`
	ExportId requests.Integer `position:"Query" name:"ExportId"`
}

// DescribeVulExportInfoResponse is the response struct for api DescribeVulExportInfo
type DescribeVulExportInfoResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Id           int    `json:"Id" xml:"Id"`
	FileName     string `json:"FileName" xml:"FileName"`
	CurrentCount int    `json:"CurrentCount" xml:"CurrentCount"`
	TotalCount   int    `json:"TotalCount" xml:"TotalCount"`
	Progress     int    `json:"Progress" xml:"Progress"`
	ExportStatus string `json:"ExportStatus" xml:"ExportStatus"`
	Message      string `json:"Message" xml:"Message"`
	Link         string `json:"Link" xml:"Link"`
}

// CreateDescribeVulExportInfoRequest creates a request to invoke DescribeVulExportInfo API
func CreateDescribeVulExportInfoRequest() (request *DescribeVulExportInfoRequest) {
	request = &DescribeVulExportInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeVulExportInfo", "vipaegis", "openAPI")
	return
}

// CreateDescribeVulExportInfoResponse creates a response to parse from DescribeVulExportInfo response
func CreateDescribeVulExportInfoResponse() (response *DescribeVulExportInfoResponse) {
	response = &DescribeVulExportInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
