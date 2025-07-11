// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package opensearchserverless

import (
	"context"
	"unique"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/service/opensearchserverless"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	inttypes "github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/internal/vcr"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*inttypes.ServicePackageFrameworkDataSource {
	return []*inttypes.ServicePackageFrameworkDataSource{
		{
			Factory:  newAccessPolicyDataSource,
			TypeName: "aws_opensearchserverless_access_policy",
			Name:     "Access Policy",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  newCollectionDataSource,
			TypeName: "aws_opensearchserverless_collection",
			Name:     "Collection",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  newLifecyclePolicyDataSource,
			TypeName: "aws_opensearchserverless_lifecycle_policy",
			Name:     "Lifecycle Policy",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  newSecurityConfigDataSource,
			TypeName: "aws_opensearchserverless_security_config",
			Name:     "Security Config",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*inttypes.ServicePackageFrameworkResource {
	return []*inttypes.ServicePackageFrameworkResource{
		{
			Factory:  newAccessPolicyResource,
			TypeName: "aws_opensearchserverless_access_policy",
			Name:     "Access Policy",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  newCollectionResource,
			TypeName: "aws_opensearchserverless_collection",
			Name:     "Collection",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  newLifecyclePolicyResource,
			TypeName: "aws_opensearchserverless_lifecycle_policy",
			Name:     "Lifecycle Policy",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  newSecurityConfigResource,
			TypeName: "aws_opensearchserverless_security_config",
			Name:     "Security Config",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  newSecurityPolicyResource,
			TypeName: "aws_opensearchserverless_security_policy",
			Name:     "Security Policy",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  newVPCEndpointResource,
			TypeName: "aws_opensearchserverless_vpc_endpoint",
			Name:     "VPC Endpoint",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*inttypes.ServicePackageSDKDataSource {
	return []*inttypes.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceSecurityPolicy,
			TypeName: "aws_opensearchserverless_security_policy",
			Name:     "Security Policy",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  dataSourceVPCEndpoint,
			TypeName: "aws_opensearchserverless_vpc_endpoint",
			Name:     "VPC Endpoint",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*inttypes.ServicePackageSDKResource {
	return []*inttypes.ServicePackageSDKResource{}
}

func (p *servicePackage) ServicePackageName() string {
	return names.OpenSearchServerless
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*opensearchserverless.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*opensearchserverless.Options){
		opensearchserverless.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *opensearchserverless.Options) {
			if region := config[names.AttrRegion].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         p.ServicePackageName(),
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		func(o *opensearchserverless.Options) {
			if inContext, ok := conns.FromContext(ctx); ok && inContext.VCREnabled() {
				tflog.Info(ctx, "overriding retry behavior to immediately return VCR errors")
				o.Retryer = conns.AddIsErrorRetryables(cfg.Retryer().(aws.RetryerV2), retry.IsErrorRetryableFunc(vcr.InteractionNotFoundRetryableFunc))
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return opensearchserverless.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*opensearchserverless.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*opensearchserverless.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *opensearchserverless.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*opensearchserverless.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
