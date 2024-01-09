// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the contents of the encrypted fields SecretString or SecretBinary
// from the specified version of a secret, whichever contains content. We recommend
// that you cache your secret values by using client-side caching. Caching secrets
// improves speed and reduces your costs. For more information, see Cache secrets
// for your applications (https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieving-secrets.html)
// . To retrieve the previous version of a secret, use VersionStage and specify
// AWSPREVIOUS. To revert to the previous version of a secret, call
// UpdateSecretVersionStage (https://docs.aws.amazon.com/cli/latest/reference/secretsmanager/update-secret-version-stage.html)
// . Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see Logging Secrets Manager events with CloudTrail (https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html)
// . Required permissions: secretsmanager:GetSecretValue . If the secret is
// encrypted using a customer-managed key instead of the Amazon Web Services
// managed key aws/secretsmanager , then you also need kms:Decrypt permissions for
// that key. For more information, see IAM policy actions for Secrets Manager (https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions)
// and Authentication and access control in Secrets Manager (https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html)
// .
func (c *Client) GetSecretValue(ctx context.Context, params *GetSecretValueInput, optFns ...func(*Options)) (*GetSecretValueOutput, error) {
	if params == nil {
		params = &GetSecretValueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSecretValue", params, optFns, c.addOperationGetSecretValueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSecretValueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSecretValueInput struct {

	// The ARN or name of the secret to retrieve. For an ARN, we recommend that you
	// specify a complete ARN rather than a partial ARN. See Finding a secret from a
	// partial ARN (https://docs.aws.amazon.com/secretsmanager/latest/userguide/troubleshoot.html#ARN_secretnamehyphen)
	// .
	//
	// This member is required.
	SecretId *string

	// The unique identifier of the version of the secret to retrieve. If you include
	// both this parameter and VersionStage , the two parameters must refer to the same
	// secret version. If you don't specify either a VersionStage or VersionId , then
	// Secrets Manager returns the AWSCURRENT version. This value is typically a
	// UUID-type (https://wikipedia.org/wiki/Universally_unique_identifier) value with
	// 32 hexadecimal digits.
	VersionId *string

	// The staging label of the version of the secret to retrieve. Secrets Manager
	// uses staging labels to keep track of different versions during the rotation
	// process. If you include both this parameter and VersionId , the two parameters
	// must refer to the same secret version. If you don't specify either a
	// VersionStage or VersionId , Secrets Manager returns the AWSCURRENT version.
	VersionStage *string

	noSmithyDocumentSerde
}

type GetSecretValueOutput struct {

	// The ARN of the secret.
	ARN *string

	// The date and time that this version of the secret was created. If you don't
	// specify which version in VersionId or VersionStage , then Secrets Manager uses
	// the AWSCURRENT version.
	CreatedDate *time.Time

	// The friendly name of the secret.
	Name *string

	// The decrypted secret value, if the secret value was originally provided as
	// binary data in the form of a byte array. The response parameter represents the
	// binary data as a base64-encoded (https://tools.ietf.org/html/rfc4648#section-4)
	// string. If the secret was created by using the Secrets Manager console, or if
	// the secret value was originally provided as a string, then this field is
	// omitted. The secret value appears in SecretString instead.
	SecretBinary []byte

	// The decrypted secret value, if the secret value was originally provided as a
	// string or through the Secrets Manager console. If this secret was created by
	// using the console, then Secrets Manager stores the information as a JSON
	// structure of key/value pairs.
	SecretString *string

	// The unique identifier of this version of the secret.
	VersionId *string

	// A list of all of the staging labels currently attached to this version of the
	// secret.
	VersionStages []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSecretValueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSecretValue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSecretValue{}, middleware.After)
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
	if err = addGetSecretValueResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetSecretValueValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSecretValue(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSecretValue(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "secretsmanager",
		OperationName: "GetSecretValue",
	}
}

type opGetSecretValueResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetSecretValueResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetSecretValueResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "secretsmanager"
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
				signingName = "secretsmanager"
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
				v4aScheme.SigningName = aws.String("secretsmanager")
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

func addGetSecretValueResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetSecretValueResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
