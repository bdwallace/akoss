// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package mediaconnectiface provides an interface to enable mocking the AWS MediaConnect service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mediaconnectiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/mediaconnect"
)

// MediaConnectAPI provides an interface to enable mocking the
// mediaconnect.MediaConnect service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS MediaConnect.
//    func myFunc(svc mediaconnectiface.MediaConnectAPI) bool {
//        // Make svc.AddFlowOutputs request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := mediaconnect.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMediaConnectClient struct {
//        mediaconnectiface.MediaConnectAPI
//    }
//    func (m *mockMediaConnectClient) AddFlowOutputs(input *mediaconnect.AddFlowOutputsInput) (*mediaconnect.AddFlowOutputsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMediaConnectClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type MediaConnectAPI interface {
	AddFlowOutputs(*mediaconnect.AddFlowOutputsInput) (*mediaconnect.AddFlowOutputsOutput, error)
	AddFlowOutputsWithContext(aws.Context, *mediaconnect.AddFlowOutputsInput, ...request.Option) (*mediaconnect.AddFlowOutputsOutput, error)
	AddFlowOutputsRequest(*mediaconnect.AddFlowOutputsInput) (*request.Request, *mediaconnect.AddFlowOutputsOutput)

	CreateFlow(*mediaconnect.CreateFlowInput) (*mediaconnect.CreateFlowOutput, error)
	CreateFlowWithContext(aws.Context, *mediaconnect.CreateFlowInput, ...request.Option) (*mediaconnect.CreateFlowOutput, error)
	CreateFlowRequest(*mediaconnect.CreateFlowInput) (*request.Request, *mediaconnect.CreateFlowOutput)

	DeleteFlow(*mediaconnect.DeleteFlowInput) (*mediaconnect.DeleteFlowOutput, error)
	DeleteFlowWithContext(aws.Context, *mediaconnect.DeleteFlowInput, ...request.Option) (*mediaconnect.DeleteFlowOutput, error)
	DeleteFlowRequest(*mediaconnect.DeleteFlowInput) (*request.Request, *mediaconnect.DeleteFlowOutput)

	DescribeFlow(*mediaconnect.DescribeFlowInput) (*mediaconnect.DescribeFlowOutput, error)
	DescribeFlowWithContext(aws.Context, *mediaconnect.DescribeFlowInput, ...request.Option) (*mediaconnect.DescribeFlowOutput, error)
	DescribeFlowRequest(*mediaconnect.DescribeFlowInput) (*request.Request, *mediaconnect.DescribeFlowOutput)

	GrantFlowEntitlements(*mediaconnect.GrantFlowEntitlementsInput) (*mediaconnect.GrantFlowEntitlementsOutput, error)
	GrantFlowEntitlementsWithContext(aws.Context, *mediaconnect.GrantFlowEntitlementsInput, ...request.Option) (*mediaconnect.GrantFlowEntitlementsOutput, error)
	GrantFlowEntitlementsRequest(*mediaconnect.GrantFlowEntitlementsInput) (*request.Request, *mediaconnect.GrantFlowEntitlementsOutput)

	ListEntitlements(*mediaconnect.ListEntitlementsInput) (*mediaconnect.ListEntitlementsOutput, error)
	ListEntitlementsWithContext(aws.Context, *mediaconnect.ListEntitlementsInput, ...request.Option) (*mediaconnect.ListEntitlementsOutput, error)
	ListEntitlementsRequest(*mediaconnect.ListEntitlementsInput) (*request.Request, *mediaconnect.ListEntitlementsOutput)

	ListEntitlementsPages(*mediaconnect.ListEntitlementsInput, func(*mediaconnect.ListEntitlementsOutput, bool) bool) error
	ListEntitlementsPagesWithContext(aws.Context, *mediaconnect.ListEntitlementsInput, func(*mediaconnect.ListEntitlementsOutput, bool) bool, ...request.Option) error

	ListFlows(*mediaconnect.ListFlowsInput) (*mediaconnect.ListFlowsOutput, error)
	ListFlowsWithContext(aws.Context, *mediaconnect.ListFlowsInput, ...request.Option) (*mediaconnect.ListFlowsOutput, error)
	ListFlowsRequest(*mediaconnect.ListFlowsInput) (*request.Request, *mediaconnect.ListFlowsOutput)

	ListFlowsPages(*mediaconnect.ListFlowsInput, func(*mediaconnect.ListFlowsOutput, bool) bool) error
	ListFlowsPagesWithContext(aws.Context, *mediaconnect.ListFlowsInput, func(*mediaconnect.ListFlowsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*mediaconnect.ListTagsForResourceInput) (*mediaconnect.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *mediaconnect.ListTagsForResourceInput, ...request.Option) (*mediaconnect.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*mediaconnect.ListTagsForResourceInput) (*request.Request, *mediaconnect.ListTagsForResourceOutput)

	RemoveFlowOutput(*mediaconnect.RemoveFlowOutputInput) (*mediaconnect.RemoveFlowOutputOutput, error)
	RemoveFlowOutputWithContext(aws.Context, *mediaconnect.RemoveFlowOutputInput, ...request.Option) (*mediaconnect.RemoveFlowOutputOutput, error)
	RemoveFlowOutputRequest(*mediaconnect.RemoveFlowOutputInput) (*request.Request, *mediaconnect.RemoveFlowOutputOutput)

	RevokeFlowEntitlement(*mediaconnect.RevokeFlowEntitlementInput) (*mediaconnect.RevokeFlowEntitlementOutput, error)
	RevokeFlowEntitlementWithContext(aws.Context, *mediaconnect.RevokeFlowEntitlementInput, ...request.Option) (*mediaconnect.RevokeFlowEntitlementOutput, error)
	RevokeFlowEntitlementRequest(*mediaconnect.RevokeFlowEntitlementInput) (*request.Request, *mediaconnect.RevokeFlowEntitlementOutput)

	StartFlow(*mediaconnect.StartFlowInput) (*mediaconnect.StartFlowOutput, error)
	StartFlowWithContext(aws.Context, *mediaconnect.StartFlowInput, ...request.Option) (*mediaconnect.StartFlowOutput, error)
	StartFlowRequest(*mediaconnect.StartFlowInput) (*request.Request, *mediaconnect.StartFlowOutput)

	StopFlow(*mediaconnect.StopFlowInput) (*mediaconnect.StopFlowOutput, error)
	StopFlowWithContext(aws.Context, *mediaconnect.StopFlowInput, ...request.Option) (*mediaconnect.StopFlowOutput, error)
	StopFlowRequest(*mediaconnect.StopFlowInput) (*request.Request, *mediaconnect.StopFlowOutput)

	TagResource(*mediaconnect.TagResourceInput) (*mediaconnect.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *mediaconnect.TagResourceInput, ...request.Option) (*mediaconnect.TagResourceOutput, error)
	TagResourceRequest(*mediaconnect.TagResourceInput) (*request.Request, *mediaconnect.TagResourceOutput)

	UntagResource(*mediaconnect.UntagResourceInput) (*mediaconnect.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *mediaconnect.UntagResourceInput, ...request.Option) (*mediaconnect.UntagResourceOutput, error)
	UntagResourceRequest(*mediaconnect.UntagResourceInput) (*request.Request, *mediaconnect.UntagResourceOutput)

	UpdateFlowEntitlement(*mediaconnect.UpdateFlowEntitlementInput) (*mediaconnect.UpdateFlowEntitlementOutput, error)
	UpdateFlowEntitlementWithContext(aws.Context, *mediaconnect.UpdateFlowEntitlementInput, ...request.Option) (*mediaconnect.UpdateFlowEntitlementOutput, error)
	UpdateFlowEntitlementRequest(*mediaconnect.UpdateFlowEntitlementInput) (*request.Request, *mediaconnect.UpdateFlowEntitlementOutput)

	UpdateFlowOutput(*mediaconnect.UpdateFlowOutputInput) (*mediaconnect.UpdateFlowOutputOutput, error)
	UpdateFlowOutputWithContext(aws.Context, *mediaconnect.UpdateFlowOutputInput, ...request.Option) (*mediaconnect.UpdateFlowOutputOutput, error)
	UpdateFlowOutputRequest(*mediaconnect.UpdateFlowOutputInput) (*request.Request, *mediaconnect.UpdateFlowOutputOutput)

	UpdateFlowSource(*mediaconnect.UpdateFlowSourceInput) (*mediaconnect.UpdateFlowSourceOutput, error)
	UpdateFlowSourceWithContext(aws.Context, *mediaconnect.UpdateFlowSourceInput, ...request.Option) (*mediaconnect.UpdateFlowSourceOutput, error)
	UpdateFlowSourceRequest(*mediaconnect.UpdateFlowSourceInput) (*request.Request, *mediaconnect.UpdateFlowSourceOutput)
}

var _ MediaConnectAPI = (*mediaconnect.MediaConnect)(nil)
