package qualitycheck

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

// UpdateRuleForAnt invokes the qualitycheck.UpdateRuleForAnt API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/updateruleforant.html
func (client *Client) UpdateRuleForAnt(request *UpdateRuleForAntRequest) (response *UpdateRuleForAntResponse, err error) {
	response = CreateUpdateRuleForAntResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateRuleForAntWithChan invokes the qualitycheck.UpdateRuleForAnt API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/updateruleforant.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateRuleForAntWithChan(request *UpdateRuleForAntRequest) (<-chan *UpdateRuleForAntResponse, <-chan error) {
	responseChan := make(chan *UpdateRuleForAntResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateRuleForAnt(request)
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

// UpdateRuleForAntWithCallback invokes the qualitycheck.UpdateRuleForAnt API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/updateruleforant.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateRuleForAntWithCallback(request *UpdateRuleForAntRequest, callback func(response *UpdateRuleForAntResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateRuleForAntResponse
		var err error
		defer close(result)
		response, err = client.UpdateRuleForAnt(request)
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

// UpdateRuleForAntRequest is the request struct for api UpdateRuleForAnt
type UpdateRuleForAntRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// UpdateRuleForAntResponse is the response struct for api UpdateRuleForAnt
type UpdateRuleForAntResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Success   bool                   `json:"Success" xml:"Success"`
	Code      string                 `json:"Code" xml:"Code"`
	Message   string                 `json:"Message" xml:"Message"`
	Data      DataInUpdateRuleForAnt `json:"Data" xml:"Data"`
}

// CreateUpdateRuleForAntRequest creates a request to invoke UpdateRuleForAnt API
func CreateUpdateRuleForAntRequest() (request *UpdateRuleForAntRequest) {
	request = &UpdateRuleForAntRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "UpdateRuleForAnt", "", "")
	return
}

// CreateUpdateRuleForAntResponse creates a response to parse from UpdateRuleForAnt response
func CreateUpdateRuleForAntResponse() (response *UpdateRuleForAntResponse) {
	response = &UpdateRuleForAntResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
