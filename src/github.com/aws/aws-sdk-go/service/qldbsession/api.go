// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package qldbsession

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opSendCommand = "SendCommand"

// SendCommandRequest generates a "aws/request.Request" representing the
// client's request for the SendCommand operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See SendCommand for more information on using the SendCommand
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the SendCommandRequest method.
//    req, resp := client.SendCommandRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/qldb-session-2019-07-11/SendCommand
func (c *QLDBSession) SendCommandRequest(input *SendCommandInput) (req *request.Request, output *SendCommandOutput) {
	op := &request.Operation{
		Name:       opSendCommand,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SendCommandInput{}
	}

	output = &SendCommandOutput{}
	req = c.newRequest(op, input, output)
	return
}

// SendCommand API operation for Amazon QLDB Session.
//
// Sends a command to an Amazon QLDB ledger.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon QLDB Session's
// API operation SendCommand for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeBadRequestException "BadRequestException"
//   Returned if the request is malformed or contains an error such as an invalid
//   parameter value or a missing required parameter.
//
//   * ErrCodeInvalidSessionException "InvalidSessionException"
//   Returned if the session doesn't exist anymore because it timed-out or expired.
//
//   * ErrCodeOccConflictException "OccConflictException"
//   Returned when a transaction cannot be written to the journal due to a failure
//   in the verification phase of Optimistic Concurrency Control.
//
//   * ErrCodeRateExceededException "RateExceededException"
//   Returned when the rate of requests exceeds the allowed throughput.
//
//   * ErrCodeLimitExceededException "LimitExceededException"
//   Returned if a resource limit such as number of active sessions is exceeded.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/qldb-session-2019-07-11/SendCommand
func (c *QLDBSession) SendCommand(input *SendCommandInput) (*SendCommandOutput, error) {
	req, out := c.SendCommandRequest(input)
	return out, req.Send()
}

// SendCommandWithContext is the same as SendCommand with the addition of
// the ability to pass a context and additional request options.
//
// See SendCommand for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *QLDBSession) SendCommandWithContext(ctx aws.Context, input *SendCommandInput, opts ...request.Option) (*SendCommandOutput, error) {
	req, out := c.SendCommandRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// Contains the details of the transaction to abort.
type AbortTransactionRequest struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AbortTransactionRequest) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AbortTransactionRequest) GoString() string {
	return s.String()
}

// Contains the details of the aborted transaction.
type AbortTransactionResult struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AbortTransactionResult) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AbortTransactionResult) GoString() string {
	return s.String()
}

// Contains the details of the transaction to commit.
type CommitTransactionRequest struct {
	_ struct{} `type:"structure"`

	// Specifies the commit digest for the transaction to commit. For every active
	// transaction, the commit digest must be passed. QLDB validates CommitDigest
	// and rejects the commit with an error if the digest computed on the client
	// does not match the digest computed by QLDB.
	//
	// CommitDigest is automatically base64 encoded/decoded by the SDK.
	//
	// CommitDigest is a required field
	CommitDigest []byte `type:"blob" required:"true"`

	// Specifies the transaction id of the transaction to commit.
	//
	// TransactionId is a required field
	TransactionId *string `min:"22" type:"string" required:"true"`
}

// String returns the string representation
func (s CommitTransactionRequest) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CommitTransactionRequest) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CommitTransactionRequest) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CommitTransactionRequest"}
	if s.CommitDigest == nil {
		invalidParams.Add(request.NewErrParamRequired("CommitDigest"))
	}
	if s.TransactionId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransactionId"))
	}
	if s.TransactionId != nil && len(*s.TransactionId) < 22 {
		invalidParams.Add(request.NewErrParamMinLen("TransactionId", 22))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCommitDigest sets the CommitDigest field's value.
func (s *CommitTransactionRequest) SetCommitDigest(v []byte) *CommitTransactionRequest {
	s.CommitDigest = v
	return s
}

// SetTransactionId sets the TransactionId field's value.
func (s *CommitTransactionRequest) SetTransactionId(v string) *CommitTransactionRequest {
	s.TransactionId = &v
	return s
}

// Contains the details of the committed transaction.
type CommitTransactionResult struct {
	_ struct{} `type:"structure"`

	// The commit digest of the committed transaction.
	//
	// CommitDigest is automatically base64 encoded/decoded by the SDK.
	CommitDigest []byte `type:"blob"`

	// The transaction id of the committed transaction.
	TransactionId *string `min:"22" type:"string"`
}

// String returns the string representation
func (s CommitTransactionResult) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CommitTransactionResult) GoString() string {
	return s.String()
}

// SetCommitDigest sets the CommitDigest field's value.
func (s *CommitTransactionResult) SetCommitDigest(v []byte) *CommitTransactionResult {
	s.CommitDigest = v
	return s
}

