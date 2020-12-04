package vs

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

// DescribeVsCertificateList invokes the vs.DescribeVsCertificateList API synchronously
func (client *Client) DescribeVsCertificateList(request *DescribeVsCertificateListRequest) (response *DescribeVsCertificateListResponse, err error) {
	response = CreateDescribeVsCertificateListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVsCertificateListWithChan invokes the vs.DescribeVsCertificateList API asynchronously
func (client *Client) DescribeVsCertificateListWithChan(request *DescribeVsCertificateListRequest) (<-chan *DescribeVsCertificateListResponse, <-chan error) {
	responseChan := make(chan *DescribeVsCertificateListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVsCertificateList(request)
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

// DescribeVsCertificateListWithCallback invokes the vs.DescribeVsCertificateList API asynchronously
func (client *Client) DescribeVsCertificateListWithCallback(request *DescribeVsCertificateListRequest, callback func(response *DescribeVsCertificateListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVsCertificateListResponse
		var err error
		defer close(result)
		response, err = client.DescribeVsCertificateList(request)
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

// DescribeVsCertificateListRequest is the request struct for api DescribeVsCertificateList
type DescribeVsCertificateListRequest struct {
	*requests.RpcRequest
	ShowLog    string           `position:"Query" name:"ShowLog"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeVsCertificateListResponse is the response struct for api DescribeVsCertificateList
type DescribeVsCertificateListResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	CertificateListModel CertificateListModel `json:"CertificateListModel" xml:"CertificateListModel"`
}

// CreateDescribeVsCertificateListRequest creates a request to invoke DescribeVsCertificateList API
func CreateDescribeVsCertificateListRequest() (request *DescribeVsCertificateListRequest) {
	request = &DescribeVsCertificateListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeVsCertificateList", "vs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVsCertificateListResponse creates a response to parse from DescribeVsCertificateList response
func CreateDescribeVsCertificateListResponse() (response *DescribeVsCertificateListResponse) {
	response = &DescribeVsCertificateListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
