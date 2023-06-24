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

// StartImageVulScan invokes the sas.StartImageVulScan API synchronously
func (client *Client) StartImageVulScan(request *StartImageVulScanRequest) (response *StartImageVulScanResponse, err error) {
	response = CreateStartImageVulScanResponse()
	err = client.DoAction(request, response)
	return
}

// StartImageVulScanWithChan invokes the sas.StartImageVulScan API asynchronously
func (client *Client) StartImageVulScanWithChan(request *StartImageVulScanRequest) (<-chan *StartImageVulScanResponse, <-chan error) {
	responseChan := make(chan *StartImageVulScanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartImageVulScan(request)
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

// StartImageVulScanWithCallback invokes the sas.StartImageVulScan API asynchronously
func (client *Client) StartImageVulScanWithCallback(request *StartImageVulScanRequest, callback func(response *StartImageVulScanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartImageVulScanResponse
		var err error
		defer close(result)
		response, err = client.StartImageVulScan(request)
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

// StartImageVulScanRequest is the request struct for api StartImageVulScan
type StartImageVulScanRequest struct {
	*requests.RpcRequest
	RepoId         string `position:"Query" name:"RepoId"`
	RepoNamespace  string `position:"Query" name:"RepoNamespace"`
	SourceIp       string `position:"Query" name:"SourceIp"`
	ImageDigest    string `position:"Query" name:"ImageDigest"`
	RepName        string `position:"Query" name:"RepName"`
	Lang           string `position:"Query" name:"Lang"`
	ImageTag       string `position:"Query" name:"ImageTag"`
	RepoInstanceId string `position:"Query" name:"RepoInstanceId"`
	ImageLayer     string `position:"Query" name:"ImageLayer"`
	RepoRegionId   string `position:"Query" name:"RepoRegionId"`
}

// StartImageVulScanResponse is the response struct for api StartImageVulScan
type StartImageVulScanResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartImageVulScanRequest creates a request to invoke StartImageVulScan API
func CreateStartImageVulScanRequest() (request *StartImageVulScanRequest) {
	request = &StartImageVulScanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "StartImageVulScan", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartImageVulScanResponse creates a response to parse from StartImageVulScan response
func CreateStartImageVulScanResponse() (response *StartImageVulScanResponse) {
	response = &StartImageVulScanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
