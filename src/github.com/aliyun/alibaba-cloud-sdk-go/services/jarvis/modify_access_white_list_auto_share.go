package jarvis

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

// ModifyAccessWhiteListAutoShare invokes the jarvis.ModifyAccessWhiteListAutoShare API synchronously
// api document: https://help.aliyun.com/api/jarvis/modifyaccesswhitelistautoshare.html
func (client *Client) ModifyAccessWhiteListAutoShare(request *ModifyAccessWhiteListAutoShareRequest) (response *ModifyAccessWhiteListAutoShareResponse, err error) {
	response = CreateModifyAccessWhiteListAutoShareResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyAccessWhiteListAutoShareWithChan invokes the jarvis.ModifyAccessWhiteListAutoShare API asynchronously
// api document: https://help.aliyun.com/api/jarvis/modifyaccesswhitelistautoshare.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAccessWhiteListAutoShareWithChan(request *ModifyAccessWhiteListAutoShareRequest) (<-chan *ModifyAccessWhiteListAutoShareResponse, <-chan error) {
	responseChan := make(chan *ModifyAccessWhiteListAutoShareResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyAccessWhiteListAutoShare(request)
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

// ModifyAccessWhiteListAutoShareWithCallback invokes the jarvis.ModifyAccessWhiteListAutoShare API asynchronously
// api document: https://help.aliyun.com/api/jarvis/modifyaccesswhitelistautoshare.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAccessWhiteListAutoShareWithCallback(request *ModifyAccessWhiteListAutoShareRequest, callback func(response *ModifyAccessWhiteListAutoShareResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyAccessWhiteListAutoShareResponse
		var err error
		defer close(result)
		response, err = client.ModifyAccessWhiteListAutoShare(request)
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

// ModifyAccessWhiteListAutoShareRequest is the request struct for api ModifyAccessWhiteListAutoShare
type ModifyAccessWhiteListAutoShareRequest struct {
	*requests.RpcRequest
	SrcIP         string           `position:"Query" name:"SrcIP"`
	SourceIp      string           `position:"Query" name:"SourceIp"`
	AutoConfig    requests.Integer `position:"Query" name:"AutoConfig"`
	ProductName   string           `position:"Query" name:"ProductName"`
	WhiteListType requests.Integer `position:"Query" name:"WhiteListType"`
	Lang          string           `position:"Query" name:"Lang"`
	SourceCode    string           `position:"Query" name:"SourceCode"`
}

// ModifyAccessWhiteListAutoShareResponse is the response struct for api ModifyAccessWhiteListAutoShare
type ModifyAccessWhiteListAutoShareResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Module    string `json:"Module" xml:"Module"`
}

// CreateModifyAccessWhiteListAutoShareRequest creates a request to invoke ModifyAccessWhiteListAutoShare API
func CreateModifyAccessWhiteListAutoShareRequest() (request *ModifyAccessWhiteListAutoShareRequest) {
	request = &ModifyAccessWhiteListAutoShareRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jarvis", "2018-02-06", "ModifyAccessWhiteListAutoShare", "jarvis", "openAPI")
	return
}

// CreateModifyAccessWhiteListAutoShareResponse creates a response to parse from ModifyAccessWhiteListAutoShare response
func CreateModifyAccessWhiteListAutoShareResponse() (response *ModifyAccessWhiteListAutoShareResponse) {
	response = &ModifyAccessWhiteListAutoShareResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
