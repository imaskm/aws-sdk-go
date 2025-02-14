// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAliasExistsException for service response error code
	// "AliasExistsException".
	//
	// This exception is thrown when a user tries to confirm the account with an
	// email address or phone number that has already been supplied as an alias
	// for a different user profile. This exception indicates that an account with
	// this email address or phone already exists in a user pool that you've configured
	// to use email address or phone number as a sign-in alias.
	ErrCodeAliasExistsException = "AliasExistsException"

	// ErrCodeCodeDeliveryFailureException for service response error code
	// "CodeDeliveryFailureException".
	//
	// This exception is thrown when a verification code fails to deliver successfully.
	ErrCodeCodeDeliveryFailureException = "CodeDeliveryFailureException"

	// ErrCodeCodeMismatchException for service response error code
	// "CodeMismatchException".
	//
	// This exception is thrown if the provided code doesn't match what the server
	// was expecting.
	ErrCodeCodeMismatchException = "CodeMismatchException"

	// ErrCodeConcurrentModificationException for service response error code
	// "ConcurrentModificationException".
	//
	// This exception is thrown if two or more modifications are happening concurrently.
	ErrCodeConcurrentModificationException = "ConcurrentModificationException"

	// ErrCodeDuplicateProviderException for service response error code
	// "DuplicateProviderException".
	//
	// This exception is thrown when the provider is already supported by the user
	// pool.
	ErrCodeDuplicateProviderException = "DuplicateProviderException"

	// ErrCodeEnableSoftwareTokenMFAException for service response error code
	// "EnableSoftwareTokenMFAException".
	//
	// This exception is thrown when there is a code mismatch and the service fails
	// to configure the software token TOTP multi-factor authentication (MFA).
	ErrCodeEnableSoftwareTokenMFAException = "EnableSoftwareTokenMFAException"

	// ErrCodeExpiredCodeException for service response error code
	// "ExpiredCodeException".
	//
	// This exception is thrown if a code has expired.
	ErrCodeExpiredCodeException = "ExpiredCodeException"

	// ErrCodeGroupExistsException for service response error code
	// "GroupExistsException".
	//
	// This exception is thrown when Amazon Cognito encounters a group that already
	// exists in the user pool.
	ErrCodeGroupExistsException = "GroupExistsException"

	// ErrCodeInternalErrorException for service response error code
	// "InternalErrorException".
	//
	// This exception is thrown when Amazon Cognito encounters an internal error.
	ErrCodeInternalErrorException = "InternalErrorException"

	// ErrCodeInvalidEmailRoleAccessPolicyException for service response error code
	// "InvalidEmailRoleAccessPolicyException".
	//
	// This exception is thrown when Amazon Cognito isn't allowed to use your email
	// identity. HTTP status code: 400.
	ErrCodeInvalidEmailRoleAccessPolicyException = "InvalidEmailRoleAccessPolicyException"

	// ErrCodeInvalidLambdaResponseException for service response error code
	// "InvalidLambdaResponseException".
	//
	// This exception is thrown when Amazon Cognito encounters an invalid Lambda
	// response.
	ErrCodeInvalidLambdaResponseException = "InvalidLambdaResponseException"

	// ErrCodeInvalidOAuthFlowException for service response error code
	// "InvalidOAuthFlowException".
	//
	// This exception is thrown when the specified OAuth flow is not valid.
	ErrCodeInvalidOAuthFlowException = "InvalidOAuthFlowException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// This exception is thrown when the Amazon Cognito service encounters an invalid
	// parameter.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidPasswordException for service response error code
	// "InvalidPasswordException".
	//
	// This exception is thrown when Amazon Cognito encounters an invalid password.
	ErrCodeInvalidPasswordException = "InvalidPasswordException"

	// ErrCodeInvalidSmsRoleAccessPolicyException for service response error code
	// "InvalidSmsRoleAccessPolicyException".
	//
	// This exception is returned when the role provided for SMS configuration doesn't
	// have permission to publish using Amazon SNS.
	ErrCodeInvalidSmsRoleAccessPolicyException = "InvalidSmsRoleAccessPolicyException"

	// ErrCodeInvalidSmsRoleTrustRelationshipException for service response error code
	// "InvalidSmsRoleTrustRelationshipException".
	//
	// This exception is thrown when the trust relationship is not valid for the
	// role provided for SMS configuration. This can happen if you don't trust cognito-idp.amazonaws.com
	// or the external ID provided in the role does not match what is provided in
	// the SMS configuration for the user pool.
	ErrCodeInvalidSmsRoleTrustRelationshipException = "InvalidSmsRoleTrustRelationshipException"

	// ErrCodeInvalidUserPoolConfigurationException for service response error code
	// "InvalidUserPoolConfigurationException".
	//
	// This exception is thrown when the user pool configuration is not valid.
	ErrCodeInvalidUserPoolConfigurationException = "InvalidUserPoolConfigurationException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// This exception is thrown when a user exceeds the limit for a requested Amazon
	// Web Services resource.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeMFAMethodNotFoundException for service response error code
	// "MFAMethodNotFoundException".
	//
	// This exception is thrown when Amazon Cognito can't find a multi-factor authentication
	// (MFA) method.
	ErrCodeMFAMethodNotFoundException = "MFAMethodNotFoundException"

	// ErrCodeNotAuthorizedException for service response error code
	// "NotAuthorizedException".
	//
	// This exception is thrown when a user isn't authorized.
	ErrCodeNotAuthorizedException = "NotAuthorizedException"

	// ErrCodePasswordResetRequiredException for service response error code
	// "PasswordResetRequiredException".
	//
	// This exception is thrown when a password reset is required.
	ErrCodePasswordResetRequiredException = "PasswordResetRequiredException"

	// ErrCodePreconditionNotMetException for service response error code
	// "PreconditionNotMetException".
	//
	// This exception is thrown when a precondition is not met.
	ErrCodePreconditionNotMetException = "PreconditionNotMetException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// This exception is thrown when the Amazon Cognito service can't find the requested
	// resource.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeScopeDoesNotExistException for service response error code
	// "ScopeDoesNotExistException".
	//
	// This exception is thrown when the specified scope doesn't exist.
	ErrCodeScopeDoesNotExistException = "ScopeDoesNotExistException"

	// ErrCodeSoftwareTokenMFANotFoundException for service response error code
	// "SoftwareTokenMFANotFoundException".
	//
	// This exception is thrown when the software token time-based one-time password
	// (TOTP) multi-factor authentication (MFA) isn't activated for the user pool.
	ErrCodeSoftwareTokenMFANotFoundException = "SoftwareTokenMFANotFoundException"

	// ErrCodeTooManyFailedAttemptsException for service response error code
	// "TooManyFailedAttemptsException".
	//
	// This exception is thrown when the user has made too many failed attempts
	// for a given action, such as sign-in.
	ErrCodeTooManyFailedAttemptsException = "TooManyFailedAttemptsException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// This exception is thrown when the user has made too many requests for a given
	// operation.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"

	// ErrCodeUnauthorizedException for service response error code
	// "UnauthorizedException".
	//
	// Exception that is thrown when the request isn't authorized. This can happen
	// due to an invalid access token in the request.
	ErrCodeUnauthorizedException = "UnauthorizedException"

	// ErrCodeUnexpectedLambdaException for service response error code
	// "UnexpectedLambdaException".
	//
	// This exception is thrown when Amazon Cognito encounters an unexpected exception
	// with Lambda.
	ErrCodeUnexpectedLambdaException = "UnexpectedLambdaException"

	// ErrCodeUnsupportedIdentityProviderException for service response error code
	// "UnsupportedIdentityProviderException".
	//
	// This exception is thrown when the specified identifier isn't supported.
	ErrCodeUnsupportedIdentityProviderException = "UnsupportedIdentityProviderException"

	// ErrCodeUnsupportedOperationException for service response error code
	// "UnsupportedOperationException".
	//
	// Exception that is thrown when you attempt to perform an operation that isn't
	// enabled for the user pool client.
	ErrCodeUnsupportedOperationException = "UnsupportedOperationException"

	// ErrCodeUnsupportedTokenTypeException for service response error code
	// "UnsupportedTokenTypeException".
	//
	// Exception that is thrown when an unsupported token is passed to an operation.
	ErrCodeUnsupportedTokenTypeException = "UnsupportedTokenTypeException"

	// ErrCodeUnsupportedUserStateException for service response error code
	// "UnsupportedUserStateException".
	//
	// The request failed because the user is in an unsupported state.
	ErrCodeUnsupportedUserStateException = "UnsupportedUserStateException"

	// ErrCodeUserImportInProgressException for service response error code
	// "UserImportInProgressException".
	//
	// This exception is thrown when you're trying to modify a user pool while a
	// user import job is in progress for that pool.
	ErrCodeUserImportInProgressException = "UserImportInProgressException"

	// ErrCodeUserLambdaValidationException for service response error code
	// "UserLambdaValidationException".
	//
	// This exception is thrown when the Amazon Cognito service encounters a user
	// validation exception with the Lambda service.
	ErrCodeUserLambdaValidationException = "UserLambdaValidationException"

	// ErrCodeUserNotConfirmedException for service response error code
	// "UserNotConfirmedException".
	//
	// This exception is thrown when a user isn't confirmed successfully.
	ErrCodeUserNotConfirmedException = "UserNotConfirmedException"

	// ErrCodeUserNotFoundException for service response error code
	// "UserNotFoundException".
	//
	// This exception is thrown when a user isn't found.
	ErrCodeUserNotFoundException = "UserNotFoundException"

	// ErrCodeUserPoolAddOnNotEnabledException for service response error code
	// "UserPoolAddOnNotEnabledException".
	//
	// This exception is thrown when user pool add-ons aren't enabled.
	ErrCodeUserPoolAddOnNotEnabledException = "UserPoolAddOnNotEnabledException"

	// ErrCodeUserPoolTaggingException for service response error code
	// "UserPoolTaggingException".
	//
	// This exception is thrown when a user pool tag can't be set or updated.
	ErrCodeUserPoolTaggingException = "UserPoolTaggingException"

	// ErrCodeUsernameExistsException for service response error code
	// "UsernameExistsException".
	//
	// This exception is thrown when Amazon Cognito encounters a user name that
	// already exists in the user pool.
	ErrCodeUsernameExistsException = "UsernameExistsException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AliasExistsException":                     newErrorAliasExistsException,
	"CodeDeliveryFailureException":             newErrorCodeDeliveryFailureException,
	"CodeMismatchException":                    newErrorCodeMismatchException,
	"ConcurrentModificationException":          newErrorConcurrentModificationException,
	"DuplicateProviderException":               newErrorDuplicateProviderException,
	"EnableSoftwareTokenMFAException":          newErrorEnableSoftwareTokenMFAException,
	"ExpiredCodeException":                     newErrorExpiredCodeException,
	"GroupExistsException":                     newErrorGroupExistsException,
	"InternalErrorException":                   newErrorInternalErrorException,
	"InvalidEmailRoleAccessPolicyException":    newErrorInvalidEmailRoleAccessPolicyException,
	"InvalidLambdaResponseException":           newErrorInvalidLambdaResponseException,
	"InvalidOAuthFlowException":                newErrorInvalidOAuthFlowException,
	"InvalidParameterException":                newErrorInvalidParameterException,
	"InvalidPasswordException":                 newErrorInvalidPasswordException,
	"InvalidSmsRoleAccessPolicyException":      newErrorInvalidSmsRoleAccessPolicyException,
	"InvalidSmsRoleTrustRelationshipException": newErrorInvalidSmsRoleTrustRelationshipException,
	"InvalidUserPoolConfigurationException":    newErrorInvalidUserPoolConfigurationException,
	"LimitExceededException":                   newErrorLimitExceededException,
	"MFAMethodNotFoundException":               newErrorMFAMethodNotFoundException,
	"NotAuthorizedException":                   newErrorNotAuthorizedException,
	"PasswordResetRequiredException":           newErrorPasswordResetRequiredException,
	"PreconditionNotMetException":              newErrorPreconditionNotMetException,
	"ResourceNotFoundException":                newErrorResourceNotFoundException,
	"ScopeDoesNotExistException":               newErrorScopeDoesNotExistException,
	"SoftwareTokenMFANotFoundException":        newErrorSoftwareTokenMFANotFoundException,
	"TooManyFailedAttemptsException":           newErrorTooManyFailedAttemptsException,
	"TooManyRequestsException":                 newErrorTooManyRequestsException,
	"UnauthorizedException":                    newErrorUnauthorizedException,
	"UnexpectedLambdaException":                newErrorUnexpectedLambdaException,
	"UnsupportedIdentityProviderException":     newErrorUnsupportedIdentityProviderException,
	"UnsupportedOperationException":            newErrorUnsupportedOperationException,
	"UnsupportedTokenTypeException":            newErrorUnsupportedTokenTypeException,
	"UnsupportedUserStateException":            newErrorUnsupportedUserStateException,
	"UserImportInProgressException":            newErrorUserImportInProgressException,
	"UserLambdaValidationException":            newErrorUserLambdaValidationException,
	"UserNotConfirmedException":                newErrorUserNotConfirmedException,
	"UserNotFoundException":                    newErrorUserNotFoundException,
	"UserPoolAddOnNotEnabledException":         newErrorUserPoolAddOnNotEnabledException,
	"UserPoolTaggingException":                 newErrorUserPoolTaggingException,
	"UsernameExistsException":                  newErrorUsernameExistsException,
}
