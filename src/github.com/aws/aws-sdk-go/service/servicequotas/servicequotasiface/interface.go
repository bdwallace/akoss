// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package servicequotasiface provides an interface to enable mocking the Service Quotas service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package servicequotasiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/servicequotas"
)

// ServiceQuotasAPI provides an interface to enable mocking the
// servicequotas.ServiceQuotas service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Service Quotas.
//    func myFunc(svc servicequotasiface.ServiceQuotasAPI) bool {
//        // Make svc.AssociateServiceQuotaTemplate request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := servicequotas.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockServiceQuotasClient struct {
//        servicequotasiface.ServiceQuotasAPI
//    }
//    func (m *mockServiceQuotasClient) AssociateServiceQuotaTemplate(input *servicequotas.AssociateServiceQuotaTemplateInput) (*servicequotas.AssociateServiceQuotaTemplateOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockServiceQuotasClient{}
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
type ServiceQuotasAPI interface {
	AssociateServiceQuotaTemplate(*servicequotas.AssociateServiceQuotaTemplateInput) (*servicequotas.AssociateServiceQuotaTemplateOutput, error)
	AssociateServiceQuotaTemplateWithContext(aws.Context, *servicequotas.AssociateServiceQuotaTemplateInput, ...request.Option) (*servicequotas.AssociateServiceQuotaTemplateOutput, error)
	AssociateServiceQuotaTemplateRequest(*servicequotas.AssociateServiceQuotaTemplateInput) (*request.Request, *servicequotas.AssociateServiceQuotaTemplateOutput)

	DeleteServiceQuotaIncreaseRequestFromTemplate(*servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateInput) (*servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateOutput, error)
	DeleteServiceQuotaIncreaseRequestFromTemplateWithContext(aws.Context, *servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateInput, ...request.Option) (*servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateOutput, error)
	DeleteServiceQuotaIncreaseRequestFromTemplateRequest(*servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateInput) (*request.Request, *servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateOutput)

	DisassociateServiceQuotaTemplate(*servicequotas.DisassociateServiceQuotaTemplateInput) (*servicequotas.DisassociateServiceQuotaTemplateOutput, error)
	DisassociateServiceQuotaTemplateWithContext(aws.Context, *servicequotas.DisassociateServiceQuotaTemplateInput, ...request.Option) (*servicequotas.DisassociateServiceQuotaTemplateOutput, error)
	DisassociateServiceQuotaTemplateRequest(*servicequotas.DisassociateServiceQuotaTemplateInput) (*request.Request, *servicequotas.DisassociateServiceQuotaTemplateOutput)

	GetAWSDefaultServiceQuota(*servicequotas.GetAWSDefaultServiceQuotaInput) (*servicequotas.GetAWSDefaultServiceQuotaOutput, error)
	GetAWSDefaultServiceQuotaWithContext(aws.Context, *servicequotas.GetAWSDefaultServiceQuotaInput, ...request.Option) (*servicequotas.GetAWSDefaultServiceQuotaOutput, error)
	GetAWSDefaultServiceQuotaRequest(*servicequotas.GetAWSDefaultServiceQuotaInput) (*request.Request, *servicequotas.GetAWSDefaultServiceQuotaOutput)

	GetAssociationForServiceQuotaTemplate(*servicequotas.GetAssociationForServiceQuotaTemplateInput) (*servicequotas.GetAssociationForServiceQuotaTemplateOutput, error)
	GetAssociationForServiceQuotaTemplateWithContext(aws.Context, *servicequotas.GetAssociationForServiceQuotaTemplateInput, ...request.Option) (*servicequotas.GetAssociationForServiceQuotaTemplateOutput, error)
	GetAssociationForServiceQuotaTemplateRequest(*servicequotas.GetAssociationForServiceQuotaTemplateInput) (*request.Request, *servicequotas.GetAssociationForServiceQuotaTemplateOutput)

	GetRequestedServiceQuotaChange(*servicequotas.GetRequestedServiceQuotaChangeInput) (*servicequotas.GetRequestedServiceQuotaChangeOutput, error)
	GetRequestedServiceQuotaChangeWithContext(aws.Context, *servicequotas.GetRequestedServiceQuotaChangeInput, ...request.Option) (*servicequotas.GetRequestedServiceQuotaChangeOutput, error)
	GetRequestedServiceQuotaChangeRequest(*servicequotas.GetRequestedServiceQuotaChangeInput) (*request.Request, *servicequotas.GetRequestedServiceQuotaChangeOutput)

	GetServiceQuota(*servicequotas.GetServiceQuotaInput) (*servicequotas.GetServiceQuotaOutput, error)
	GetServiceQuotaWithContext(aws.Context, *servicequotas.GetServiceQuotaInput, ...request.Option) (*servicequotas.GetServiceQuotaOutput, error)
	GetServiceQuotaRequest(*servicequotas.GetServiceQuotaInput) (*request.Request, *servicequotas.GetServiceQuotaOutput)

	GetServiceQuotaIncreaseRequestFromTemplate(*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput) (*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput, error)
	GetServiceQuotaIncreaseRequestFromTemplateWithContext(aws.Context, *servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput, ...request.Option) (*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput, error)
	GetServiceQuotaIncreaseRequestFromTemplateRequest(*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput) (*request.Request, *servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput)

	ListAWSDefaultServiceQuotas(*servicequotas.ListAWSDefaultServiceQuotasInput) (*servicequotas.ListAWSDefaultServiceQuotasOutput, error)
	ListAWSDefaultServiceQuotasWithContext(aws.Context, *servicequotas.ListAWSDefaultServiceQuotasInput, ...request.Option) (*servicequotas.ListAWSDefaultServiceQuotasOutput, error)
	ListAWSDefaultServiceQuotasRequest(*servicequotas.ListAWSDefaultServiceQuotasInput) (*request.Request, *servicequotas.ListAWSDefaultServiceQuotasOutput)

	ListAWSDefaultServiceQuotasPages(*servicequotas.ListAWSDefaultServiceQuotasInput, func(*servicequotas.ListAWSDefaultServiceQuotasOutput, bool) bool) error
	ListAWSDefaultServiceQuotasPagesWithContext(aws.Context, *servicequotas.ListAWSDefaultServiceQuotasInput, func(*servicequotas.ListAWSDefaultServiceQuotasOutput, bool) bool, ...request.Option) error

	ListRequestedServiceQuotaChangeHistory(*servicequotas.ListRequestedServiceQuotaChangeHistoryInput) (*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, error)
	ListRequestedServiceQuotaChangeHistoryWithContext(aws.Context, *servicequotas.ListRequestedServiceQuotaChangeHistoryInput, ...request.Option) (*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, error)
	ListRequestedServiceQuotaChangeHistoryRequest(*servicequotas.ListRequestedServiceQuotaChangeHistoryInput) (*request.Request, *servicequotas.ListRequestedServiceQuotaChangeHistoryOutput)

	ListRequestedServiceQuotaChangeHistoryPages(*servicequotas.ListRequestedServiceQuotaChangeHistoryInput, func(*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, bool) bool) error
	ListRequestedServiceQuotaChangeHistoryPagesWithContext(aws.Context, *servicequotas.ListRequestedServiceQuotaChangeHistoryInput, func(*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, bool) bool, ...request.Option) error

	ListRequestedServiceQuotaChangeHistoryByQuota(*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput) (*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error)
	ListRequestedServiceQuotaChangeHistoryByQuotaWithContext(aws.Context, *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput, ...request.Option) (*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error)
	ListRequestedServiceQuotaChangeHistoryByQuotaRequest(*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput) (*request.Request, *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput)

	ListRequestedServiceQuotaChangeHistoryByQuotaPages(*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput, func(*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, bool) bool) error
	ListRequestedServiceQuotaChangeHistoryByQuotaPagesWithContext(aws.Context, *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput, func(*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, bool) bool, ...request.Option) error

	ListServiceQuotaIncreaseRequestsInTemplate(*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput) (*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, error)
	ListServiceQuotaIncreaseRequestsInTemplateWithContext(aws.Context, *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput, ...request.Option) (*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, error)
	ListServiceQuotaIncreaseRequestsInTemplateRequest(*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput) (*request.Request, *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput)

	ListServiceQuotaIncreaseRequestsInTemplatePages(*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput, func(*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, bool) bool) error
	ListServiceQuotaIncreaseRequestsInTemplatePagesWithContext(aws.Context, *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput, func(*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, bool) bool, ...request.Option) error

	ListServiceQuotas(*servicequotas.ListServiceQuotasInput) (*servicequotas.ListServiceQuotasOutput, error)
	ListServiceQuotasWithContext(aws.Context, *servicequotas.ListServiceQuotasInput, ...request.Option) (*servicequotas.ListServiceQuotasOutput, error)
	ListServiceQuotasRequest(*servicequotas.ListServiceQuotasInput) (*request.Request, *servicequotas.ListServiceQuotasOutput)

	ListServiceQuotasPages(*servicequotas.ListServiceQuotasInput, func(*servicequotas.ListServiceQuotasOutput, bool) bool) error
	ListServiceQuotasPagesWithContext(aws.Context, *servicequotas.ListServiceQuotasInput, func(*servicequotas.ListServiceQuotasOutput, bool) bool, ...request.Option) error

	ListServices(*servicequotas.ListServicesInput) (*servicequotas.ListServicesOutput, error)
	ListServicesWithContext(aws.Context, *servicequotas.ListServicesInput, ...request.Option) (*servicequotas.ListServicesOutput, error)
	ListServicesRequest(*servicequotas.ListServicesInput) (*request.Request, *servicequotas.ListServicesOutput)

	ListServicesPages(*servicequotas.ListServicesInput, func(*servicequotas.ListServicesOutput, bool) bool) error
	ListServicesPagesWithContext(aws.Context, *servicequotas.ListServicesInput, func(*servicequotas.ListServicesOutput, bool) bool, ...request.Option) error

	PutServiceQuotaIncreaseRequestIntoTemplate(*servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateInput) (*servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateOutput, error)
	PutServiceQuotaIncreaseRequestIntoTemplateWithContext(aws.Context, *servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateInput, ...request.Option) (*servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateOutput, error)
	PutServiceQuotaIncreaseRequestIntoTemplateRequest(*servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateInput) (*request.Request, *servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateOutput)

	RequestServiceQuotaIncrease(*servicequotas.RequestServiceQuotaIncreaseInput) (*servicequotas.RequestServiceQuotaIncreaseOutput, error)
	RequestServiceQuotaIncreaseWithContext(aws.Context, *servicequotas.RequestServiceQuotaIncreaseInput, ...request.Option) (*servicequotas.RequestServiceQuotaIncreaseOutput, error)
	RequestServiceQuotaIncreaseRequest(*servicequotas.RequestServiceQuotaIncreaseInput) (*request.Request, *servicequotas.RequestServiceQuotaIncreaseOutput)
}

var _ ServiceQuotasAPI = (*servicequotas.ServiceQuotas)(nil)
