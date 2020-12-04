package green

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

// VideoCancelScan invokes the green.VideoCancelScan API synchronously
func (client *Client) VideoCancelScan(request *VideoCancelScanRequest) (response *VideoCancelScanResponse, err error) {
	response = CreateVideoCancelScanResponse()
	err = client.DoAction(request, response)
	return
}

// VideoCancelScanWithChan invokes the green.VideoCancelScan API asynchronously
func (client *Client) VideoCancelScanWithChan(request *VideoCancelScanRequest) (<-chan *VideoCancelScanResponse, <-chan error) {
	responseChan := make(chan *VideoCancelScanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VideoCancelScan(request)
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

// VideoCancelScanWithCallback invokes the green.VideoCancelScan API asynchronously
func (client *Client) VideoCancelScanWithCallback(request *VideoCancelScanRequest, callback func(response *VideoCancelScanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VideoCancelScanResponse
		var err error
		defer close(result)
		response, err = client.VideoCancelScan(request)
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

// VideoCancelScanRequest is the request struct for api VideoCancelScan
type VideoCancelScanRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// VideoCancelScanResponse is the response struct for api VideoCancelScan
type VideoCancelScanResponse struct {
	*responses.BaseResponse
}

// CreateVideoCancelScanRequest creates a request to invoke VideoCancelScan API
func CreateVideoCancelScanRequest() (request *VideoCancelScanRequest) {
	request = &VideoCancelScanRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "VideoCancelScan", "/green/video/cancelscan", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateVideoCancelScanResponse creates a response to parse from VideoCancelScan response
func CreateVideoCancelScanResponse() (response *VideoCancelScanResponse) {
	response = &VideoCancelScanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
