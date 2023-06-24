package dcdn

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

// DescribeDcdnDomainBpsData invokes the dcdn.DescribeDcdnDomainBpsData API synchronously
func (client *Client) DescribeDcdnDomainBpsData(request *DescribeDcdnDomainBpsDataRequest) (response *DescribeDcdnDomainBpsDataResponse, err error) {
	response = CreateDescribeDcdnDomainBpsDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnDomainBpsDataWithChan invokes the dcdn.DescribeDcdnDomainBpsData API asynchronously
func (client *Client) DescribeDcdnDomainBpsDataWithChan(request *DescribeDcdnDomainBpsDataRequest) (<-chan *DescribeDcdnDomainBpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnDomainBpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnDomainBpsData(request)
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

// DescribeDcdnDomainBpsDataWithCallback invokes the dcdn.DescribeDcdnDomainBpsData API asynchronously
func (client *Client) DescribeDcdnDomainBpsDataWithCallback(request *DescribeDcdnDomainBpsDataRequest, callback func(response *DescribeDcdnDomainBpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnDomainBpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnDomainBpsData(request)
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

// DescribeDcdnDomainBpsDataRequest is the request struct for api DescribeDcdnDomainBpsData
type DescribeDcdnDomainBpsDataRequest struct {
	*requests.RpcRequest
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	Interval       string           `position:"Query" name:"Interval"`
}

// DescribeDcdnDomainBpsDataResponse is the response struct for api DescribeDcdnDomainBpsData
type DescribeDcdnDomainBpsDataResponse struct {
	*responses.BaseResponse
	RequestId          string                                        `json:"RequestId" xml:"RequestId"`
	DomainName         string                                        `json:"DomainName" xml:"DomainName"`
	StartTime          string                                        `json:"StartTime" xml:"StartTime"`
	EndTime            string                                        `json:"EndTime" xml:"EndTime"`
	DataInterval       string                                        `json:"DataInterval" xml:"DataInterval"`
	BpsDataPerInterval BpsDataPerIntervalInDescribeDcdnDomainBpsData `json:"BpsDataPerInterval" xml:"BpsDataPerInterval"`
}

// CreateDescribeDcdnDomainBpsDataRequest creates a request to invoke DescribeDcdnDomainBpsData API
func CreateDescribeDcdnDomainBpsDataRequest() (request *DescribeDcdnDomainBpsDataRequest) {
	request = &DescribeDcdnDomainBpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainBpsData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnDomainBpsDataResponse creates a response to parse from DescribeDcdnDomainBpsData response
func CreateDescribeDcdnDomainBpsDataResponse() (response *DescribeDcdnDomainBpsDataResponse) {
	response = &DescribeDcdnDomainBpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