// SetTransactionId sets the TransactionId field's value.
func (s *CommitTransactionResult) SetTransactionId(v string) *CommitTransactionResult {
	s.TransactionId = &v
	return s
}

// Specifies a request to end the session.
type EndSessionRequest struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EndSessionRequest) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s EndSessionRequest) GoString() string {
	return s.String()
}

// Contains the details of the ended session.
type EndSessionResult struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EndSessionResult) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s EndSessionResult) GoString() string {
	return s.String()
}

// Specifies a request to execute a statement.
type ExecuteStatementRequest struct {
	_ struct{} `type:"structure"`

	// Specifies the parameters for the parameterized statement in the request.
	Parameters []*ValueHolder `type:"list"`

	// Specifies the statement of the request.
	//
	// Statement is a required field
	Statement *string `min:"1" type:"string" required:"true"`

	// Specifies the transaction id of the request.
	//
	// TransactionId is a required field
	TransactionId *string `min:"22" type:"string" required:"true"`
}

// String returns the string representation
func (s ExecuteStatementRequest) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ExecuteStatementRequest) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ExecuteStatementRequest) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ExecuteStatementRequest"}
	if s.Statement == nil {
		invalidParams.Add(request.NewErrParamRequired("Statement"))
	}
	if s.Statement != nil && len(*s.Statement) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Statement", 1))
	}
	if s.TransactionId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransactionId"))
	}
	if s.TransactionId != nil && len(*s.TransactionId) < 22 {
		invalidParams.Add(request.NewErrParamMinLen("TransactionId", 22))
	}
	if s.Parameters != nil {
		for i, v := range s.Parameters {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Parameters", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetParameters sets the Parameters field's value.
func (s *ExecuteStatementRequest) SetParameters(v []*ValueHolder) *ExecuteStatementRequest {
	s.Parameters = v
	return s
}

// SetStatement sets the Statement field's value.
func (s *ExecuteStatementRequest) SetStatement(v string) *ExecuteStatementRequest {
	s.Statement = &v
	return s
}

// SetTransactionId sets the TransactionId field's value.
func (s *ExecuteStatementRequest) SetTransactionId(v string) *ExecuteStatementRequest {
	s.TransactionId = &v
	return s
}

// Contains the details of the executed statement.
type ExecuteStatementResult struct {
	_ struct{} `type:"structure"`

	// Contains the details of the first fetched page.
	FirstPage *Page `type:"structure"`
}

// String returns the string representation
func (s ExecuteStatementResult) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ExecuteStatementResult) GoString() string {
	return s.String()
}

// SetFirstPage sets the FirstPage field's value.
func (s *ExecuteStatementResult) SetFirstPage(v *Page) *ExecuteStatementResult {
	s.FirstPage = v
	return s
}

// Specifies the details of the page to be fetched.
type FetchPageRequest struct {
	_ struct{} `type:"structure"`

	// Specifies the next page token of the page to be fetched.
	//
	// NextPageToken is a required field
	NextPageToken *string `min:"4" type:"string" required:"true"`

	// Specifies the transaction id of the page to be fetched.
	//
	// TransactionId is a required field
	TransactionId *string `min:"22" type:"string" required:"true"`
}

// String returns the string representation
func (s FetchPageRequest) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s FetchPageRequest) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *FetchPageRequest) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "FetchPageRequest"}
	if s.NextPageToken == nil {
		invalidParams.Add(request.NewErrParamRequired("NextPageToken"))
	}
	if s.NextPageToken != nil && len(*s.NextPageToken) < 4 {
		invalidParams.Add(request.NewErrParamMinLen("NextPageToken", 4))
	}
	if s.TransactionId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransactionId"))
	}
	if s.TransactionId != nil && len(*s.TransactionId) < 22 {
		invalidParams.Add(request.NewErrParamMinLen("TransactionId", 22))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetNextPageToken sets the NextPageToken field's value.
func (s *FetchPageRequest) SetNextPageToken(v string) *FetchPageRequest {
	s.NextPageToken = &v
	return s
}

// SetTransactionId sets the TransactionId field's value.
func (s *FetchPageRequest) SetTransactionId(v string) *FetchPageRequest {
	s.TransactionId = &v
	return s
}

// Contains the page that was fetched.
type FetchPageResult struct {
	_ struct{} `type:"structure"`

	// Contains details of the fetched page.
	Page *Page `type:"structure"`
}

// String returns the string representation
func (s FetchPageResult) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s FetchPageResult) GoString() string {
	return s.String()
}

// SetPage sets the Page field's value.
func (s *FetchPageResult) SetPage(v *Page) *FetchPageResult {
	s.Page = v
	return s
}

