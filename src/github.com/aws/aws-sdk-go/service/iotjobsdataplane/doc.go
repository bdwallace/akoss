// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iotjobsdataplane provides the client and types for making API
// requests to AWS IoT Jobs Data Plane.
//
// AWS IoT Jobs is a service that allows you to define a set of jobs — remote
// operations that are sent to and executed on one or more devices connected
// to AWS IoT. For example, you can define a job that instructs a set of devices
// to download and install application or firmware updates, reboot, rotate certificates,
// or perform remote troubleshooting operations.
//
// To create a job, you make a job document which is a description of the remote
// operations to be performed, and you specify a list of targets that should
// perform the operations. The targets can be individual things, thing groups
// or both.
//
// AWS IoT Jobs sends a message to inform the targets that a job is available.
// The target starts the execution of the job by downloading the job document,
// performing the operations it specifies, and reporting its progress to AWS
// IoT. The Jobs service provides commands to track the progress of a job on
// a specific target and for all the targets of the job
//
// See https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29 for more information on this service.
//
// See iotjobsdataplane package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/iotjobsdataplane/
//
// Using the Client
//
// To contact AWS IoT Jobs Data Plane with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS IoT Jobs Data Plane client IoTJobsDataPlane for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/iotjobsdataplane/#New
package iotjobsdataplane
