package imm

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

// FindImages invokes the imm.FindImages API synchronously
func (client *Client) FindImages(request *FindImagesRequest) (response *FindImagesResponse, err error) {
	response = CreateFindImagesResponse()
	err = client.DoAction(request, response)
	return
}

// FindImagesWithChan invokes the imm.FindImages API asynchronously
func (client *Client) FindImagesWithChan(request *FindImagesRequest) (<-chan *FindImagesResponse, <-chan error) {
	responseChan := make(chan *FindImagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FindImages(request)
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

// FindImagesWithCallback invokes the imm.FindImages API asynchronously
func (client *Client) FindImagesWithCallback(request *FindImagesRequest, callback func(response *FindImagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FindImagesResponse
		var err error
		defer close(result)
		response, err = client.FindImages(request)
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

// FindImagesRequest is the request struct for api FindImages
type FindImagesRequest struct {
	*requests.RpcRequest
	RemarksArrayBIn          string           `position:"Query" name:"RemarksArrayBIn"`
	Project                  string           `position:"Query" name:"Project"`
	ExternalId               string           `position:"Query" name:"ExternalId"`
	FacesModifyTimeRange     string           `position:"Query" name:"FacesModifyTimeRange"`
	OCRContentsMatch         string           `position:"Query" name:"OCRContentsMatch"`
	Limit                    requests.Integer `position:"Query" name:"Limit"`
	RemarksDPrefix           string           `position:"Query" name:"RemarksDPrefix"`
	SourceType               string           `position:"Query" name:"SourceType"`
	Order                    string           `position:"Query" name:"Order"`
	GroupId                  string           `position:"Query" name:"GroupId"`
	OrderBy                  string           `position:"Query" name:"OrderBy"`
	TagNames                 string           `position:"Query" name:"TagNames"`
	Marker                   string           `position:"Query" name:"Marker"`
	RemarksCPrefix           string           `position:"Query" name:"RemarksCPrefix"`
	ModifyTimeRange          string           `position:"Query" name:"ModifyTimeRange"`
	AddressLineContentsMatch string           `position:"Query" name:"AddressLineContentsMatch"`
	Gender                   string           `position:"Query" name:"Gender"`
	RemarksArrayAIn          string           `position:"Query" name:"RemarksArrayAIn"`
	ImageSizeRange           string           `position:"Query" name:"ImageSizeRange"`
	RemarksBPrefix           string           `position:"Query" name:"RemarksBPrefix"`
	LocationBoundary         string           `position:"Query" name:"LocationBoundary"`
	ImageTimeRange           string           `position:"Query" name:"ImageTimeRange"`
	TagsModifyTimeRange      string           `position:"Query" name:"TagsModifyTimeRange"`
	AgeRange                 string           `position:"Query" name:"AgeRange"`
	RemarksAPrefix           string           `position:"Query" name:"RemarksAPrefix"`
	SourceUriPrefix          string           `position:"Query" name:"SourceUriPrefix"`
	Emotion                  string           `position:"Query" name:"Emotion"`
	CreateTimeRange          string           `position:"Query" name:"CreateTimeRange"`
	SetId                    string           `position:"Query" name:"SetId"`
}

// FindImagesResponse is the response struct for api FindImages
type FindImagesResponse struct {
	*responses.BaseResponse
	SetId      string       `json:"SetId" xml:"SetId"`
	NextMarker string       `json:"NextMarker" xml:"NextMarker"`
	RequestId  string       `json:"RequestId" xml:"RequestId"`
	Images     []ImagesItem `json:"Images" xml:"Images"`
}

// CreateFindImagesRequest creates a request to invoke FindImages API
func CreateFindImagesRequest() (request *FindImagesRequest) {
	request = &FindImagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "FindImages", "imm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateFindImagesResponse creates a response to parse from FindImages response
func CreateFindImagesResponse() (response *FindImagesResponse) {
	response = &FindImagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
