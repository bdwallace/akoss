// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package lexruntimeserviceiface provides an interface to enable mocking the Amazon Lex Runtime Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package lexruntimeserviceiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/lexruntimeservice"
)

// LexRuntimeServiceAPI provides an interface to enable mocking the
// lexruntimeservice.LexRuntimeService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Lex Runtime Service.
//    func myFunc(svc lexruntimeserviceiface.LexRuntimeServiceAPI) bool {
//        // Make svc.DeleteSession request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := lexruntimeservice.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockLexRuntimeServiceClient struct {
//        lexruntimeserviceiface.LexRuntimeServiceAPI
//    }
//    func (m *mockLexRuntimeServiceClient) DeleteSession(input *lexruntimeservice.DeleteSessionInput) (*lexruntimeservice.DeleteSessionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockLexRuntimeServiceClient{}
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
type LexRuntimeServiceAPI interface {
	DeleteSession(*lexruntimeservice.DeleteSessionInput) (*lexruntimeservice.DeleteSessionOutput, error)
	DeleteSessionWithContext(aws.Context, *lexruntimeservice.DeleteSessionInput, ...request.Option) (*lexruntimeservice.DeleteSessionOutput, error)
	DeleteSessionRequest(*lexruntimeservice.DeleteSessionInput) (*request.Request, *lexruntimeservice.DeleteSessionOutput)

	GetSession(*lexruntimeservice.GetSessionInput) (*lexruntimeservice.GetSessionOutput, error)
	GetSessionWithContext(aws.Context, *lexruntimeservice.GetSessionInput, ...request.Option) (*lexruntimeservice.GetSessionOutput, error)
	GetSessionRequest(*lexruntimeservice.GetSessionInput) (*request.Request, *lexruntimeservice.GetSessionOutput)

	PostContent(*lexruntimeservice.PostContentInput) (*lexruntimeservice.PostContentOutput, error)
	PostContentWithContext(aws.Context, *lexruntimeservice.PostContentInput, ...request.Option) (*lexruntimeservice.PostContentOutput, error)
	PostContentRequest(*lexruntimeservice.PostContentInput) (*request.Request, *lexruntimeservice.PostContentOutput)

	PostText(*lexruntimeservice.PostTextInput) (*lexruntimeservice.PostTextOutput, error)
	PostTextWithContext(aws.Context, *lexruntimeservice.PostTextInput, ...request.Option) (*lexruntimeservice.PostTextOutput, error)
	PostTextRequest(*lexruntimeservice.PostTextInput) (*request.Request, *lexruntimeservice.PostTextOutput)

	PutSession(*lexruntimeservice.PutSessionInput) (*lexruntimeservice.PutSessionOutput, error)
	PutSessionWithContext(aws.Context, *lexruntimeservice.PutSessionInput, ...request.Option) (*lexruntimeservice.PutSessionOutput, error)
	PutSessionRequest(*lexruntimeservice.PutSessionInput) (*request.Request, *lexruntimeservice.PutSessionOutput)
}

var _ LexRuntimeServiceAPI = (*lexruntimeservice.LexRuntimeService)(nil)
