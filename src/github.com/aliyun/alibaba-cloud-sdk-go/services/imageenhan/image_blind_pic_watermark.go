package imageenhan

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

// ImageBlindPicWatermark invokes the imageenhan.ImageBlindPicWatermark API synchronously
func (client *Client) ImageBlindPicWatermark(request *ImageBlindPicWatermarkRequest) (response *ImageBlindPicWatermarkResponse, err error) {
	response = CreateImageBlindPicWatermarkResponse()
	err = client.DoAction(request, response)
	return
}

// ImageBlindPicWatermarkWithChan invokes the imageenhan.ImageBlindPicWatermark API asynchronously
func (client *Client) ImageBlindPicWatermarkWithChan(request *ImageBlindPicWatermarkRequest) (<-chan *ImageBlindPicWatermarkResponse, <-chan error) {
	responseChan := make(chan *ImageBlindPicWatermarkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImageBlindPicWatermark(request)
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

// ImageBlindPicWatermarkWithCallback invokes the imageenhan.ImageBlindPicWatermark API asynchronously
func (client *Client) ImageBlindPicWatermarkWithCallback(request *ImageBlindPicWatermarkRequest, callback func(response *ImageBlindPicWatermarkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImageBlindPicWatermarkResponse
		var err error
		defer close(result)
		response, err = client.ImageBlindPicWatermark(request)
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

// ImageBlindPicWatermarkRequest is the request struct for api ImageBlindPicWatermark
type ImageBlindPicWatermarkRequest struct {
	*requests.RpcRequest
	WatermarkImageURL string           `position:"Body" name:"WatermarkImageURL"`
	QualityFactor     requests.Integer `position:"Body" name:"QualityFactor"`
	FunctionType      string           `position:"Body" name:"FunctionType"`
	LogoURL           string           `position:"Body" name:"LogoURL"`
	OutputFileType    string           `position:"Body" name:"OutputFileType"`
	OriginImageURL    string           `position:"Body" name:"OriginImageURL"`
}

// ImageBlindPicWatermarkResponse is the response struct for api ImageBlindPicWatermark
type ImageBlindPicWatermarkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateImageBlindPicWatermarkRequest creates a request to invoke ImageBlindPicWatermark API
func CreateImageBlindPicWatermarkRequest() (request *ImageBlindPicWatermarkRequest) {
	request = &ImageBlindPicWatermarkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imageenhan", "2019-09-30", "ImageBlindPicWatermark", "imageenhan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateImageBlindPicWatermarkResponse creates a response to parse from ImageBlindPicWatermark response
func CreateImageBlindPicWatermarkResponse() (response *ImageBlindPicWatermarkResponse) {
	response = &ImageBlindPicWatermarkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
