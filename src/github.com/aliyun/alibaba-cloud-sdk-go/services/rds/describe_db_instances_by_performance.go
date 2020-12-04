package rds

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

// DescribeDBInstancesByPerformance invokes the rds.DescribeDBInstancesByPerformance API synchronously
// api document: https://help.aliyun.com/api/rds/describedbinstancesbyperformance.html
func (client *Client) DescribeDBInstancesByPerformance(request *DescribeDBInstancesByPerformanceRequest) (response *DescribeDBInstancesByPerformanceResponse, err error) {
	response = CreateDescribeDBInstancesByPerformanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstancesByPerformanceWithChan invokes the rds.DescribeDBInstancesByPerformance API asynchronously
// api document: https://help.aliyun.com/api/rds/describedbinstancesbyperformance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstancesByPerformanceWithChan(request *DescribeDBInstancesByPerformanceRequest) (<-chan *DescribeDBInstancesByPerformanceResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstancesByPerformanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstancesByPerformance(request)
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

// DescribeDBInstancesByPerformanceWithCallback invokes the rds.DescribeDBInstancesByPerformance API asynchronously
// api document: https://help.aliyun.com/api/rds/describedbinstancesbyperformance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstancesByPerformanceWithCallback(request *DescribeDBInstancesByPerformanceRequest, callback func(response *DescribeDBInstancesByPerformanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstancesByPerformanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstancesByPerformance(request)
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

// DescribeDBInstancesByPerformanceRequest is the request struct for api DescribeDBInstancesByPerformance
type DescribeDBInstancesByPerformanceRequest struct {
	*requests.RpcRequest
	Tag4Value            string           `position:"Query" name:"Tag.4.value"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Tag2Key              string           `position:"Query" name:"Tag.2.key"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Tag3Key              string           `position:"Query" name:"Tag.3.key"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	Tag1Value            string           `position:"Query" name:"Tag.1.value"`
	SortKey              string           `position:"Query" name:"SortKey"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	Tag3Value            string           `position:"Query" name:"Tag.3.value"`
	ProxyId              string           `position:"Query" name:"proxyId"`
	Tag5Key              string           `position:"Query" name:"Tag.5.key"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tag5Value            string           `position:"Query" name:"Tag.5.value"`
	Tags                 string           `position:"Query" name:"Tags"`
	Tag1Key              string           `position:"Query" name:"Tag.1.key"`
	SortMethod           string           `position:"Query" name:"SortMethod"`
	Tag2Value            string           `position:"Query" name:"Tag.2.value"`
	Tag4Key              string           `position:"Query" name:"Tag.4.key"`
}

// DescribeDBInstancesByPerformanceResponse is the response struct for api DescribeDBInstancesByPerformance
type DescribeDBInstancesByPerformanceResponse struct {
	*responses.BaseResponse
	RequestId        string                                  `json:"RequestId" xml:"RequestId"`
	PageNumber       int                                     `json:"PageNumber" xml:"PageNumber"`
	TotalRecordCount int                                     `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageRecordCount  int                                     `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            ItemsInDescribeDBInstancesByPerformance `json:"Items" xml:"Items"`
}

// CreateDescribeDBInstancesByPerformanceRequest creates a request to invoke DescribeDBInstancesByPerformance API
func CreateDescribeDBInstancesByPerformanceRequest() (request *DescribeDBInstancesByPerformanceRequest) {
	request = &DescribeDBInstancesByPerformanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstancesByPerformance", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBInstancesByPerformanceResponse creates a response to parse from DescribeDBInstancesByPerformance response
func CreateDescribeDBInstancesByPerformanceResponse() (response *DescribeDBInstancesByPerformanceResponse) {
	response = &DescribeDBInstancesByPerformanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
