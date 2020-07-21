// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalizeruntime

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opGetPersonalizedRanking = "GetPersonalizedRanking"

// GetPersonalizedRankingRequest generates a "aws/request.Request" representing the
// client's request for the GetPersonalizedRanking operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetPersonalizedRanking for more information on using the GetPersonalizedRanking
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetPersonalizedRankingRequest method.
//    req, resp := client.GetPersonalizedRankingRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/personalize-runtime-2018-05-22/GetPersonalizedRanking
func (c *PersonalizeRuntime) GetPersonalizedRankingRequest(input *GetPersonalizedRankingInput) (req *request.Request, output *GetPersonalizedRankingOutput) {
	op := &request.Operation{
		Name:       opGetPersonalizedRanking,
		HTTPMethod: "POST",
		HTTPPath:   "/personalize-ranking",
	}

	if input == nil {
		input = &GetPersonalizedRankingInput{}
	}

	output = &GetPersonalizedRankingOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetPersonalizedRanking API operation for Amazon Personalize Runtime.
//
// Re-ranks a list of recommended items for the given user. The first item in
// the list is deemed the most likely item to be of interest to the user.
//
// The solution backing the campaign must have been created using a recipe of
// type PERSONALIZED_RANKING.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Personalize Runtime's
// API operation GetPersonalizedRanking for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInvalidInputException "InvalidInputException"
//   Provide a valid value for the field or parameter.
//
//   * ErrCodeResourceNotFoundException "ResourceNotFoundException"
//   The specified resource does not exist.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/personalize-runtime-2018-05-22/GetPersonalizedRanking
func (c *PersonalizeRuntime) GetPersonalizedRanking(input *GetPersonalizedRankingInput) (*GetPersonalizedRankingOutput, error) {
	req, out := c.GetPersonalizedRankingRequest(input)
	return out, req.Send()
}

// GetPersonalizedRankingWithContext is the same as GetPersonalizedRanking with the addition of
// the ability to pass a context and additional request options.
//
// See GetPersonalizedRanking for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PersonalizeRuntime) GetPersonalizedRankingWithContext(ctx aws.Context, input *GetPersonalizedRankingInput, opts ...request.Option) (*GetPersonalizedRankingOutput, error) {
	req, out := c.GetPersonalizedRankingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetRecommendations = "GetRecommendations"

// GetRecommendationsRequest generates a "aws/request.Request" representing the
// client's request for the GetRecommendations operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetRecommendations for more information on using the GetRecommendations
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetRecommendationsRequest method.
//    req, resp := client.GetRecommendationsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/personalize-runtime-2018-05-22/GetRecommendations
func (c *PersonalizeRuntime) GetRecommendationsRequest(input *GetRecommendationsInput) (req *request.Request, output *GetRecommendationsOutput) {
	op := &request.Operation{
		Name:       opGetRecommendations,
		HTTPMethod: "POST",
		HTTPPath:   "/recommendations",
	}

	if input == nil {
		input = &GetRecommendationsInput{}
	}

	output = &GetRecommendationsOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetRecommendations API operation for Amazon Personalize Runtime.
//
// Returns a list of recommended items. The required input depends on the recipe
// type used to create the solution backing the campaign, as follows:
//
//    * RELATED_ITEMS - itemId required, userId not used
//
//    * USER_PERSONALIZATION - itemId optional, userId required
//
// Campaigns that are backed by a solution created using a recipe of type PERSONALIZED_RANKING
// use the API.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Personalize Runtime's
// API operation GetRecommendations for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInvalidInputException "InvalidInputException"
//   Provide a valid value for the field or parameter.
//
//   * ErrCodeResourceNotFoundException "ResourceNotFoundException"
//   The specified resource does not exist.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/personalize-runtime-2018-05-22/GetRecommendations
func (c *PersonalizeRuntime) GetRecommendations(input *GetRecommendationsInput) (*GetRecommendationsOutput, error) {
	req, out := c.GetRecommendationsRequest(input)
	return out, req.Send()
}

// GetRecommendationsWithContext is the same as GetRecommendations with the addition of
// the ability to pass a context and additional request options.
//
// See GetRecommendations for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PersonalizeRuntime) GetRecommendationsWithContext(ctx aws.Context, input *GetRecommendationsInput, opts ...request.Option) (*GetRecommendationsOutput, error) {
	req, out := c.GetRecommendationsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetPersonalizedRankingInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the campaign to use for generating the
	// personalized ranking.
	//
	// CampaignArn is a required field
	CampaignArn *string `locationName:"campaignArn" type:"string" required:"true"`

	// A list of items (itemId's) to rank. If an item was not included in the training
	// dataset, the item is appended to the end of the reranked list.
	//
	// InputList is a required field
	InputList []*string `locationName:"inputList" type:"list" required:"true"`

	// The user for which you want the campaign to provide a personalized ranking.
	//
	// UserId is a required field
	UserId *string `locationName:"userId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPersonalizedRankingInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetPersonalizedRankingInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPersonalizedRankingInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetPersonalizedRankingInput"}
	if s.CampaignArn == nil {
		invalidParams.Add(request.NewErrParamRequired("CampaignArn"))
	}
	if s.InputList == nil {
		invalidParams.Add(request.NewErrParamRequired("InputList"))
	}
	if s.UserId == nil {
		invalidParams.Add(request.NewErrParamRequired("UserId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCampaignArn sets the CampaignArn field's value.
func (s *GetPersonalizedRankingInput) SetCampaignArn(v string) *GetPersonalizedRankingInput {
	s.CampaignArn = &v
	return s
}

// SetInputList sets the InputList field's value.
func (s *GetPersonalizedRankingInput) SetInputList(v []*string) *GetPersonalizedRankingInput {
	s.InputList = v
	return s
}

// SetUserId sets the UserId field's value.
func (s *GetPersonalizedRankingInput) SetUserId(v string) *GetPersonalizedRankingInput {
	s.UserId = &v
	return s
}

type GetPersonalizedRankingOutput struct {
	_ struct{} `type:"structure"`

	// A list of items in order of most likely interest to the user.
	PersonalizedRanking []*PredictedItem `locationName:"personalizedRanking" type:"list"`
}

// String returns the string representation
func (s GetPersonalizedRankingOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetPersonalizedRankingOutput) GoString() string {
	return s.String()
}

// SetPersonalizedRanking sets the PersonalizedRanking field's value.
func (s *GetPersonalizedRankingOutput) SetPersonalizedRanking(v []*PredictedItem) *GetPersonalizedRankingOutput {
	s.PersonalizedRanking = v
	return s
}

type GetRecommendationsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the campaign to use for getting recommendations.
	//
	// CampaignArn is a required field
	CampaignArn *string `locationName:"campaignArn" type:"string" required:"true"`

	// The item ID to provide recommendations for.
	//
	// Required for RELATED_ITEMS recipe type.
	ItemId *string `locationName:"itemId" type:"string"`

	// The number of results to return. The default is 25. The maximum is 100.
	NumResults *int64 `locationName:"numResults" type:"integer"`

	// The user ID to provide recommendations for.
	//
	// Required for USER_PERSONALIZATION recipe type.
	UserId *string `locationName:"userId" type:"string"`
}

// String returns the string representation
func (s GetRecommendationsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRecommendationsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRecommendationsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetRecommendationsInput"}
	if s.CampaignArn == nil {
		invalidParams.Add(request.NewErrParamRequired("CampaignArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCampaignArn sets the CampaignArn field's value.
func (s *GetRecommendationsInput) SetCampaignArn(v string) *GetRecommendationsInput {
	s.CampaignArn = &v
	return s
}

// SetItemId sets the ItemId field's value.
func (s *GetRecommendationsInput) SetItemId(v string) *GetRecommendationsInput {
	s.ItemId = &v
	return s
}

// SetNumResults sets the NumResults field's value.
func (s *GetRecommendationsInput) SetNumResults(v int64) *GetRecommendationsInput {
	s.NumResults = &v
	return s
}

// SetUserId sets the UserId field's value.
func (s *GetRecommendationsInput) SetUserId(v string) *GetRecommendationsInput {
	s.UserId = &v
	return s
}

type GetRecommendationsOutput struct {
	_ struct{} `type:"structure"`

	// A list of recommendations.
	ItemList []*PredictedItem `locationName:"itemList" type:"list"`
}

// String returns the string representation
func (s GetRecommendationsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRecommendationsOutput) GoString() string {
	return s.String()
}

// SetItemList sets the ItemList field's value.
func (s *GetRecommendationsOutput) SetItemList(v []*PredictedItem) *GetRecommendationsOutput {
	s.ItemList = v
	return s
}

// An object that identifies an item.
//
// The and APIs return a list of PredictedItems.
type PredictedItem struct {
	_ struct{} `type:"structure"`

	// The recommended item ID.
	ItemId *string `locationName:"itemId" type:"string"`
}

// String returns the string representation
func (s PredictedItem) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PredictedItem) GoString() string {
	return s.String()
}

// SetItemId sets the ItemId field's value.
func (s *PredictedItem) SetItemId(v string) *PredictedItem {
	s.ItemId = &v
	return s
}
