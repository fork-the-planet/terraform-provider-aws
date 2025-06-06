// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oam_test

import (
	"fmt"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func testAccObservabilityAccessManagerLinkDataSource_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_oam_link.test"
	resourceName := "aws_oam_link.test"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckAlternateAccount(t)
			acctest.PreCheckPartitionHasService(t, names.ObservabilityAccessManagerEndpointID)
			testAccPreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.ObservabilityAccessManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5FactoriesAlternate(ctx, t),
		Steps: []resource.TestStep{
			{
				Config: testAccLinkDataSourceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, names.AttrARN, resourceName, names.AttrARN),
					resource.TestCheckResourceAttrPair(dataSourceName, "label", resourceName, "label"),
					resource.TestCheckResourceAttrPair(dataSourceName, "label_template", resourceName, "label_template"),
					resource.TestCheckResourceAttrPair(dataSourceName, "link_id", resourceName, "link_id"),
					resource.TestCheckResourceAttr(dataSourceName, "resource_types.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "resource_types.0", "AWS::CloudWatch::Metric"),
					resource.TestCheckResourceAttrPair(dataSourceName, "sink_arn", resourceName, "sink_arn"),
					resource.TestCheckResourceAttr(dataSourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttr(dataSourceName, acctest.CtTagsKey1, acctest.CtValue1),
				),
			},
		},
	})
}

func testAccObservabilityAccessManagerLinkDataSource_logGroupConfiguration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_oam_link.test"
	resourceName := "aws_oam_link.test"
	filter := "LogGroupName LIKE 'aws/lambda/%' OR LogGroupName LIKE 'AWSLogs%'"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckAlternateAccount(t)
			acctest.PreCheckPartitionHasService(t, names.ObservabilityAccessManagerEndpointID)
			testAccPreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.ObservabilityAccessManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5FactoriesAlternate(ctx, t),
		Steps: []resource.TestStep{
			{
				Config: testAccLinkDataSourceConfig_logGroupConfiguration(rName, filter),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, names.AttrARN, resourceName, names.AttrARN),
					resource.TestCheckResourceAttrPair(dataSourceName, "label", resourceName, "label"),
					resource.TestCheckResourceAttr(dataSourceName, "label_template", "$AccountName"),
					resource.TestCheckResourceAttr(dataSourceName, "link_configuration.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "link_configuration.0.log_group_configuration.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "link_configuration.0.log_group_configuration.0.filter", filter),
					resource.TestCheckResourceAttr(dataSourceName, "link_configuration.0.metric_configuration.#", "0"),
					resource.TestCheckResourceAttrPair(dataSourceName, "link_id", resourceName, "link_id"),
					resource.TestCheckResourceAttr(dataSourceName, "resource_types.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "resource_types.0", "AWS::Logs::LogGroup"),
					resource.TestCheckResourceAttrPair(dataSourceName, "sink_arn", resourceName, "sink_arn"),
					resource.TestCheckResourceAttr(dataSourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttr(dataSourceName, acctest.CtTagsKey1, acctest.CtValue1),
				),
			},
		},
	})
}

func testAccObservabilityAccessManagerLinkDataSource_metricConfiguration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_oam_link.test"
	resourceName := "aws_oam_link.test"
	filter := "Namespace IN ('AWS/EC2', 'AWS/ELB', 'AWS/S3')"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckAlternateAccount(t)
			acctest.PreCheckPartitionHasService(t, names.ObservabilityAccessManagerEndpointID)
			testAccPreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.ObservabilityAccessManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5FactoriesAlternate(ctx, t),
		Steps: []resource.TestStep{
			{
				Config: testAccLinkDataSourceConfig_metricConfiguration(rName, filter),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, names.AttrARN, resourceName, names.AttrARN),
					resource.TestCheckResourceAttrPair(dataSourceName, "label", resourceName, "label"),
					resource.TestCheckResourceAttr(dataSourceName, "label_template", "$AccountName"),
					resource.TestCheckResourceAttr(dataSourceName, "link_configuration.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "link_configuration.0.log_group_configuration.#", "0"),
					resource.TestCheckResourceAttr(dataSourceName, "link_configuration.0.metric_configuration.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "link_configuration.0.metric_configuration.0.filter", filter),
					resource.TestCheckResourceAttrPair(dataSourceName, "link_id", resourceName, "link_id"),
					resource.TestCheckResourceAttr(dataSourceName, "resource_types.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "resource_types.0", "AWS::CloudWatch::Metric"),
					resource.TestCheckResourceAttrPair(dataSourceName, "sink_arn", resourceName, "sink_arn"),
					resource.TestCheckResourceAttr(dataSourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttr(dataSourceName, acctest.CtTagsKey1, acctest.CtValue1),
				),
			},
		},
	})
}

