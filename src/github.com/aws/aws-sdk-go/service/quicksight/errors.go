// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You don't have access to this. The provided credentials couldn't be validated.
	// You might not be authorized to carry out the request. Ensure that your account
	// is authorized to use the Amazon QuickSight service, that your policies have
	// the correct permissions, and that you are using the correct access keys.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeDomainNotWhitelistedException for service response error code
	// "DomainNotWhitelistedException".
	//
	// The domain specified is not on the allowlist. All domains for embedded dashboards
	// must be added to the approved list by an Amazon QuickSight admin.
	ErrCodeDomainNotWhitelistedException = "DomainNotWhitelistedException"

	// ErrCodeIdentityTypeNotSupportedException for service response error code
	// "IdentityTypeNotSupportedException".
	//
	// The identity type specified is not supported. Supported identity types include
	// IAM and QUICKSIGHT.
	ErrCodeIdentityTypeNotSupportedException = "IdentityTypeNotSupportedException"

	// ErrCodeInternalFailureException for service response error code
	// "InternalFailureException".
	//
	// An internal failure occurred.
	ErrCodeInternalFailureException = "InternalFailureException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The NextToken value isn't valid.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidParameterValueException for service response error code
	// "InvalidParameterValueException".
	//
	// One or more parameters don't have a valid value.
	ErrCodeInvalidParameterValueException = "InvalidParameterValueException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// A limit is exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodePreconditionNotMetException for service response error code
	// "PreconditionNotMetException".
	//
	// One or more preconditions aren't met.
	ErrCodePreconditionNotMetException = "PreconditionNotMetException"

	// ErrCodeResourceExistsException for service response error code
	// "ResourceExistsException".
	//
	// The resource specified doesn't exist.
	ErrCodeResourceExistsException = "ResourceExistsException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// One or more resources can't be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeResourceUnavailableException for service response error code
	// "ResourceUnavailableException".
	//
	// This resource is currently unavailable.
	ErrCodeResourceUnavailableException = "ResourceUnavailableException"

	// ErrCodeSessionLifetimeInMinutesInvalidException for service response error code
	// "SessionLifetimeInMinutesInvalidException".
	//
	// The number of minutes specified for the lifetime of a session is not valid.
	// The session lifetime must be from 15 to 600 minutes.
	ErrCodeSessionLifetimeInMinutesInvalidException = "SessionLifetimeInMinutesInvalidException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// Access is throttled.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeUnsupportedUserEditionException for service response error code
	// "UnsupportedUserEditionException".
	//
	// This error indicates that you are calling an operation on an Amazon QuickSight
	// subscription where the edition doesn't include support for that operation.
	// Amazon QuickSight currently has Standard Edition and Enterprise Edition.
	// Not every operation and capability is available in every edition.
	ErrCodeUnsupportedUserEditionException = "UnsupportedUserEditionException"

	// ErrCodeUserNotFoundException for service response error code
	// "QuickSightUserNotFoundException".
	//
	// The user is not found. This error can happen in any operation that requires
	// finding a user based on a provided user name, such as DeleteUser, DescribeUser,
	// and so on.
	ErrCodeUserNotFoundException = "QuickSightUserNotFoundException"
)
