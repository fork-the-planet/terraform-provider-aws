// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validators

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	inttypes "github.com/hashicorp/terraform-provider-aws/internal/types"
)

// awsRegionValidator validates that a string Attribute's value is a valid AWS Region.
type awsRegionValidator struct{}

// Description describes the validation in plain text formatting.
func (validator awsRegionValidator) Description(_ context.Context) string {
	return "value must be a valid AWS Region Code"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator awsRegionValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// ValidateString performs the validation.
func (validator awsRegionValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	if !inttypes.IsAWSRegion(request.ConfigValue.ValueString()) {
		response.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			request.Path,
			validator.Description(ctx),
			request.ConfigValue.ValueString(),
		))
		return
	}
}

// AWSRegion returns a string validator which ensures that any configured
// attribute value:
//
//   - Is a string, which represents a valid AWS Region.
//
// Null (unconfigured) and unknown (known after apply) values are skipped.
func AWSRegion() validator.String { // nosemgrep:ci.aws-in-func-name
	return awsRegionValidator{}
}
