// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apprunner_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	tfknownvalue "github.com/hashicorp/terraform-provider-aws/internal/acctest/knownvalue"
	tfstatecheck "github.com/hashicorp/terraform-provider-aws/internal/acctest/statecheck"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfapprunner "github.com/hashicorp/terraform-provider-aws/internal/service/apprunner"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccAppRunnerVPCIngressConnection_basic(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_apprunner_vpc_ingress_connection.test"
	vpcResourceName := "aws_vpc.test"
	vpcEndpointResourceName := "aws_vpc_endpoint.test"
	appRunnerServiceResourceName := "aws_apprunner_service.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AppRunnerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVPCIngressConnectionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCIngressConnectionConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVPCIngressConnectionExists(ctx, resourceName),
					acctest.MatchResourceAttrRegionalARN(ctx, resourceName, names.AttrARN, "apprunner", regexache.MustCompile(fmt.Sprintf(`vpcingressconnection/%s/.+`, rName))),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					resource.TestCheckResourceAttr(resourceName, names.AttrStatus, string(types.VpcIngressConnectionStatusAvailable)),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrDomainName),
					resource.TestCheckResourceAttrPair(resourceName, "service_arn", appRunnerServiceResourceName, names.AttrARN),
					resource.TestCheckResourceAttrPair(resourceName, "ingress_vpc_configuration.0.vpc_id", vpcResourceName, names.AttrID),
					resource.TestCheckResourceAttrPair(resourceName, "ingress_vpc_configuration.0.vpc_endpoint_id", vpcEndpointResourceName, names.AttrID),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAppRunnerVPCIngressConnection_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_apprunner_vpc_ingress_connection.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AppRunnerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVPCIngressConnectionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCIngressConnectionConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVPCIngressConnectionExists(ctx, resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfapprunner.ResourceVPCIngressConnection(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAppRunnerVPCIngressConnection_Identity_ExistingResource(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_apprunner_vpc_ingress_connection.test"

	resource.ParallelTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_12_0),
		},
		PreCheck:     func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.AppRunnerServiceID),
		CheckDestroy: testAccCheckVPCIngressConnectionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"aws": {
						Source:            "hashicorp/aws",
						VersionConstraint: "5.100.0",
					},
				},
				Config: testAccVPCIngressConnectionConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVPCIngressConnectionExists(ctx, resourceName),
				),
				ConfigStateChecks: []statecheck.StateCheck{
					tfstatecheck.ExpectNoIdentity(resourceName),
				},
			},
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"aws": {
						Source:            "hashicorp/aws",
						VersionConstraint: "6.0.0",
					},
				},
				Config: testAccVPCIngressConnectionConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVPCIngressConnectionExists(ctx, resourceName),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectIdentity(resourceName, map[string]knownvalue.Check{
						names.AttrARN: knownvalue.Null(),
					}),
				},
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				Config:                   testAccVPCIngressConnectionConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVPCIngressConnectionExists(ctx, resourceName),
				),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectIdentity(resourceName, map[string]knownvalue.Check{
						names.AttrARN: tfknownvalue.RegionalARNRegexp("apprunner", regexache.MustCompile(`vpcingressconnection/.+`)),
					}),
				},
			},
		},
	})
}

func testAccCheckVPCIngressConnectionDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_apprunner_vpc_ingress_connection" {
				continue
			}

			conn := acctest.Provider.Meta().(*conns.AWSClient).AppRunnerClient(ctx)

			_, err := tfapprunner.FindVPCIngressConnectionByARN(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("App Runner VPC Ingress Connection %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckVPCIngressConnectionExists(ctx context.Context, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).AppRunnerClient(ctx)

		_, err := tfapprunner.FindVPCIngressConnectionByARN(ctx, conn, rs.Primary.ID)

		return err
	}
}

func testAccVPCIngressConnectionConfig_base(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
data "aws_region" "current" {}

resource "aws_apprunner_service" "test" {
  service_name = %[1]q

  source_configuration {
    image_repository {
      image_configuration {
        port = "8000"
      }
      image_identifier      = "public.ecr.aws/aws-containers/hello-app-runner:latest"
      image_repository_type = "ECR_PUBLIC"
    }
    auto_deployments_enabled = false
  }

  network_configuration {
    ingress_configuration {
      is_publicly_accessible = false
    }
  }
}

resource "aws_vpc_endpoint" "test" {
  vpc_id            = aws_vpc.test.id
  service_name      = "com.amazonaws.${data.aws_region.current.name}.apprunner.requests"
  vpc_endpoint_type = "Interface"

  subnet_ids = aws_subnet.test[*].id

  security_group_ids = [
    aws_vpc.test.default_security_group_id,
  ]

  tags = {
    Name = %[1]q
  }
}
`, rName))
}

func testAccVPCIngressConnectionConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccVPCIngressConnectionConfig_base(rName), fmt.Sprintf(`
resource "aws_apprunner_vpc_ingress_connection" "test" {
  name        = %[1]q
  service_arn = aws_apprunner_service.test.arn

  ingress_vpc_configuration {
    vpc_id          = aws_vpc.test.id
    vpc_endpoint_id = aws_vpc_endpoint.test.id
  }
}
`, rName))
}
