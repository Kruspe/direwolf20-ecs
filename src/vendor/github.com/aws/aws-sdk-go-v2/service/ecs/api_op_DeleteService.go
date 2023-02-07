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

// Deletes a specified service within a cluster. You can delete a service if you
// have no running tasks in it and the desired task count is zero. If the service
// is actively maintaining tasks, you can't delete it, and you must update the
// service to a desired task count of zero. For more information, see
// UpdateService. When you delete a service, if there are still running tasks that
// require cleanup, the service status moves from ACTIVE to DRAINING, and the
// service is no longer visible in the console or in the ListServices API
// operation. After all tasks have transitioned to either STOPPING or STOPPED
// status, the service status moves from DRAINING to INACTIVE. Services in the
// DRAINING or INACTIVE status can still be viewed with the DescribeServices API
// operation. However, in the future, INACTIVE services may be cleaned up and
// purged from Amazon ECS record keeping, and DescribeServices calls on those
// services return a ServiceNotFoundException error. If you attempt to create a new
// service with the same name as an existing service in either ACTIVE or DRAINING
// status, you receive an error.
func (c *Client) DeleteService(ctx context.Context, params *DeleteServiceInput, optFns ...func(*Options)) (*DeleteServiceOutput, error) {
	if params == nil {
		params = &DeleteServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteService", params, optFns, c.addOperationDeleteServiceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteServiceInput struct {

	// The name of the service to delete.
	//
	// This member is required.
	Service *string

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// service to delete. If you do not specify a cluster, the default cluster is
	// assumed.
	Cluster *string

	// If true, allows you to delete a service even if it wasn't scaled down to zero
	// tasks. It's only necessary to use this if the service uses the REPLICA
	// scheduling strategy.
	Force *bool

	noSmithyDocumentSerde
}

type DeleteServiceOutput struct {

	// The full description of the deleted service.
	Service *types.Service

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteService{}, middleware.After)
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
	if err = addOpDeleteServiceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteService(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteService(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DeleteService",
	}
}
