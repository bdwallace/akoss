// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codedeploy provides the client and types for making API
// requests to AWS CodeDeploy.
//
// AWS CodeDeploy is a deployment service that automates application deployments
// to Amazon EC2 instances, on-premises instances running in your own facility,
// serverless AWS Lambda functions, or applications in an Amazon ECS service.
//
// You can deploy a nearly unlimited variety of application content, such as
// an updated Lambda function, updated applications in an Amazon ECS service,
// code, web and configuration files, executables, packages, scripts, multimedia
// files, and so on. AWS CodeDeploy can deploy application content stored in
// Amazon S3 buckets, GitHub repositories, or Bitbucket repositories. You do
// not need to make changes to your existing code before you can use AWS CodeDeploy.
//
// AWS CodeDeploy makes it easier for you to rapidly release new features, helps
// you avoid downtime during application deployment, and handles the complexity
// of updating your applications, without many of the risks associated with
// error-prone manual deployments.
//
// AWS CodeDeploy Components
//
// Use the information in this guide to help you work with the following AWS
// CodeDeploy components:
//
//    * Application: A name that uniquely identifies the application you want
//    to deploy. AWS CodeDeploy uses this name, which functions as a container,
//    to ensure the correct combination of revision, deployment configuration,
//    and deployment group are referenced during a deployment.
//
//    * Deployment group: A set of individual instances, CodeDeploy Lambda deployment
//    configuration settings, or an Amazon ECS service and network details.
//    A Lambda deployment group specifies how to route traffic to a new version
//    of a Lambda function. An Amazon ECS deployment group specifies the service
//    created in Amazon ECS to deploy, a load balancer, and a listener to reroute
//    production traffic to an updated containerized application. An EC2/On-premises
//    deployment group contains individually tagged instances, Amazon EC2 instances
//    in Amazon EC2 Auto Scaling groups, or both. All deployment groups can
//    specify optional trigger, alarm, and rollback settings.
//
//    * Deployment configuration: A set of deployment rules and deployment success
//    and failure conditions used by AWS CodeDeploy during a deployment.
//
//    * Deployment: The process and the components used when updating a Lambda
//    function, a containerized application in an Amazon ECS service, or of
//    installing content on one or more instances.
//
//    * Application revisions: For an AWS Lambda deployment, this is an AppSpec
//    file that specifies the Lambda function to be updated and one or more
//    functions to validate deployment lifecycle events. For an Amazon ECS deployment,
//    this is an AppSpec file that specifies the Amazon ECS task definition,
//    container, and port where production traffic is rerouted. For an EC2/On-premises
//    deployment, this is an archive file that contains source content—source
//    code, webpages, executable files, and deployment scripts—along with
//    an AppSpec file. Revisions are stored in Amazon S3 buckets or GitHub repositories.
//    For Amazon S3, a revision is uniquely identified by its Amazon S3 object
//    key and its ETag, version, or both. For GitHub, a revision is uniquely
//    identified by its commit ID.
//
// This guide also contains information to help you get details about the instances
// in your deployments, to make on-premises instances available for AWS CodeDeploy
// deployments, to get details about a Lambda function deployment, and to get
// details about Amazon ECS service deployments.
//
// AWS CodeDeploy Information Resources
//
//    * AWS CodeDeploy User Guide (https://docs.aws.amazon.com/codedeploy/latest/userguide)
//
//    * AWS CodeDeploy API Reference Guide (https://docs.aws.amazon.com/codedeploy/latest/APIReference/)
//
//    * AWS CLI Reference for AWS CodeDeploy (https://docs.aws.amazon.com/cli/latest/reference/deploy/index.html)
//
//    * AWS CodeDeploy Developer Forum (https://forums.aws.amazon.com/forum.jspa?forumID=179)
//
// See https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06 for more information on this service.
//
// See codedeploy package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/codedeploy/
//
// Using the Client
//
// To contact AWS CodeDeploy with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS CodeDeploy client CodeDeploy for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/codedeploy/#New
package codedeploy
