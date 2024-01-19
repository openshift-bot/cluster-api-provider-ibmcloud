// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Uploads a batch of log events to the specified log stream. The sequence token
// is now ignored in PutLogEvents actions. PutLogEvents actions are always
// accepted and never return InvalidSequenceTokenException or
// DataAlreadyAcceptedException even if the sequence token is not valid. You can
// use parallel PutLogEvents actions on the same log stream. The batch of events
// must satisfy the following constraints:
//   - The maximum batch size is 1,048,576 bytes. This size is calculated as the
//     sum of all event messages in UTF-8, plus 26 bytes for each log event.
//   - None of the log events in the batch can be more than 2 hours in the future.
//   - None of the log events in the batch can be more than 14 days in the past.
//     Also, none of the log events can be from earlier than the retention period of
//     the log group.
//   - The log events in the batch must be in chronological order by their
//     timestamp. The timestamp is the time that the event occurred, expressed as the
//     number of milliseconds after Jan 1, 1970 00:00:00 UTC . (In Amazon Web
//     Services Tools for PowerShell and the Amazon Web Services SDK for .NET, the
//     timestamp is specified in .NET format: yyyy-mm-ddThh:mm:ss . For example,
//     2017-09-15T13:45:30 .)
//   - A batch of log events in a single request cannot span more than 24 hours.
//     Otherwise, the operation fails.
//   - Each log event can be no larger than 256 KB.
//   - The maximum number of log events in a batch is 10,000.
//   - The quota of five requests per second per log stream has been removed.
//     Instead, PutLogEvents actions are throttled based on a per-second per-account
//     quota. You can request an increase to the per-second throttling quota by using
//     the Service Quotas service.
//
// If a call to PutLogEvents returns "UnrecognizedClientException" the most likely
// cause is a non-valid Amazon Web Services access key ID or secret key.
func (c *Client) PutLogEvents(ctx context.Context, params *PutLogEventsInput, optFns ...func(*Options)) (*PutLogEventsOutput, error) {
	if params == nil {
		params = &PutLogEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutLogEvents", params, optFns, c.addOperationPutLogEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutLogEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLogEventsInput struct {

	// The log events.
	//
	// This member is required.
	LogEvents []types.InputLogEvent

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string

	// The name of the log stream.
	//
	// This member is required.
	LogStreamName *string

	// The sequence token obtained from the response of the previous PutLogEvents
	// call. The sequenceToken parameter is now ignored in PutLogEvents actions.
	// PutLogEvents actions are now accepted and never return
	// InvalidSequenceTokenException or DataAlreadyAcceptedException even if the
	// sequence token is not valid.
	SequenceToken *string

	noSmithyDocumentSerde
}

type PutLogEventsOutput struct {

	// The next sequence token. This field has been deprecated. The sequence token is
	// now ignored in PutLogEvents actions. PutLogEvents actions are always accepted
	// even if the sequence token is not valid. You can use parallel PutLogEvents
	// actions on the same log stream and you do not need to wait for the response of a
	// previous PutLogEvents action to obtain the nextSequenceToken value.
	NextSequenceToken *string

	// The rejected events.
	RejectedLogEventsInfo *types.RejectedLogEventsInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutLogEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutLogEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutLogEvents{}, middleware.After)
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
	if err = addPutLogEventsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutLogEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutLogEvents(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutLogEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "PutLogEvents",
	}
}

type opPutLogEventsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opPutLogEventsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opPutLogEventsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "logs"
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
				signingName = "logs"
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
				v4aScheme.SigningName = aws.String("logs")
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

func addPutLogEventsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opPutLogEventsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
