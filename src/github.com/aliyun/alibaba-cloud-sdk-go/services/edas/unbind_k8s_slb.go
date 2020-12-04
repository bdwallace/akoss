package edas

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

// UnbindK8sSlb invokes the edas.UnbindK8sSlb API synchronously
func (client *Client) UnbindK8sSlb(request *UnbindK8sSlbRequest) (response *UnbindK8sSlbResponse, err error) {
	response = CreateUnbindK8sSlbResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindK8sSlbWithChan invokes the edas.UnbindK8sSlb API asynchronously
func (client *Client) UnbindK8sSlbWithChan(request *UnbindK8sSlbRequest) (<-chan *UnbindK8sSlbResponse, <-chan error) {
	responseChan := make(chan *UnbindK8sSlbResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindK8sSlb(request)
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

// UnbindK8sSlbWithCallback invokes the edas.UnbindK8sSlb API asynchronously
func (client *Client) UnbindK8sSlbWithCallback(request *UnbindK8sSlbRequest, callback func(response *UnbindK8sSlbResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindK8sSlbResponse
		var err error
		defer close(result)
		response, err = client.UnbindK8sSlb(request)
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

// UnbindK8sSlbRequest is the request struct for api UnbindK8sSlb
type UnbindK8sSlbRequest struct {
	*requests.RoaRequest
	AppId     string `position:"Query" name:"AppId"`
	ClusterId string `position:"Query" name:"ClusterId"`
	Type      string `position:"Query" name:"Type"`
}

// UnbindK8sSlbResponse is the response struct for api UnbindK8sSlb
type UnbindK8sSlbResponse struct {
	*responses.BaseResponse
	Code          int    `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	ChangeOrderId string `json:"ChangeOrderId" xml:"ChangeOrderId"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateUnbindK8sSlbRequest creates a request to invoke UnbindK8sSlb API
func CreateUnbindK8sSlbRequest() (request *UnbindK8sSlbRequest) {
	request = &UnbindK8sSlbRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "UnbindK8sSlb", "/pop/v5/k8s/acs/k8s_slb_binding", "edas", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateUnbindK8sSlbResponse creates a response to parse from UnbindK8sSlb response
func CreateUnbindK8sSlbResponse() (response *UnbindK8sSlbResponse) {
	response = &UnbindK8sSlbResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
