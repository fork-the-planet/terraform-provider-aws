// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53domains

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/hashicorp/aws-sdk-go-base/v2/endpoints"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*route53domains.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))

	return route53domains.NewFromConfig(cfg,
		route53domains.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *route53domains.Options) {
			if config["partition"].(string) == endpoints.AwsPartitionID {
				// Route 53 Domains is only available in AWS Commercial us-east-1 Region.
				if cfg.Region != endpoints.UsEast1RegionID {
					tflog.Info(ctx, "overriding region", map[string]any{
						"original_region": cfg.Region,
						"override_region": endpoints.UsEast1RegionID,
					})
				}
				o.Region = endpoints.UsEast1RegionID
			}
		},
	), nil
}
