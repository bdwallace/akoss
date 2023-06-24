package unimkt

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

// AccountOperate invokes the unimkt.AccountOperate API synchronously
// api document: https://help.aliyun.com/api/unimkt/accountoperate.html
func (client *Client) AccountOperate(request *AccountOperateRequest) (response *AccountOperateResponse, err error) {
	response = CreateAccountOperateResponse()
	err = client.DoAction(request, response)
	return
}

// AccountOperateWithChan invokes the unimkt.AccountOperate API asynchronously
// api document: https://help.aliyun.com/api/unimkt/accountoperate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AccountOperateWithChan(request *AccountOperateRequest) (<-chan *AccountOperateResponse, <-chan error) {
	responseChan := make(chan *AccountOperateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AccountOperate(request)
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

// AccountOperateWithCallback invokes the unimkt.AccountOperate API asynchronously
// api document: https://help.aliyun.com/api/unimkt/accountoperate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AccountOperateWithCallback(request *AccountOperateRequest, callback func(response *AccountOperateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AccountOperateResponse
		var err error
		defer close(result)
		response, err = client.AccountOperate(request)
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

// AccountOperateRequest is the request struct for api AccountOperate
type AccountOperateRequest struct {
	*requests.RpcRequest
	AccountManagerNumber string           `position:"Body" name:"AccountManagerNumber"`
	FromUserId           string           `position:"Body" name:"FromUserId"`
	ToAccountNo          string           `position:"Body" name:"ToAccountNo"`
	CataloguePrice       string           `position:"Body" name:"CataloguePrice"`
	BpId                 string           `position:"Body" name:"BpId"`
	OperateName          string           `position:"Body" name:"OperateName"`
	Balance              string           `position:"Body" name:"Balance"`
	ActualAmount         string           `position:"Body" name:"ActualAmount"`
	ProxyUserNick        string           `position:"Body" name:"ProxyUserNick"`
	FromAccountNo        string           `position:"Body" name:"FromAccountNo"`
	PriceVersion         string           `position:"Body" name:"PriceVersion"`
	CreateTime           string           `position:"Body" name:"CreateTime"`
	Amount               string           `position:"Body" name:"Amount"`
	AccountManagerName   string           `position:"Body" name:"AccountManagerName"`
	ToUserId             string           `position:"Body" name:"ToUserId"`
	MarketCount          requests.Integer `position:"Body" name:"MarketCount"`
	ProxyUserId          string           `position:"Body" name:"ProxyUserId"`
	DiscountRate         string           `position:"Body" name:"DiscountRate"`
	AccuActualAmount     string           `position:"Body" name:"AccuActualAmount"`
	FlowId               string           `position:"Body" name:"FlowId"`
	PreDebit             string           `position:"Body" name:"PreDebit"`
	AccuAmount           string           `position:"Body" name:"AccuAmount"`
}

// AccountOperateResponse is the response struct for api AccountOperate
type AccountOperateResponse struct {
	*responses.BaseResponse
	Status    int    `json:"Status" xml:"Status"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAccountOperateRequest creates a request to invoke AccountOperate API
func CreateAccountOperateRequest() (request *AccountOperateRequest) {
	request = &AccountOperateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-07", "AccountOperate", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAccountOperateResponse creates a response to parse from AccountOperate response
func CreateAccountOperateResponse() (response *AccountOperateResponse) {
	response = &AccountOperateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