// Contains details of the fetched page.
type Page struct {
	_ struct{} `type:"structure"`

	// The token of the next page.
	NextPageToken *string `min:"4" type:"string"`

	// A structure that contains values in multiple encoding formats.
	Values []*ValueHolder `type:"list"`
}

// String returns the string representation
func (s Page) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Page) GoString() string {
	return s.String()
}

// SetNextPageToken sets the NextPageToken field's value.
func (s *Page) SetNextPageToken(v string) *Page {
	s.NextPageToken = &v
	return s
}

// SetValues sets the Values field's value.
func (s *Page) SetValues(v []*ValueHolder) *Page {
	s.Values = v
	return s
}

type SendCommandInput struct {
	_ struct{} `type:"structure"`

	// Command to abort the current transaction.
	AbortTransaction *AbortTransactionRequest `type:"structure"`

	// Command to commit the specified transaction.
	CommitTransaction *CommitTransactionRequest `type:"structure"`

	// Command to end the current session.
	EndSession *EndSessionRequest `type:"structure"`

	// Command to execute a statement in the specified transaction.
	ExecuteStatement *ExecuteStatementRequest `type:"structure"`

	// Command to fetch a page.
	FetchPage *FetchPageRequest `type:"structure"`

	// Specifies the session token for the current command. A session token is constant
	// throughout the life of the session.
	//
	// To obtain a session token, run the StartSession command. This SessionToken
	// is required for every subsequent command that is issued during the current
	// session.
	SessionToken *string `min:"4" type:"string"`

	// Command to start a new session. A session token is obtained as part of the
	// response.
	StartSession *StartSessionRequest `type:"structure"`

	// Command to start a new transaction.
	StartTransaction *StartTransactionRequest `type:"structure"`
}

