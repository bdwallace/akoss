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

// DescribeEnsNetDistrict invokes the ens.DescribeEnsNetDistrict API synchronously
func (client *Client) DescribeEnsNetDistrict(request *DescribeEnsNetDistrictRequest) (response *DescribeEnsNetDistrictResponse, err error) {
	response = CreateDescribeEnsNetDistrictResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEnsNetDistrictWithChan invokes the ens.DescribeEnsNetDistrict API asynchronously
func (client *Client) DescribeEnsNetDistrictWithChan(request *DescribeEnsNetDistrictRequest) (<-chan *DescribeEnsNetDistrictResponse, <-chan error) {
	responseChan := make(chan *DescribeEnsNetDistrictResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEnsNetDistrict(request)
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

// DescribeEnsNetDistrictWithCallback invokes the ens.DescribeEnsNetDistrict API asynchronously
func (client *Client) DescribeEnsNetDistrictWithCallback(request *DescribeEnsNetDistrictRequest, callback func(response *DescribeEnsNetDistrictResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEnsNetDistrictResponse
		var err error
		defer close(result)
		response, err = client.DescribeEnsNetDistrict(request)
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

// DescribeEnsNetDistrictRequest is the request struct for api DescribeEnsNetDistrict
type DescribeEnsNetDistrictRequest struct {
	*requests.RpcRequest
	NetLevelCode    string `position:"Query" name:"NetLevelCode"`
	NetDistrictCode string `position:"Query" name:"NetDistrictCode"`
	Version         string `position:"Query" name:"Version"`
}

// DescribeEnsNetDistrictResponse is the response struct for api DescribeEnsNetDistrict
type DescribeEnsNetDistrictResponse struct {
	*responses.BaseResponse
	RequestId       string                                  `json:"RequestId" xml:"RequestId"`
	Code            int                                     `json:"Code" xml:"Code"`
	EnsNetDistricts EnsNetDistrictsInDescribeEnsNetDistrict `json:"EnsNetDistricts" xml:"EnsNetDistricts"`
}

// CreateDescribeEnsNetDistrictRequest creates a request to invoke DescribeEnsNetDistrict API
func CreateDescribeEnsNetDistrictRequest() (request *DescribeEnsNetDistrictRequest) {
	request = &DescribeEnsNetDistrictRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeEnsNetDistrict", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeEnsNetDistrictResponse creates a response to parse from DescribeEnsNetDistrict response
func CreateDescribeEnsNetDistrictResponse() (response *DescribeEnsNetDistrictResponse) {
	response = &DescribeEnsNetDistrictResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
