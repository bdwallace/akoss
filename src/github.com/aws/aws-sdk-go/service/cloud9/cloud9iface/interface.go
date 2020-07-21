// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloud9iface provides an interface to enable mocking the AWS Cloud9 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloud9iface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloud9"
)

// Cloud9API provides an interface to enable mocking the
// cloud9.Cloud9 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Cloud9.
//    func myFunc(svc cloud9iface.Cloud9API) bool {
//        // Make svc.CreateEnvironmentEC2 request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cloud9.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCloud9Client struct {
//        cloud9iface.Cloud9API
//    }
//    func (m *mockCloud9Client) CreateEnvironmentEC2(input *cloud9.CreateEnvironmentEC2Input) (*cloud9.CreateEnvironmentEC2Output, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCloud9Client{}
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
type Cloud9API interface {
	CreateEnvironmentEC2(*cloud9.CreateEnvironmentEC2Input) (*cloud9.CreateEnvironmentEC2Output, error)
	CreateEnvironmentEC2WithContext(aws.Context, *cloud9.CreateEnvironmentEC2Input, ...request.Option) (*cloud9.CreateEnvironmentEC2Output, error)
	CreateEnvironmentEC2Request(*cloud9.CreateEnvironmentEC2Input) (*request.Request, *cloud9.CreateEnvironmentEC2Output)

	CreateEnvironmentMembership(*cloud9.CreateEnvironmentMembershipInput) (*cloud9.CreateEnvironmentMembershipOutput, error)
	CreateEnvironmentMembershipWithContext(aws.Context, *cloud9.CreateEnvironmentMembershipInput, ...request.Option) (*cloud9.CreateEnvironmentMembershipOutput, error)
	CreateEnvironmentMembershipRequest(*cloud9.CreateEnvironmentMembershipInput) (*request.Request, *cloud9.CreateEnvironmentMembershipOutput)

	DeleteEnvironment(*cloud9.DeleteEnvironmentInput) (*cloud9.DeleteEnvironmentOutput, error)
	DeleteEnvironmentWithContext(aws.Context, *cloud9.DeleteEnvironmentInput, ...request.Option) (*cloud9.DeleteEnvironmentOutput, error)
	DeleteEnvironmentRequest(*cloud9.DeleteEnvironmentInput) (*request.Request, *cloud9.DeleteEnvironmentOutput)

	DeleteEnvironmentMembership(*cloud9.DeleteEnvironmentMembershipInput) (*cloud9.DeleteEnvironmentMembershipOutput, error)
	DeleteEnvironmentMembershipWithContext(aws.Context, *cloud9.DeleteEnvironmentMembershipInput, ...request.Option) (*cloud9.DeleteEnvironmentMembershipOutput, error)
	DeleteEnvironmentMembershipRequest(*cloud9.DeleteEnvironmentMembershipInput) (*request.Request, *cloud9.DeleteEnvironmentMembershipOutput)

	DescribeEnvironmentMemberships(*cloud9.DescribeEnvironmentMembershipsInput) (*cloud9.DescribeEnvironmentMembershipsOutput, error)
	DescribeEnvironmentMembershipsWithContext(aws.Context, *cloud9.DescribeEnvironmentMembershipsInput, ...request.Option) (*cloud9.DescribeEnvironmentMembershipsOutput, error)
	DescribeEnvironmentMembershipsRequest(*cloud9.DescribeEnvironmentMembershipsInput) (*request.Request, *cloud9.DescribeEnvironmentMembershipsOutput)

	DescribeEnvironmentMembershipsPages(*cloud9.DescribeEnvironmentMembershipsInput, func(*cloud9.DescribeEnvironmentMembershipsOutput, bool) bool) error
	DescribeEnvironmentMembershipsPagesWithContext(aws.Context, *cloud9.DescribeEnvironmentMembershipsInput, func(*cloud9.DescribeEnvironmentMembershipsOutput, bool) bool, ...request.Option) error

	DescribeEnvironmentStatus(*cloud9.DescribeEnvironmentStatusInput) (*cloud9.DescribeEnvironmentStatusOutput, error)
	DescribeEnvironmentStatusWithContext(aws.Context, *cloud9.DescribeEnvironmentStatusInput, ...request.Option) (*cloud9.DescribeEnvironmentStatusOutput, error)
	DescribeEnvironmentStatusRequest(*cloud9.DescribeEnvironmentStatusInput) (*request.Request, *cloud9.DescribeEnvironmentStatusOutput)

	DescribeEnvironments(*cloud9.DescribeEnvironmentsInput) (*cloud9.DescribeEnvironmentsOutput, error)
	DescribeEnvironmentsWithContext(aws.Context, *cloud9.DescribeEnvironmentsInput, ...request.Option) (*cloud9.DescribeEnvironmentsOutput, error)
	DescribeEnvironmentsRequest(*cloud9.DescribeEnvironmentsInput) (*request.Request, *cloud9.DescribeEnvironmentsOutput)

	ListEnvironments(*cloud9.ListEnvironmentsInput) (*cloud9.ListEnvironmentsOutput, error)
	ListEnvironmentsWithContext(aws.Context, *cloud9.ListEnvironmentsInput, ...request.Option) (*cloud9.ListEnvironmentsOutput, error)
	ListEnvironmentsRequest(*cloud9.ListEnvironmentsInput) (*request.Request, *cloud9.ListEnvironmentsOutput)

	ListEnvironmentsPages(*cloud9.ListEnvironmentsInput, func(*cloud9.ListEnvironmentsOutput, bool) bool) error
	ListEnvironmentsPagesWithContext(aws.Context, *cloud9.ListEnvironmentsInput, func(*cloud9.ListEnvironmentsOutput, bool) bool, ...request.Option) error

	UpdateEnvironment(*cloud9.UpdateEnvironmentInput) (*cloud9.UpdateEnvironmentOutput, error)
	UpdateEnvironmentWithContext(aws.Context, *cloud9.UpdateEnvironmentInput, ...request.Option) (*cloud9.UpdateEnvironmentOutput, error)
	UpdateEnvironmentRequest(*cloud9.UpdateEnvironmentInput) (*request.Request, *cloud9.UpdateEnvironmentOutput)

	UpdateEnvironmentMembership(*cloud9.UpdateEnvironmentMembershipInput) (*cloud9.UpdateEnvironmentMembershipOutput, error)
	UpdateEnvironmentMembershipWithContext(aws.Context, *cloud9.UpdateEnvironmentMembershipInput, ...request.Option) (*cloud9.UpdateEnvironmentMembershipOutput, error)
	UpdateEnvironmentMembershipRequest(*cloud9.UpdateEnvironmentMembershipInput) (*request.Request, *cloud9.UpdateEnvironmentMembershipOutput)
}

var _ Cloud9API = (*cloud9.Cloud9)(nil)