// String returns the string representation
func (s SendCommandInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SendCommandInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendCommandInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SendCommandInput"}
	if s.SessionToken != nil && len(*s.SessionToken) < 4 {
		invalidParams.Add(request.NewErrParamMinLen("SessionToken", 4))
	}
	if s.CommitTransaction != nil {
		if err := s.CommitTransaction.Validate(); err != nil {
			invalidParams.AddNested("CommitTransaction", err.(request.ErrInvalidParams))
		}
	}
	if s.ExecuteStatement != nil {
		if err := s.ExecuteStatement.Validate(); err != nil {
			invalidParams.AddNested("ExecuteStatement", err.(request.ErrInvalidParams))
		}
	}
	if s.FetchPage != nil {
		if err := s.FetchPage.Validate(); err != nil {
			invalidParams.AddNested("FetchPage", err.(request.ErrInvalidParams))
		}
	}
	if s.StartSession != nil {
		if err := s.StartSession.Validate(); err != nil {
			invalidParams.AddNested("StartSession", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAbortTransaction sets the AbortTransaction field's value.
func (s *SendCommandInput) SetAbortTransaction(v *AbortTransactionRequest) *SendCommandInput {
	s.AbortTransaction = v
	return s
}

// SetCommitTransaction sets the CommitTransaction field's value.
func (s *SendCommandInput) SetCommitTransaction(v *CommitTransactionRequest) *SendCommandInput {
	s.CommitTransaction = v
	return s
}

// SetEndSession sets the EndSession field's value.
func (s *SendCommandInput) SetEndSession(v *EndSessionRequest) *SendCommandInput {
	s.EndSession = v
	return s
}

// SetExecuteStatement sets the ExecuteStatement field's value.
func (s *SendCommandInput) SetExecuteStatement(v *ExecuteStatementRequest) *SendCommandInput {
	s.ExecuteStatement = v
	return s
}

// SetFetchPage sets the FetchPage field's value.
func (s *SendCommandInput) SetFetchPage(v *FetchPageRequest) *SendCommandInput {
	s.FetchPage = v
	return s
}

// SetSessionToken sets the SessionToken field's value.
func (s *SendCommandInput) SetSessionToken(v string) *SendCommandInput {
	s.SessionToken = &v
	return s
}

// SetStartSession sets the StartSession field's value.
func (s *SendCommandInput) SetStartSession(v *StartSessionRequest) *SendCommandInput {
	s.StartSession = v
	return s
}

// SetStartTransaction sets the StartTransaction field's value.
func (s *SendCommandInput) SetStartTransaction(v *StartTransactionRequest) *SendCommandInput {
	s.StartTransaction = v
	return s
}

type SendCommandOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of the aborted transaction.
	AbortTransaction *AbortTransactionResult `type:"structure"`

	// Contains the details of the committed transaction.
	CommitTransaction *CommitTransactionResult `type:"structure"`

	// Contains the details of the ended session.
	EndSession *EndSessionResult `type:"structure"`

	// Contains the details of the executed statement.
	ExecuteStatement *ExecuteStatementResult `type:"structure"`

	// Contains the details of the fetched page.
	FetchPage *FetchPageResult `type:"structure"`

	// Contains the details of the started session that includes a session token.
	// This SessionToken is required for every subsequent command that is issued
	// during the current session.
	StartSession *StartSessionResult `type:"structure"`

	// Contains the details of the started transaction.
	StartTransaction *StartTransactionResult `type:"structure"`
}

// String returns the string representation
func (s SendCommandOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SendCommandOutput) GoString() string {
	return s.String()
}

// SetAbortTransaction sets the AbortTransaction field's value.
func (s *SendCommandOutput) SetAbortTransaction(v *AbortTransactionResult) *SendCommandOutput {
	s.AbortTransaction = v
	return s
}

// SetCommitTransaction sets the CommitTransaction field's value.
func (s *SendCommandOutput) SetCommitTransaction(v *CommitTransactionResult) *SendCommandOutput {
	s.CommitTransaction = v
	return s
}

// SetEndSession sets the EndSession field's value.
func (s *SendCommandOutput) SetEndSession(v *EndSessionResult) *SendCommandOutput {
	s.EndSession = v
	return s
}

// SetExecuteStatement sets the ExecuteStatement field's value.
func (s *SendCommandOutput) SetExecuteStatement(v *ExecuteStatementResult) *SendCommandOutput {
	s.ExecuteStatement = v
	return s
}

// SetFetchPage sets the FetchPage field's value.
func (s *SendCommandOutput) SetFetchPage(v *FetchPageResult) *SendCommandOutput {
	s.FetchPage = v
	return s
}

// SetStartSession sets the StartSession field's value.
func (s *SendCommandOutput) SetStartSession(v *StartSessionResult) *SendCommandOutput {
	s.StartSession = v
	return s
}

// SetStartTransaction sets the StartTransaction field's value.
func (s *SendCommandOutput) SetStartTransaction(v *StartTransactionResult) *SendCommandOutput {
	s.StartTransaction = v
	return s
}

// Specifies a request to start a a new session.
type StartSessionRequest struct {
	_ struct{} `type:"structure"`

	// The name of the ledger to start a new session against.
	//
	// LedgerName is a required field
	LedgerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartSessionRequest) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartSessionRequest) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartSessionRequest) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "StartSessionRequest"}
	if s.LedgerName == nil {
		invalidParams.Add(request.NewErrParamRequired("LedgerName"))
	}
	if s.LedgerName != nil && len(*s.LedgerName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("LedgerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetLedgerName sets the LedgerName field's value.
func (s *StartSessionRequest) SetLedgerName(v string) *StartSessionRequest {
	s.LedgerName = &v
	return s
}

// Contains the details of the started session.
type StartSessionResult struct {
	_ struct{} `type:"structure"`

	// Session token of the started session. This SessionToken is required for every
	// subsequent command that is issued during the current session.
	SessionToken *string `min:"4" type:"string"`
}

// String returns the string representation
func (s StartSessionResult) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartSessionResult) GoString() string {
	return s.String()
}

// SetSessionToken sets the SessionToken field's value.
func (s *StartSessionResult) SetSessionToken(v string) *StartSessionResult {
	s.SessionToken = &v
	return s
}

// Specifies a request to start a transaction.
type StartTransactionRequest struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartTransactionRequest) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartTransactionRequest) GoString() string {
	return s.String()
}

// Contains the details of the started transaction.
type StartTransactionResult struct {
	_ struct{} `type:"structure"`

	// The transaction id of the started transaction.
	TransactionId *string `min:"22" type:"string"`
}

// String returns the string representation
func (s StartTransactionResult) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartTransactionResult) GoString() string {
	return s.String()
}

// SetTransactionId sets the TransactionId field's value.
func (s *StartTransactionResult) SetTransactionId(v string) *StartTransactionResult {
	s.TransactionId = &v
	return s
}

// A structure that can contains values in multiple encoding formats.
type ValueHolder struct {
	_ struct{} `type:"structure"`

	// An Amazon Ion binary value contained in a ValueHolder structure.
	//
	// IonBinary is automatically base64 encoded/decoded by the SDK.
	IonBinary []byte `min:"1" type:"blob"`

	// An Amazon Ion plaintext value contained in a ValueHolder structure.
	IonText *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ValueHolder) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ValueHolder) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ValueHolder) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ValueHolder"}
	if s.IonBinary != nil && len(s.IonBinary) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("IonBinary", 1))
	}
	if s.IonText != nil && len(*s.IonText) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("IonText", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetIonBinary sets the IonBinary field's value.
func (s *ValueHolder) SetIonBinary(v []byte) *ValueHolder {
	s.IonBinary = v
	return s
}

// SetIonText sets the IonText field's value.
func (s *ValueHolder) SetIonText(v string) *ValueHolder {
	s.IonText = &v
	return s
}
