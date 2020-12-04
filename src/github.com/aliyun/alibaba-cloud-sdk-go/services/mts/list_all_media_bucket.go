package mts

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

// ListAllMediaBucket invokes the mts.ListAllMediaBucket API synchronously
func (client *Client) ListAllMediaBucket(request *ListAllMediaBucketRequest) (response *ListAllMediaBucketResponse, err error) {
	response = CreateListAllMediaBucketResponse()
	err = client.DoAction(request, response)
	return
}

// ListAllMediaBucketWithChan invokes the mts.ListAllMediaBucket API asynchronously
func (client *Client) ListAllMediaBucketWithChan(request *ListAllMediaBucketRequest) (<-chan *ListAllMediaBucketResponse, <-chan error) {
	responseChan := make(chan *ListAllMediaBucketResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAllMediaBucket(request)
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

// ListAllMediaBucketWithCallback invokes the mts.ListAllMediaBucket API asynchronously
func (client *Client) ListAllMediaBucketWithCallback(request *ListAllMediaBucketRequest, callback func(response *ListAllMediaBucketResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAllMediaBucketResponse
		var err error
		defer close(result)
		response, err = client.ListAllMediaBucket(request)
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

// ListAllMediaBucketRequest is the request struct for api ListAllMediaBucket
type ListAllMediaBucketRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NextPageToken        string           `position:"Query" name:"NextPageToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	MaximumPageSize      requests.Integer `position:"Query" name:"MaximumPageSize"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ListAllMediaBucketResponse is the response struct for api ListAllMediaBucket
type ListAllMediaBucketResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	NextPageToken   string          `json:"NextPageToken" xml:"NextPageToken"`
	MediaBucketList MediaBucketList `json:"MediaBucketList" xml:"MediaBucketList"`
}

// CreateListAllMediaBucketRequest creates a request to invoke ListAllMediaBucket API
func CreateListAllMediaBucketRequest() (request *ListAllMediaBucketRequest) {
	request = &ListAllMediaBucketRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ListAllMediaBucket", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAllMediaBucketResponse creates a response to parse from ListAllMediaBucket response
func CreateListAllMediaBucketResponse() (response *ListAllMediaBucketResponse) {
	response = &ListAllMediaBucketResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
