// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package docdbelastic_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go-v2/service/docdbelastic"
	awstypes "github.com/aws/aws-sdk-go-v2/service/docdbelastic/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	tfstatecheck "github.com/hashicorp/terraform-provider-aws/internal/acctest/statecheck"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	tfdocdbelastic "github.com/hashicorp/terraform-provider-aws/internal/service/docdbelastic"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccDocDBElasticCluster_basic(t *testing.T) {
	ctx := acctest.Context(t)

	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var cluster awstypes.Cluster
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_docdbelastic_cluster.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccPreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.DocDBElasticServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckClusterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccClusterConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterExists(ctx, resourceName, &cluster),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					acctest.MatchResourceAttrRegionalARN(ctx, resourceName, names.AttrARN, "docdb-elastic", regexache.MustCompile(`cluster/`+verify.UUIDRegexPattern+`$`)),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrEndpoint),
					resource.TestCheckResourceAttrPair(resourceName, names.AttrID, resourceName, names.AttrARN),
					resource.TestCheckResourceAttr(resourceName, "shard_capacity", "2"),
					resource.TestCheckResourceAttr(resourceName, "shard_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "admin_user_name", "testuser"),
					resource.TestCheckResourceAttr(resourceName, "admin_user_password", "testpassword"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "vpc_security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, names.AttrPreferredMaintenanceWindow, "Tue:04:00-Tue:04:30"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"admin_user_password",
				},
			},
		},
	})
}

func TestAccDocDBElasticCluster_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var cluster awstypes.Cluster
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_docdbelastic_cluster.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccPreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.DocDBElasticServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckClusterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccClusterConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterExists(ctx, resourceName, &cluster),
					acctest.CheckFrameworkResourceDisappears(ctx, acctest.Provider, tfdocdbelastic.ResourceCluster, resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDocDBElasticCluster_tags(t *testing.T) {
	ctx := acctest.Context(t)

	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var cluster awstypes.Cluster
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_docdbelastic_cluster.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccPreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.DocDBElasticServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckClusterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccClusterConfig_tags1(rName, acctest.CtKey1, acctest.CtValue1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterExists(ctx, resourceName, &cluster),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsKey1, acctest.CtValue1),
				),
			},
			{
				Config: testAccClusterConfig_tags2(rName, acctest.CtKey1, acctest.CtValue1, acctest.CtKey2, acctest.CtValue2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterExists(ctx, resourceName, &cluster),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "2"),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsKey1, acctest.CtValue1),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsKey2, acctest.CtValue2),
				),
			},
			{
				Config: testAccClusterConfig_tags1(rName, acctest.CtKey2, acctest.CtValue2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterExists(ctx, resourceName, &cluster),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsKey2, acctest.CtValue2),
				),
			},
		},
	})
}

func TestAccDocDBElasticCluster_update(t *testing.T) {
	ctx := acctest.Context(t)

	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var cluster awstypes.Cluster
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_docdbelastic_cluster.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccPreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.DocDBElasticServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckClusterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccClusterConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterExists(ctx, resourceName, &cluster),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					acctest.MatchResourceAttrRegionalARN(ctx, resourceName, names.AttrARN, "docdb-elastic", regexache.MustCompile(`cluster/`+verify.UUIDRegexPattern+`$`)),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrEndpoint),
					resource.TestCheckResourceAttr(resourceName, "shard_capacity", "2"),
					resource.TestCheckResourceAttr(resourceName, "shard_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "admin_user_name", "testuser"),
					resource.TestCheckResourceAttr(resourceName, "admin_user_password", "testpassword"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "vpc_security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, names.AttrPreferredMaintenanceWindow, "Tue:04:00-Tue:04:30"),
				),
			},
			{
				Config: testAccClusterConfig_update(rName, 4, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterExists(ctx, resourceName, &cluster),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					acctest.MatchResourceAttrRegionalARN(ctx, resourceName, names.AttrARN, "docdb-elastic", regexache.MustCompile(`cluster/`+verify.UUIDRegexPattern+`$`)),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrEndpoint),
					resource.TestCheckResourceAttr(resourceName, "shard_capacity", "4"),
					resource.TestCheckResourceAttr(resourceName, "shard_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "admin_user_name", "testuser"),
					resource.TestCheckResourceAttr(resourceName, "admin_user_password", "testpassword"),
					resource.TestCheckResourceAttr(resourceName, "backup_retention_period", "1"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "vpc_security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, names.AttrPreferredMaintenanceWindow, "Tue:04:00-Tue:04:30"),
				),
			},
			{
				Config: testAccClusterConfig_update(rName, 4, 3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterExists(ctx, resourceName, &cluster),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					acctest.MatchResourceAttrRegionalARN(ctx, resourceName, names.AttrARN, "docdb-elastic", regexache.MustCompile(`cluster/`+verify.UUIDRegexPattern+`$`)),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrEndpoint),
					resource.TestCheckResourceAttr(resourceName, "shard_capacity", "4"),
					resource.TestCheckResourceAttr(resourceName, "shard_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "admin_user_name", "testuser"),
					resource.TestCheckResourceAttr(resourceName, "admin_user_password", "testpassword"),
					resource.TestCheckResourceAttr(resourceName, "backup_retention_period", "3"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "vpc_security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, names.AttrPreferredMaintenanceWindow, "Tue:04:00-Tue:04:30"),
				),
			},
		},
	})
}

