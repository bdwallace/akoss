package green

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

// DeleteSimilarityImage invokes the green.DeleteSimilarityImage API synchronously
func (client *Client) DeleteSimilarityImage(request *DeleteSimilarityImageRequest) (response *DeleteSimilarityImageResponse, err error) {
	response = CreateDeleteSimilarityImageResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSimilarityImageWithChan invokes the green.DeleteSimilarityImage API asynchronously
func (client *Client) DeleteSimilarityImageWithChan(request *DeleteSimilarityImageRequest) (<-chan *DeleteSimilarityImageResponse, <-chan error) {
	responseChan := make(chan *DeleteSimilarityImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSimilarityImage(request)
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

// DeleteSimilarityImageWithCallback invokes the green.DeleteSimilarityImage API asynchronously
func (client *Client) DeleteSimilarityImageWithCallback(request *DeleteSimilarityImageRequest, callback func(response *DeleteSimilarityImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSimilarityImageResponse
		var err error
		defer close(result)
		response, err = client.DeleteSimilarityImage(request)
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

// DeleteSimilarityImageRequest is the request struct for api DeleteSimilarityImage
type DeleteSimilarityImageRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// DeleteSimilarityImageResponse is the response struct for api DeleteSimilarityImage
type DeleteSimilarityImageResponse struct {
	*responses.BaseResponse
}

// CreateDeleteSimilarityImageRequest creates a request to invoke DeleteSimilarityImage API
func CreateDeleteSimilarityImageRequest() (request *DeleteSimilarityImageRequest) {
	request = &DeleteSimilarityImageRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "DeleteSimilarityImage", "/green/similarity/image/delete", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteSimilarityImageResponse creates a response to parse from DeleteSimilarityImage response
func CreateDeleteSimilarityImageResponse() (response *DeleteSimilarityImageResponse) {
	response = &DeleteSimilarityImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
