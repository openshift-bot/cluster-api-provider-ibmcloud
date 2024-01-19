// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the modifications made to your Reserved Instances. If no parameter is
// specified, information about all your Reserved Instances modification requests
// is returned. If a modification ID is specified, only information about the
// specific modification is returned. For more information, see Modifying Reserved
// Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-modifying.html)
// in the Amazon EC2 User Guide.
func (c *Client) DescribeReservedInstancesModifications(ctx context.Context, params *DescribeReservedInstancesModificationsInput, optFns ...func(*Options)) (*DescribeReservedInstancesModificationsOutput, error) {
	if params == nil {
		params = &DescribeReservedInstancesModificationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReservedInstancesModifications", params, optFns, c.addOperationDescribeReservedInstancesModificationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReservedInstancesModificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeReservedInstancesModifications.
type DescribeReservedInstancesModificationsInput struct {

	// One or more filters.
	//   - client-token - The idempotency token for the modification request.
	//   - create-date - The time when the modification request was created.
	//   - effective-date - The time when the modification becomes effective.
	//   - modification-result.reserved-instances-id - The ID for the Reserved
	//   Instances created as part of the modification request. This ID is only available
	//   when the status of the modification is fulfilled .
	//   - modification-result.target-configuration.availability-zone - The
	//   Availability Zone for the new Reserved Instances.
	//   - modification-result.target-configuration.instance-count - The number of new
	//   Reserved Instances.
	//   - modification-result.target-configuration.instance-type - The instance type
	//   of the new Reserved Instances.
	//   - reserved-instances-id - The ID of the Reserved Instances modified.
	//   - reserved-instances-modification-id - The ID of the modification request.
	//   - status - The status of the Reserved Instances modification request (
	//   processing | fulfilled | failed ).
	//   - status-message - The reason for the status.
	//   - update-date - The time when the modification request was last updated.
	Filters []types.Filter

	// The token to retrieve the next page of results.
	NextToken *string

	// IDs for the submitted modification request.
	ReservedInstancesModificationIds []string

	noSmithyDocumentSerde
}

// Contains the output of DescribeReservedInstancesModifications.
type DescribeReservedInstancesModificationsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The Reserved Instance modification information.
	ReservedInstancesModifications []types.ReservedInstancesModification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReservedInstancesModificationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeReservedInstancesModifications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeReservedInstancesModifications{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addDescribeReservedInstancesModificationsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedInstancesModifications(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeReservedInstancesModificationsAPIClient is a client that implements the
// DescribeReservedInstancesModifications operation.
type DescribeReservedInstancesModificationsAPIClient interface {
	DescribeReservedInstancesModifications(context.Context, *DescribeReservedInstancesModificationsInput, ...func(*Options)) (*DescribeReservedInstancesModificationsOutput, error)
}

var _ DescribeReservedInstancesModificationsAPIClient = (*Client)(nil)

// DescribeReservedInstancesModificationsPaginatorOptions is the paginator options
// for DescribeReservedInstancesModifications
type DescribeReservedInstancesModificationsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReservedInstancesModificationsPaginator is a paginator for
// DescribeReservedInstancesModifications
type DescribeReservedInstancesModificationsPaginator struct {
	options   DescribeReservedInstancesModificationsPaginatorOptions
	client    DescribeReservedInstancesModificationsAPIClient
	params    *DescribeReservedInstancesModificationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeReservedInstancesModificationsPaginator returns a new
// DescribeReservedInstancesModificationsPaginator
func NewDescribeReservedInstancesModificationsPaginator(client DescribeReservedInstancesModificationsAPIClient, params *DescribeReservedInstancesModificationsInput, optFns ...func(*DescribeReservedInstancesModificationsPaginatorOptions)) *DescribeReservedInstancesModificationsPaginator {
	if params == nil {
		params = &DescribeReservedInstancesModificationsInput{}
	}

	options := DescribeReservedInstancesModificationsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReservedInstancesModificationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReservedInstancesModificationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReservedInstancesModifications page.
func (p *DescribeReservedInstancesModificationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReservedInstancesModificationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.DescribeReservedInstancesModifications(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeReservedInstancesModifications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeReservedInstancesModifications",
	}
}

type opDescribeReservedInstancesModificationsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeReservedInstancesModificationsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeReservedInstancesModificationsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "ec2"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "ec2"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("ec2")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addDescribeReservedInstancesModificationsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeReservedInstancesModificationsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
