package elasticsearch

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

// UpgradeEngineVersion invokes the elasticsearch.UpgradeEngineVersion API synchronously
func (client *Client) UpgradeEngineVersion(request *UpgradeEngineVersionRequest) (response *UpgradeEngineVersionResponse, err error) {
	response = CreateUpgradeEngineVersionResponse()
	err = client.DoAction(request, response)
	return
}

// UpgradeEngineVersionWithChan invokes the elasticsearch.UpgradeEngineVersion API asynchronously
func (client *Client) UpgradeEngineVersionWithChan(request *UpgradeEngineVersionRequest) (<-chan *UpgradeEngineVersionResponse, <-chan error) {
	responseChan := make(chan *UpgradeEngineVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpgradeEngineVersion(request)
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

// UpgradeEngineVersionWithCallback invokes the elasticsearch.UpgradeEngineVersion API asynchronously
func (client *Client) UpgradeEngineVersionWithCallback(request *UpgradeEngineVersionRequest, callback func(response *UpgradeEngineVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpgradeEngineVersionResponse
		var err error
		defer close(result)
		response, err = client.UpgradeEngineVersion(request)
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

// UpgradeEngineVersionRequest is the request struct for api UpgradeEngineVersion
type UpgradeEngineVersionRequest struct {
	*requests.RoaRequest
	InstanceId  string           `position:"Path" name:"InstanceId"`
	DryRun      requests.Boolean `position:"Query" name:"dryRun"`
	ClientToken string           `position:"Query" name:"clientToken"`
}

// UpgradeEngineVersionResponse is the response struct for api UpgradeEngineVersion
type UpgradeEngineVersionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateUpgradeEngineVersionRequest creates a request to invoke UpgradeEngineVersion API
func CreateUpgradeEngineVersionRequest() (request *UpgradeEngineVersionRequest) {
	request = &UpgradeEngineVersionRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "UpgradeEngineVersion", "/openapi/instances/[InstanceId]/actions/upgrade-version", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpgradeEngineVersionResponse creates a response to parse from UpgradeEngineVersion response
func CreateUpgradeEngineVersionResponse() (response *UpgradeEngineVersionResponse) {
	response = &UpgradeEngineVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}