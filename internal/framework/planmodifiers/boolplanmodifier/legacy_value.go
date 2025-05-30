// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package boolplanmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// LegacyValue returns a plan modifier that prevents `known after apply` during creation plans for
// attributes that must be `Computed,Optional` for legacy value reasons.
func LegacyValue() planmodifier.Bool {
	return legacyValueModifier{}
}

type legacyValueModifier struct{}

func (m legacyValueModifier) Description(_ context.Context) string {
	return ""
}

func (m legacyValueModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m legacyValueModifier) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	// Exit if another planmodifier has set the value
	if !req.PlanValue.IsUnknown() {
		return
	}

	// Use value from Config if set
	if !req.ConfigValue.IsNull() {
		return
	}

	// Do nothing if there is an unknown configuration value, otherwise interpolation gets messed up.
	if req.ConfigValue.IsUnknown() {
		return
	}

	resp.PlanValue = types.BoolValue(false)
}
