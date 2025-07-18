// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssoadmin_test

import (
	"fmt"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccSSOAdminPrincipalApplicationAssignmentsDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_ssoadmin_principal_application_assignments.test"
	applicationResourceName := "aws_ssoadmin_application.test"
	userResourceName := "aws_identitystore_user.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.SSOAdminEndpointID)
			acctest.PreCheckSSOAdminInstances(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.SSOAdminServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccPrincipalApplicationAssignmentsDataSourceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "principal_id", userResourceName, "user_id"),
					resource.TestCheckResourceAttr(dataSourceName, "principal_type", "USER"),
					resource.TestCheckResourceAttr(dataSourceName, "application_assignments.#", "1"),
					resource.TestCheckResourceAttrPair(dataSourceName, "application_assignments.0.application_arn", applicationResourceName, names.AttrARN),
					resource.TestCheckResourceAttrPair(dataSourceName, "application_assignments.0.principal_id", userResourceName, "user_id"),
					resource.TestCheckResourceAttr(dataSourceName, "application_assignments.0.principal_type", "USER"),
				),
			},
		},
	})
}

func testAccPrincipalApplicationAssignmentsDataSourceConfigBase(rName string) string {
	return fmt.Sprintf(`
data "aws_ssoadmin_instances" "test" {}

resource "aws_ssoadmin_application" "test" {
  name                     = %[1]q
  application_provider_arn = %[2]q
  instance_arn             = tolist(data.aws_ssoadmin_instances.test.arns)[0]
}

resource "aws_identitystore_user" "test" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]

  display_name = "Acceptance Test"
  user_name    = %[1]q

  name {
    family_name = "Doe"
    given_name  = "John"
  }
}

resource "aws_ssoadmin_application_assignment" "test" {
  application_arn = aws_ssoadmin_application.test.arn
  principal_id    = aws_identitystore_user.test.user_id
  principal_type  = "USER"
}
`, rName, testAccApplicationProviderARN)
}

func testAccPrincipalApplicationAssignmentsDataSourceConfig_basic(rName string) string {
	return acctest.ConfigCompose(
		testAccPrincipalApplicationAssignmentsDataSourceConfigBase(rName),
		`
data "aws_ssoadmin_principal_application_assignments" "test" {
  depends_on = [aws_ssoadmin_application_assignment.test]

  instance_arn   = tolist(data.aws_ssoadmin_instances.test.arns)[0]
  principal_id   = aws_identitystore_user.test.user_id
  principal_type = "USER"
}
`)
}
