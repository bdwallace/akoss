package scdn

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

// DescribeScdnDomainBpsData invokes the scdn.DescribeScdnDomainBpsData API synchronously
func (client *Client) DescribeScdnDomainBpsData(request *DescribeScdnDomainBpsDataRequest) (response *DescribeScdnDomainBpsDataResponse, err error) {
	response = CreateDescribeScdnDomainBpsDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnDomainBpsDataWithChan invokes the scdn.DescribeScdnDomainBpsData API asynchronously
func (client *Client) DescribeScdnDomainBpsDataWithChan(request *DescribeScdnDomainBpsDataRequest) (<-chan *DescribeScdnDomainBpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnDomainBpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnDomainBpsData(request)
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

// DescribeScdnDomainBpsDataWithCallback invokes the scdn.DescribeScdnDomainBpsData API asynchronously
func (client *Client) DescribeScdnDomainBpsDataWithCallback(request *DescribeScdnDomainBpsDataRequest, callback func(response *DescribeScdnDomainBpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnDomainBpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnDomainBpsData(request)
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

// DescribeScdnDomainBpsDataRequest is the request struct for api DescribeScdnDomainBpsData
type DescribeScdnDomainBpsDataRequest struct {
	*requests.RpcRequest
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	Interval       string           `position:"Query" name:"Interval"`
}

// DescribeScdnDomainBpsDataResponse is the response struct for api DescribeScdnDomainBpsData
type DescribeScdnDomainBpsDataResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	DomainName         string             `json:"DomainName" xml:"DomainName"`
	StartTime          string             `json:"StartTime" xml:"StartTime"`
	EndTime            string             `json:"EndTime" xml:"EndTime"`
	DataInterval       string             `json:"DataInterval" xml:"DataInterval"`
	BpsDataPerInterval BpsDataPerInterval `json:"BpsDataPerInterval" xml:"BpsDataPerInterval"`
}

// CreateDescribeScdnDomainBpsDataRequest creates a request to invoke DescribeScdnDomainBpsData API
func CreateDescribeScdnDomainBpsDataRequest() (request *DescribeScdnDomainBpsDataRequest) {
	request = &DescribeScdnDomainBpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnDomainBpsData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeScdnDomainBpsDataResponse creates a response to parse from DescribeScdnDomainBpsData response
func CreateDescribeScdnDomainBpsDataResponse() (response *DescribeScdnDomainBpsDataResponse) {
	response = &DescribeScdnDomainBpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
