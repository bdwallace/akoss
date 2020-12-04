package sas

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

// DescribeImageVulList invokes the sas.DescribeImageVulList API synchronously
func (client *Client) DescribeImageVulList(request *DescribeImageVulListRequest) (response *DescribeImageVulListResponse, err error) {
	response = CreateDescribeImageVulListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeImageVulListWithChan invokes the sas.DescribeImageVulList API asynchronously
func (client *Client) DescribeImageVulListWithChan(request *DescribeImageVulListRequest) (<-chan *DescribeImageVulListResponse, <-chan error) {
	responseChan := make(chan *DescribeImageVulListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeImageVulList(request)
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

// DescribeImageVulListWithCallback invokes the sas.DescribeImageVulList API asynchronously
func (client *Client) DescribeImageVulListWithCallback(request *DescribeImageVulListRequest, callback func(response *DescribeImageVulListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeImageVulListResponse
		var err error
		defer close(result)
		response, err = client.DescribeImageVulList(request)
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

// DescribeImageVulListRequest is the request struct for api DescribeImageVulList
type DescribeImageVulListRequest struct {
	*requests.RpcRequest
	RepoId        string           `position:"Query" name:"RepoId"`
	StatusList    string           `position:"Query" name:"StatusList"`
	CveId         string           `position:"Query" name:"CveId"`
	Remark        string           `position:"Query" name:"Remark"`
	Type          string           `position:"Query" name:"Type"`
	CreateTsStart requests.Integer `position:"Query" name:"CreateTsStart"`
	SourceIp      string           `position:"Query" name:"SourceIp"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	Digest        string           `position:"Query" name:"Digest"`
	ModifyTsStart requests.Integer `position:"Query" name:"ModifyTsStart"`
	Tag           string           `position:"Query" name:"Tag"`
	Lang          string           `position:"Query" name:"Lang"`
	ModifyTsEnd   requests.Integer `position:"Query" name:"ModifyTsEnd"`
	Level         string           `position:"Query" name:"Level"`
	Resource      string           `position:"Query" name:"Resource"`
	GroupId       string           `position:"Query" name:"GroupId"`
	Dealed        string           `position:"Query" name:"Dealed"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	BatchName     string           `position:"Query" name:"BatchName"`
	AliasName     string           `position:"Query" name:"AliasName"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	RepoName      string           `position:"Query" name:"RepoName"`
	Name          string           `position:"Query" name:"Name"`
	Ids           string           `position:"Query" name:"Ids"`
	CreateTsEnd   requests.Integer `position:"Query" name:"CreateTsEnd"`
	Necessity     string           `position:"Query" name:"Necessity"`
	Uuids         string           `position:"Query" name:"Uuids"`
}

// DescribeImageVulListResponse is the response struct for api DescribeImageVulList
type DescribeImageVulListResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	PageSize    int         `json:"PageSize" xml:"PageSize"`
	CurrentPage int         `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount  int         `json:"TotalCount" xml:"TotalCount"`
	VulRecords  []VulRecord `json:"VulRecords" xml:"VulRecords"`
}

// CreateDescribeImageVulListRequest creates a request to invoke DescribeImageVulList API
func CreateDescribeImageVulListRequest() (request *DescribeImageVulListRequest) {
	request = &DescribeImageVulListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeImageVulList", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeImageVulListResponse creates a response to parse from DescribeImageVulList response
func CreateDescribeImageVulListResponse() (response *DescribeImageVulListResponse) {
	response = &DescribeImageVulListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
