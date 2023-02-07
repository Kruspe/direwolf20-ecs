// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deregisters an Amazon ECS container instance from the specified cluster. This
// instance is no longer available to run tasks. If you intend to use the container
// instance for some other purpose after deregistration, we recommend that you stop
// all of the tasks running on the container instance before deregistration. That
// prevents any orphaned tasks from consuming resources. Deregistering a container
// instance removes the instance from a cluster, but it doesn't terminate the EC2
// instance. If you are finished using the instance, be sure to terminate it in the
// Amazon EC2 console to stop billing. If you terminate a running container
// instance, Amazon ECS automatically deregisters the instance from your cluster
// (stopped container instances or instances with disconnected agents aren't
// automatically deregistered when terminated).
func (c *Client) DeregisterContainerInstance(ctx context.Context, params *DeregisterContainerInstanceInput, optFns ...func(*Options)) (*DeregisterContainerInstanceOutput, error) {
	if params == nil {
		params = &DeregisterContainerInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterContainerInstance", params, optFns, c.addOperationDeregisterContainerInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterContainerInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterContainerInstanceInput struct {

	// The container instance ID or full ARN of the container instance to deregister.
	// For more information about the ARN format, see Amazon Resource Name (ARN)
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids)
	// in the Amazon ECS Developer Guide.
	//
	// This member is required.
	ContainerInstance *string

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// container instance to deregister. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string

	// Forces the container instance to be deregistered. If you have tasks running on
	// the container instance when you deregister it with the force option, these tasks
	// remain running until you terminate the instance or the tasks stop through some
	// other means, but they're orphaned (no longer monitored or accounted for by
	// Amazon ECS). If an orphaned task on your container instance is part of an Amazon
	// ECS service, then the service scheduler starts another copy of that task, on a
	// different container instance if possible. Any containers in orphaned service
	// tasks that are registered with a Classic Load Balancer or an Application Load
	// Balancer target group are deregistered. They begin connection draining according
	// to the settings on the load balancer or target group.
	Force *bool

	noSmithyDocumentSerde
}

type DeregisterContainerInstanceOutput struct {

	// The container instance that was deregistered.
	ContainerInstance *types.ContainerInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterContainerInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterContainerInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterContainerInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeregisterContainerInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterContainerInstance(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeregisterContainerInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DeregisterContainerInstance",
	}
}
