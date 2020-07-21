// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqs

const (

	// ErrCodeBatchEntryIdsNotDistinct for service response error code
	// "AWS.SimpleQueueService.BatchEntryIdsNotDistinct".
	//
	// Two or more batch entries in the request have the same Id.
	ErrCodeBatchEntryIdsNotDistinct = "AWS.SimpleQueueService.BatchEntryIdsNotDistinct"

	// ErrCodeBatchRequestTooLong for service response error code
	// "AWS.SimpleQueueService.BatchRequestTooLong".
	//
	// The length of all the messages put together is more than the limit.
	ErrCodeBatchRequestTooLong = "AWS.SimpleQueueService.BatchRequestTooLong"

	// ErrCodeEmptyBatchRequest for service response error code
	// "AWS.SimpleQueueService.EmptyBatchRequest".
	//
	// The batch request doesn't contain any entries.
	ErrCodeEmptyBatchRequest = "AWS.SimpleQueueService.EmptyBatchRequest"

	// ErrCodeInvalidAttributeName for service response error code
	// "InvalidAttributeName".
	//
	// The specified attribute doesn't exist.
	ErrCodeInvalidAttributeName = "InvalidAttributeName"

	// ErrCodeInvalidBatchEntryId for service response error code
	// "AWS.SimpleQueueService.InvalidBatchEntryId".
	//
	// The Id of a batch entry in a batch request doesn't abide by the specification.
	ErrCodeInvalidBatchEntryId = "AWS.SimpleQueueService.InvalidBatchEntryId"

	// ErrCodeInvalidIdFormat for service response error code
	// "InvalidIdFormat".
	//
	// The specified receipt handle isn't valid for the current version.
	ErrCodeInvalidIdFormat = "InvalidIdFormat"

	// ErrCodeInvalidMessageContents for service response error code
	// "InvalidMessageContents".
	//
	// The message contains characters outside the allowed set.
	ErrCodeInvalidMessageContents = "InvalidMessageContents"

	// ErrCodeMessageNotInflight for service response error code
	// "AWS.SimpleQueueService.MessageNotInflight".
	//
	// The specified message isn't in flight.
	ErrCodeMessageNotInflight = "AWS.SimpleQueueService.MessageNotInflight"

	// ErrCodeOverLimit for service response error code
	// "OverLimit".
	//
	// The specified action violates a limit. For example, ReceiveMessage returns
	// this error if the maximum number of inflight messages is reached and AddPermission
	// returns this error if the maximum number of permissions for the queue is
	// reached.
	ErrCodeOverLimit = "OverLimit"

	// ErrCodePurgeQueueInProgress for service response error code
	// "AWS.SimpleQueueService.PurgeQueueInProgress".
	//
	// Indicates that the specified queue previously received a PurgeQueue request
	// within the last 60 seconds (the time it can take to delete the messages in
	// the queue).
	ErrCodePurgeQueueInProgress = "AWS.SimpleQueueService.PurgeQueueInProgress"

	// ErrCodeQueueDeletedRecently for service response error code
	// "AWS.SimpleQueueService.QueueDeletedRecently".
	//
	// You must wait 60 seconds after deleting a queue before you can create another
	// queue with the same name.
	ErrCodeQueueDeletedRecently = "AWS.SimpleQueueService.QueueDeletedRecently"

	// ErrCodeQueueDoesNotExist for service response error code
	// "AWS.SimpleQueueService.NonExistentQueue".
	//
	// The specified queue doesn't exist.
	ErrCodeQueueDoesNotExist = "AWS.SimpleQueueService.NonExistentQueue"

	// ErrCodeQueueNameExists for service response error code
	// "QueueAlreadyExists".
	//
	// A queue with this name already exists. Amazon SQS returns this error only
	// if the request includes attributes whose values differ from those of the
	// existing queue.
	ErrCodeQueueNameExists = "QueueAlreadyExists"

	// ErrCodeReceiptHandleIsInvalid for service response error code
	// "ReceiptHandleIsInvalid".
	//
	// The specified receipt handle isn't valid.
	ErrCodeReceiptHandleIsInvalid = "ReceiptHandleIsInvalid"

	// ErrCodeTooManyEntriesInBatchRequest for service response error code
	// "AWS.SimpleQueueService.TooManyEntriesInBatchRequest".
	//
	// The batch request contains more entries than permissible.
	ErrCodeTooManyEntriesInBatchRequest = "AWS.SimpleQueueService.TooManyEntriesInBatchRequest"

	// ErrCodeUnsupportedOperation for service response error code
	// "AWS.SimpleQueueService.UnsupportedOperation".
	//
	// Error code 400. Unsupported operation.
	ErrCodeUnsupportedOperation = "AWS.SimpleQueueService.UnsupportedOperation"
)
