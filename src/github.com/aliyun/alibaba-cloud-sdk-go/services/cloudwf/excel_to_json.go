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

// ExcelToJson invokes the cloudwf.ExcelToJson API synchronously
// api document: https://help.aliyun.com/api/cloudwf/exceltojson.html
func (client *Client) ExcelToJson(request *ExcelToJsonRequest) (response *ExcelToJsonResponse, err error) {
	response = CreateExcelToJsonResponse()
	err = client.DoAction(request, response)
	return
}

// ExcelToJsonWithChan invokes the cloudwf.ExcelToJson API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/exceltojson.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ExcelToJsonWithChan(request *ExcelToJsonRequest) (<-chan *ExcelToJsonResponse, <-chan error) {
	responseChan := make(chan *ExcelToJsonResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExcelToJson(request)
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

// ExcelToJsonWithCallback invokes the cloudwf.ExcelToJson API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/exceltojson.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ExcelToJsonWithCallback(request *ExcelToJsonRequest, callback func(response *ExcelToJsonResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExcelToJsonResponse
		var err error
		defer close(result)
		response, err = client.ExcelToJson(request)
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

// ExcelToJsonRequest is the request struct for api ExcelToJson
type ExcelToJsonRequest struct {
	*requests.RpcRequest
	UploadData string `position:"Query" name:"UploadData"`
}

// ExcelToJsonResponse is the response struct for api ExcelToJson
type ExcelToJsonResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateExcelToJsonRequest creates a request to invoke ExcelToJson API
func CreateExcelToJsonRequest() (request *ExcelToJsonRequest) {
	request = &ExcelToJsonRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ExcelToJson", "cloudwf", "openAPI")
	return
}

// CreateExcelToJsonResponse creates a response to parse from ExcelToJson response
func CreateExcelToJsonResponse() (response *ExcelToJsonResponse) {
	response = &ExcelToJsonResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