func TestAccDocDBElasticCluster_Identity_ExistingResource(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_docdbelastic_cluster.test"

	resource.ParallelTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_12_0),
		},
		PreCheck:     func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:   acctest.ErrorCheck(t, names.DocDBElasticServiceID),
		CheckDestroy: testAccCheckClusterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"aws": {
						Source:            "hashicorp/aws",
						VersionConstraint: "5.100.0",
					},
				},
				Config: testAccClusterConfig_basic(rName),
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
				Config: testAccClusterConfig_basic(rName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectIdentityValueMatchesState(resourceName, tfjsonpath.New(names.AttrARN)),
				},
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				Config:                   testAccClusterConfig_basic(rName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionNoop),
					},
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectIdentityValueMatchesState(resourceName, tfjsonpath.New(names.AttrARN)),
				},
			},
		},
	})
}

func testAccCheckClusterDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).DocDBElasticClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_docdbelastic_cluster" {
				continue
			}

			_, err := tfdocdbelastic.FindClusterByID(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return create.Error(names.DocDBElastic, create.ErrActionCheckingDestroyed, tfdocdbelastic.ResNameCluster, rs.Primary.ID, errors.New("not destroyed"))
		}

		return nil
	}
}

func testAccCheckClusterExists(ctx context.Context, name string, cluster *awstypes.Cluster) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return create.Error(names.DocDBElastic, create.ErrActionCheckingExistence, tfdocdbelastic.ResNameCluster, name, errors.New("not found"))
		}

		if rs.Primary.ID == "" {
			return create.Error(names.DocDBElastic, create.ErrActionCheckingExistence, tfdocdbelastic.ResNameCluster, name, errors.New("not set"))
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).DocDBElasticClient(ctx)
		resp, err := tfdocdbelastic.FindClusterByID(ctx, conn, rs.Primary.ID)

		if err != nil {
			return create.Error(names.DocDBElastic, create.ErrActionCheckingExistence, tfdocdbelastic.ResNameCluster, rs.Primary.ID, err)
		}

		*cluster = *resp

		return nil
	}
}

func testAccPreCheck(ctx context.Context, t *testing.T) {
	conn := acctest.Provider.Meta().(*conns.AWSClient).DocDBElasticClient(ctx)

	input := &docdbelastic.ListClustersInput{}
	_, err := conn.ListClusters(ctx, input)

	if acctest.PreCheckSkipError(err) {
		t.Skipf("skipping acceptance testing: %s", err)
	}

	if err != nil {
		t.Fatalf("unexpected PreCheck error: %s", err)
	}
}

func testAccClusterBaseConfig(rName string) string {
	return acctest.ConfigCompose(
		acctest.ConfigVPCWithSubnets(rName, 2),
		fmt.Sprintf(`
resource "aws_security_group" "test" {
  name   = %[1]q
  vpc_id = aws_vpc.test.id
}
`, rName),
	)
}

func testAccClusterConfig_basic(rName string) string {
	return acctest.ConfigCompose(
		testAccClusterBaseConfig(rName),
		fmt.Sprintf(`
resource "aws_docdbelastic_cluster" "test" {
  name           = %[1]q
  shard_capacity = 2
  shard_count    = 1

  admin_user_name     = "testuser"
  admin_user_password = "testpassword"
  auth_type           = "PLAIN_TEXT"

  preferred_maintenance_window = "Tue:04:00-Tue:04:30"

  vpc_security_group_ids = [
    aws_security_group.test.id
  ]

  subnet_ids = [
    aws_subnet.test[0].id,
    aws_subnet.test[1].id
  ]
}
`, rName))
}

func testAccClusterConfig_update(rName string, shardCapacity, backupRetentionPeriod int) string {
	return acctest.ConfigCompose(
		testAccClusterBaseConfig(rName),
		fmt.Sprintf(`
resource "aws_docdbelastic_cluster" "test" {
  name           = %[1]q
  shard_capacity = %[2]d
  shard_count    = 1

  admin_user_name     = "testuser"
  admin_user_password = "testpassword"
  auth_type           = "PLAIN_TEXT"

  preferred_maintenance_window = "Tue:04:00-Tue:04:30"
  backup_retention_period      = %[3]d

  vpc_security_group_ids = [
    aws_security_group.test.id
  ]

  subnet_ids = [
    aws_subnet.test[0].id,
    aws_subnet.test[1].id
  ]
}
`, rName, shardCapacity, backupRetentionPeriod))
}

func testAccClusterConfig_tags1(rName, key1, value1 string) string {
	return acctest.ConfigCompose(
		testAccClusterBaseConfig(rName),
		fmt.Sprintf(`
resource "aws_docdbelastic_cluster" "test" {
  name           = %[1]q
  shard_capacity = 2
  shard_count    = 1

  admin_user_name     = "testuser"
  admin_user_password = "testpassword"
  auth_type           = "PLAIN_TEXT"

  preferred_maintenance_window = "Tue:04:00-Tue:04:30"

  vpc_security_group_ids = [
    aws_security_group.test.id
  ]

  subnet_ids = [
    aws_subnet.test[0].id,
    aws_subnet.test[1].id
  ]

  tags = {
    %[2]q = %[3]q
  }
}
`, rName, key1, value1))
}

func testAccClusterConfig_tags2(rName, key1, value1, key2, value2 string) string {
	return acctest.ConfigCompose(
		testAccClusterBaseConfig(rName),
		fmt.Sprintf(`
resource "aws_docdbelastic_cluster" "test" {
  name           = %[1]q
  shard_capacity = 2
  shard_count    = 1

  admin_user_name     = "testuser"
  admin_user_password = "testpassword"
  auth_type           = "PLAIN_TEXT"

  preferred_maintenance_window = "Tue:04:00-Tue:04:30"

  vpc_security_group_ids = [
    aws_security_group.test.id
  ]

  subnet_ids = [
    aws_subnet.test[0].id,
    aws_subnet.test[1].id
  ]

  tags = {
    %[2]q = %[3]q
    %[4]q = %[5]q
  }
}
`, rName, key1, value1, key2, value2))
}
