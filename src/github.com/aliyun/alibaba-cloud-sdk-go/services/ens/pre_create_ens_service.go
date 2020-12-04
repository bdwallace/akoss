package ens

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

// PreCreateEnsService invokes the ens.PreCreateEnsService API synchronously
func (client *Client) PreCreateEnsService(request *PreCreateEnsServiceRequest) (response *PreCreateEnsServiceResponse, err error) {
	response = CreatePreCreateEnsServiceResponse()
	err = client.DoAction(request, response)
	return
}

// PreCreateEnsServiceWithChan invokes the ens.PreCreateEnsService API asynchronously
func (client *Client) PreCreateEnsServiceWithChan(request *PreCreateEnsServiceRequest) (<-chan *PreCreateEnsServiceResponse, <-chan error) {
	responseChan := make(chan *PreCreateEnsServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PreCreateEnsService(request)
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

// PreCreateEnsServiceWithCallback invokes the ens.PreCreateEnsService API asynchronously
func (client *Client) PreCreateEnsServiceWithCallback(request *PreCreateEnsServiceRequest, callback func(response *PreCreateEnsServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PreCreateEnsServiceResponse
		var err error
		defer close(result)
		response, err = client.PreCreateEnsService(request)
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

// PreCreateEnsServiceRequest is the request struct for api PreCreateEnsService
type PreCreateEnsServiceRequest struct {
	*requests.RpcRequest
	BandwidthType           string `position:"Query" name:"BandwidthType"`
	SchedulingPriceStrategy string `position:"Query" name:"SchedulingPriceStrategy"`
	ImageId                 string `position:"Query" name:"ImageId"`
	InstanceSpec            string `position:"Query" name:"InstanceSpec"`
	KeyPairName             string `position:"Query" name:"KeyPairName"`
	UserData                string `position:"Query" name:"UserData"`
	Password                string `position:"Query" name:"Password"`
	BuyResourcesDetail      string `position:"Query" name:"BuyResourcesDetail"`
	SystemDiskSize          string `position:"Query" name:"SystemDiskSize"`
	InstanceBandwithdLimit  string `position:"Query" name:"InstanceBandwithdLimit"`
	EnsServiceName          string `position:"Query" name:"EnsServiceName"`
	Version                 string `position:"Query" name:"Version"`
	NetLevel                string `position:"Query" name:"NetLevel"`
	SchedulingStrategy      string `position:"Query" name:"SchedulingStrategy"`
	DataDiskSize            string `position:"Query" name:"DataDiskSize"`
}

// PreCreateEnsServiceResponse is the response struct for api PreCreateEnsService
type PreCreateEnsServiceResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	Code               int    `json:"Code" xml:"Code"`
	EnsServiceId       string `json:"EnsServiceId" xml:"EnsServiceId"`
	NetLevel           string `json:"NetLevel" xml:"NetLevel"`
	BuyResourcesDetail string `json:"BuyResourcesDetail" xml:"BuyResourcesDetail"`
}

// CreatePreCreateEnsServiceRequest creates a request to invoke PreCreateEnsService API
func CreatePreCreateEnsServiceRequest() (request *PreCreateEnsServiceRequest) {
	request = &PreCreateEnsServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "PreCreateEnsService", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePreCreateEnsServiceResponse creates a response to parse from PreCreateEnsService response
func CreatePreCreateEnsServiceResponse() (response *PreCreateEnsServiceResponse) {
	response = &PreCreateEnsServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
