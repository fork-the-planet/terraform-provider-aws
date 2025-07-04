// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rbin_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rbin"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfrbin "github.com/hashicorp/terraform-provider-aws/internal/service/rbin"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccRBinRule_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var rule rbin.GetRuleOutput
	description := "my test description"
	resourceType := "EBS_SNAPSHOT"
	resourceName := "aws_rbin_rule.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, rbin.ServiceID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, rbin.ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRuleDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccRuleConfig_basic1(description, resourceType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRuleExists(ctx, resourceName, &rule),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, description),
					resource.TestCheckResourceAttr(resourceName, "exclude_resource_tags.#", "0"),
					resource.TestCheckResourceAttr(resourceName, names.AttrResourceType, resourceType),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "retention_period.*", map[string]string{
						"retention_period_value": "10",
						"retention_period_unit":  "DAYS",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "resource_tags.*", map[string]string{
						"resource_tag_key":   "some_tag1",
						"resource_tag_value": "some_value1",
					}),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccRuleConfig_basic2(description, resourceType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRuleExists(ctx, resourceName, &rule),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, description),
					resource.TestCheckResourceAttr(resourceName, "exclude_resource_tags.#", "0"),
					resource.TestCheckResourceAttr(resourceName, names.AttrResourceType, resourceType),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "retention_period.*", map[string]string{
						"retention_period_value": "10",
						"retention_period_unit":  "DAYS",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "resource_tags.*", map[string]string{
						"resource_tag_key":   "some_tag3",
						"resource_tag_value": "some_value3",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "resource_tags.*", map[string]string{
						"resource_tag_key":   "some_tag4",
						"resource_tag_value": "some_value4",
					}),
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

func TestAccRBinRule_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var rbinrule rbin.GetRuleOutput
	description := "my test description"
	resourceType := "EBS_SNAPSHOT"
	resourceName := "aws_rbin_rule.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, rbin.ServiceID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, rbin.ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRuleDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccRuleConfig_basic1(description, resourceType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRuleExists(ctx, resourceName, &rbinrule),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfrbin.ResourceRule(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRBinRule_excludeResourceTags(t *testing.T) {
	ctx := acctest.Context(t)
	var rule rbin.GetRuleOutput
	description := "my test description"
	resourceType := "EBS_SNAPSHOT"
	resourceName := "aws_rbin_rule.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, rbin.ServiceID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, rbin.ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRuleDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccRuleConfig_excludeResourceTags1(description, resourceType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRuleExists(ctx, resourceName, &rule),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, description),
					resource.TestCheckResourceAttr(resourceName, names.AttrResourceType, resourceType),

					resource.TestCheckResourceAttr(resourceName, "exclude_resource_tags.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "exclude_resource_tags.*", map[string]string{
						"resource_tag_key":   "some_tag1",
						"resource_tag_value": "some_value1",
					}),
					resource.TestCheckResourceAttr(resourceName, "resource_tags.#", "0"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "retention_period.*", map[string]string{
						"retention_period_value": "10",
						"retention_period_unit":  "DAYS",
					}),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccRuleConfig_excludeResourceTags2(description, resourceType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRuleExists(ctx, resourceName, &rule),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, description),
					resource.TestCheckResourceAttr(resourceName, names.AttrResourceType, resourceType),
					resource.TestCheckResourceAttr(resourceName, "exclude_resource_tags.#", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "exclude_resource_tags.*", map[string]string{
						"resource_tag_key":   "some_tag3",
						"resource_tag_value": "some_value3",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "exclude_resource_tags.*", map[string]string{
						"resource_tag_key":   "some_tag4",
						"resource_tag_value": "some_value4",
					}),
					resource.TestCheckResourceAttr(resourceName, "resource_tags.#", "0"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "retention_period.*", map[string]string{
						"retention_period_value": "10",
						"retention_period_unit":  "DAYS",
					}),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccRuleConfig_basic1(description, resourceType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRuleExists(ctx, resourceName, &rule),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, description),
					resource.TestCheckResourceAttr(resourceName, "exclude_resource_tags.#", "0"),
					resource.TestCheckResourceAttr(resourceName, names.AttrResourceType, resourceType),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "retention_period.*", map[string]string{
						"retention_period_value": "10",
						"retention_period_unit":  "DAYS",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "resource_tags.*", map[string]string{
						"resource_tag_key":   "some_tag1",
						"resource_tag_value": "some_value1",
					}),
				),
			},
		},
	})
}

func TestAccRBinRule_lockConfig(t *testing.T) {
	ctx := acctest.Context(t)
	var rule rbin.GetRuleOutput
	resourceType := "EBS_SNAPSHOT"
	resourceName := "aws_rbin_rule.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, rbin.ServiceID)
		},
		ErrorCheck:               acctest.ErrorCheck(t, rbin.ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRuleDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccRuleConfig_lockConfig(resourceType, "DAYS", "7"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRuleExists(ctx, resourceName, &rule),
					resource.TestCheckResourceAttr(resourceName, "lock_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "lock_configuration.0.unlock_delay.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "lock_configuration.0.unlock_delay.0.unlock_delay_unit", "DAYS"),
					resource.TestCheckResourceAttr(resourceName, "lock_configuration.0.unlock_delay.0.unlock_delay_value", "7"),
				),
			},
		},
	})
}