func testAccLinkDataSourceConfig_basic(rName string) string {
	return acctest.ConfigCompose(
		acctest.ConfigAlternateAccountProvider(),
		fmt.Sprintf(`
data "aws_caller_identity" "source" {}
data "aws_partition" "source" {}

data "aws_caller_identity" "monitoring" {
  provider = "awsalternate"
}
data "aws_partition" "monitoring" {
  provider = "awsalternate"
}

resource "aws_oam_sink" "test" {
  provider = "awsalternate"

  name = %[1]q
}

resource "aws_oam_sink_policy" "test" {
  provider = "awsalternate"

  sink_identifier = aws_oam_sink.test.arn
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["oam:CreateLink", "oam:UpdateLink"]
        Effect   = "Allow"
        Resource = "*"
        Principal = {
          "AWS" = "arn:${data.aws_partition.source.partition}:iam::${data.aws_caller_identity.source.account_id}:root"
        }
        Condition = {
          "ForAnyValue:StringEquals" = {
            "oam:ResourceTypes" = ["AWS::CloudWatch::Metric", "AWS::Logs::LogGroup"]
          }
        }
      }
    ]
  })
}

resource "aws_oam_link" "test" {
  label_template  = "$AccountName"
  resource_types  = ["AWS::CloudWatch::Metric"]
  sink_identifier = aws_oam_sink.test.arn

  tags = {
    key1 = "value1"
  }

  depends_on = [
    aws_oam_sink_policy.test
  ]
}

data aws_oam_link "test" {
  link_identifier = aws_oam_link.test.arn
}
`, rName))
}

func testAccLinkDataSourceConfig_logGroupConfiguration(rName, filter string) string {
	return acctest.ConfigCompose(
		acctest.ConfigAlternateAccountProvider(),
		fmt.Sprintf(`
data "aws_caller_identity" "source" {}
data "aws_partition" "source" {}

data "aws_caller_identity" "monitoring" {
  provider = "awsalternate"
}
data "aws_partition" "monitoring" {
  provider = "awsalternate"
}

resource "aws_oam_sink" "test" {
  provider = "awsalternate"

  name = %[1]q
}

resource "aws_oam_sink_policy" "test" {
  provider = "awsalternate"

  sink_identifier = aws_oam_sink.test.arn
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["oam:CreateLink", "oam:UpdateLink"]
        Effect   = "Allow"
        Resource = "*"
        Principal = {
          "AWS" = "arn:${data.aws_partition.source.partition}:iam::${data.aws_caller_identity.source.account_id}:root"
        }
        Condition = {
          "ForAnyValue:StringEquals" = {
            "oam:ResourceTypes" = ["AWS::CloudWatch::Metric", "AWS::Logs::LogGroup"]
          }
        }
      }
    ]
  })
}

resource "aws_oam_link" "test" {
  label_template = "$AccountName"
  link_configuration {
    log_group_configuration {
      filter = %[2]q
    }
  }
  resource_types  = ["AWS::Logs::LogGroup"]
  sink_identifier = aws_oam_sink.test.arn

  tags = {
    key1 = "value1"
  }

  depends_on = [
    aws_oam_sink_policy.test
  ]
}

data aws_oam_link "test" {
  link_identifier = aws_oam_link.test.arn
}
`, rName, filter))
}

func testAccLinkDataSourceConfig_metricConfiguration(rName, filter string) string {
	return acctest.ConfigCompose(
		acctest.ConfigAlternateAccountProvider(),
		fmt.Sprintf(`
data "aws_caller_identity" "source" {}
data "aws_partition" "source" {}

data "aws_caller_identity" "monitoring" {
  provider = "awsalternate"
}
data "aws_partition" "monitoring" {
  provider = "awsalternate"
}

resource "aws_oam_sink" "test" {
  provider = "awsalternate"

  name = %[1]q
}

resource "aws_oam_sink_policy" "test" {
  provider = "awsalternate"

  sink_identifier = aws_oam_sink.test.arn
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["oam:CreateLink", "oam:UpdateLink"]
        Effect   = "Allow"
        Resource = "*"
        Principal = {
          "AWS" = "arn:${data.aws_partition.source.partition}:iam::${data.aws_caller_identity.source.account_id}:root"
        }
        Condition = {
          "ForAnyValue:StringEquals" = {
            "oam:ResourceTypes" = ["AWS::CloudWatch::Metric", "AWS::Logs::LogGroup"]
          }
        }
      }
    ]
  })
}

resource "aws_oam_link" "test" {
  label_template = "$AccountName"
  link_configuration {
    metric_configuration {
      filter = %[2]q
    }
  }
  resource_types  = ["AWS::CloudWatch::Metric"]
  sink_identifier = aws_oam_sink.test.arn

  tags = {
    key1 = "value1"
  }

  depends_on = [
    aws_oam_sink_policy.test
  ]
}

data aws_oam_link "test" {
  link_identifier = aws_oam_link.test.arn
}
`, rName, filter))
}