func TestAccRBinRule_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var rule rbin.GetRuleOutput
	resourceType := "EBS_SNAPSHOT"
	resourceName := "aws_rbin_rule.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.RBin)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.RBin),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRuleDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccRuleConfig_tags1(resourceType, acctest.CtKey1, acctest.CtValue1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRuleExists(ctx, resourceName, &rule),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsKey1, acctest.CtValue1),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccRuleConfig_tags2(resourceType, acctest.CtKey1, acctest.CtValue1Updated, acctest.CtKey2, acctest.CtValue2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRuleExists(ctx, resourceName, &rule),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "2"),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsKey1, acctest.CtValue1Updated),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsKey2, acctest.CtValue2),
				),
			},
			{
				Config: testAccRuleConfig_tags1(resourceType, acctest.CtKey1, acctest.CtValue1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRuleExists(ctx, resourceName, &rule),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsKey1, acctest.CtValue1),
				),
			},
		},
	})
}

func testAccCheckRuleDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).RBinClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_rbin_rule" {
				continue
			}

			_, err := tfrbin.FindRuleByID(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("RBin Rule %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckRuleExists(ctx context.Context, n string, v *rbin.GetRuleOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).RBinClient(ctx)

		output, err := tfrbin.FindRuleByID(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccRuleConfig_basic1(description, resourceType string) string {
	return fmt.Sprintf(`
resource "aws_rbin_rule" "test" {
  description   = %[1]q
  resource_type = %[2]q

  resource_tags {
    resource_tag_key   = "some_tag1"
    resource_tag_value = "some_value1"
  }

  retention_period {
    retention_period_value = 10
    retention_period_unit  = "DAYS"
  }
}
`, description, resourceType)
}

func testAccRuleConfig_basic2(description, resourceType string) string {
	return fmt.Sprintf(`
resource "aws_rbin_rule" "test" {
  description   = %[1]q
  resource_type = %[2]q

  resource_tags {
    resource_tag_key   = "some_tag3"
    resource_tag_value = "some_value3"
  }

  resource_tags {
    resource_tag_key   = "some_tag4"
    resource_tag_value = "some_value4"
  }

  retention_period {
    retention_period_value = 10
    retention_period_unit  = "DAYS"
  }
}
`, description, resourceType)
}

func testAccRuleConfig_excludeResourceTags1(description, resourceType string) string {
	return fmt.Sprintf(`
resource "aws_rbin_rule" "test" {
  description   = %[1]q
  resource_type = %[2]q

  exclude_resource_tags {
    resource_tag_key   = "some_tag1"
    resource_tag_value = "some_value1"
  }

  retention_period {
    retention_period_value = 10
    retention_period_unit  = "DAYS"
  }
}
`, description, resourceType)
}

func testAccRuleConfig_excludeResourceTags2(description, resourceType string) string {
	return fmt.Sprintf(`
resource "aws_rbin_rule" "test" {
  description   = %[1]q
  resource_type = %[2]q

  exclude_resource_tags {
    resource_tag_key   = "some_tag3"
    resource_tag_value = "some_value3"
  }

  exclude_resource_tags {
    resource_tag_key   = "some_tag4"
    resource_tag_value = "some_value4"
  }

  retention_period {
    retention_period_value = 10
    retention_period_unit  = "DAYS"
  }
}
`, description, resourceType)
}

func testAccRuleConfig_lockConfig(resourceType, delay_unit1, delay_value1 string) string {
	return fmt.Sprintf(`
resource "aws_rbin_rule" "test" {
  resource_type = %[1]q

  retention_period {
    retention_period_value = 10
    retention_period_unit  = "DAYS"
  }

  lock_configuration {
    unlock_delay {
      unlock_delay_unit  = %[2]q
      unlock_delay_value = %[3]q
    }
  }
}
`, resourceType, delay_unit1, delay_value1)
}

func testAccRuleConfig_tags1(resourceType, tag1Key, tag1Value string) string {
	return fmt.Sprintf(`
resource "aws_rbin_rule" "test" {
  resource_type = %[1]q

  resource_tags {
    resource_tag_key   = "some_tag"
    resource_tag_value = ""
  }

  retention_period {
    retention_period_value = 10
    retention_period_unit  = "DAYS"
  }

  tags = {
    %[2]q = %[3]q
  }
}
`, resourceType, tag1Key, tag1Value)
}

func testAccRuleConfig_tags2(resourceType, tag1Key, tag1Value, tag2Key, tag2Value string) string {
	return fmt.Sprintf(`
resource "aws_rbin_rule" "test" {
  resource_type = %[1]q

  resource_tags {
    resource_tag_key   = "some_tag"
    resource_tag_value = ""
  }

  retention_period {
    retention_period_value = 10
    retention_period_unit  = "DAYS"
  }

  tags = {
    %[2]q = %[3]q
    %[4]q = %[5]q
  }
}
`, resourceType, tag1Key, tag1Value, tag2Key, tag2Value)
}
