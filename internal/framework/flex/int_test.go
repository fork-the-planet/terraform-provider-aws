// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package flex_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
)

func TestInt64FromFramework(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input    types.Int64
		expected *int64
	}
	tests := map[string]testCase{
		"valid int64": {
			input:    types.Int64Value(42),
			expected: aws.Int64(42),
		},
		"zero int64": {
			input:    types.Int64Value(0),
			expected: aws.Int64(0),
		},
		"null int64": {
			input:    types.Int64Null(),
			expected: nil,
		},
		"unknown int64": {
			input:    types.Int64Unknown(),
			expected: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := flex.Int64FromFramework(context.Background(), test.input)

			if diff := cmp.Diff(got, test.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestInt64ToFramework(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input    *int64
		expected types.Int64
	}
	tests := map[string]testCase{
		"valid int64": {
			input:    aws.Int64(42),
			expected: types.Int64Value(42),
		},
		"zero int64": {
			input:    aws.Int64(0),
			expected: types.Int64Value(0),
		},
		"nil int64": {
			input:    nil,
			expected: types.Int64Null(),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := flex.Int64ToFramework(context.Background(), test.input)

			if diff := cmp.Diff(got, test.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestInt64ToFrameworkLegacy(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input    *int64
		expected types.Int64
	}
	tests := map[string]testCase{
		"valid int64": {
			input:    aws.Int64(42),
			expected: types.Int64Value(42),
		},
		"zero int64": {
			input:    aws.Int64(0),
			expected: types.Int64Value(0),
		},
		"nil int64": {
			input:    nil,
			expected: types.Int64Value(0),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := flex.Int64ToFrameworkLegacy(context.Background(), test.input)

			if diff := cmp.Diff(got, test.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestInt32ToFrameworkInt64(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input    *int32
		expected types.Int64
	}
	tests := map[string]testCase{
		"valid int64": {
			input:    aws.Int32(42),
			expected: types.Int64Value(42),
		},
		"zero int64": {
			input:    aws.Int32(0),
			expected: types.Int64Value(0),
		},
		"nil int64": {
			input:    nil,
			expected: types.Int64Null(),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := flex.Int32ToFrameworkInt64(context.Background(), test.input)

			if diff := cmp.Diff(got, test.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestInt32ToFrameworkInt64Legacy(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input    *int32
		expected types.Int64
	}
	tests := map[string]testCase{
		"valid int64": {
			input:    aws.Int32(42),
			expected: types.Int64Value(42),
		},
		"zero int64": {
			input:    aws.Int32(0),
			expected: types.Int64Value(0),
		},
		"nil int64": {
			input:    nil,
			expected: types.Int64Value(0),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := flex.Int32ToFrameworkInt64Legacy(context.Background(), test.input)

			if diff := cmp.Diff(got, test.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestInt32FromFrameworkInt64(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input    types.Int64
		expected *int32
	}
	tests := map[string]testCase{
		"valid int64": {
			input:    types.Int64Value(42),
			expected: aws.Int32(42),
		},
		"zero int64": {
			input:    types.Int64Value(0),
			expected: aws.Int32(0),
		},
		"null int64": {
			input:    types.Int64Null(),
			expected: nil,
		},
		"unknown int64": {
			input:    types.Int64Unknown(),
			expected: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := flex.Int32FromFrameworkInt64(context.Background(), test.input)

			if diff := cmp.Diff(got, test.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestInt32FromFrameworkLegacy(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input    types.Int32
		expected *int32
	}
	tests := map[string]testCase{
		"valid int32": {
			input:    types.Int32Value(42),
			expected: aws.Int32(42),
		},
		"zero int32": {
			input:    types.Int32Value(0),
			expected: nil,
		},
		"null int32": {
			input:    types.Int32Null(),
			expected: nil,
		},
		"unknown int32": {
			input:    types.Int32Unknown(),
			expected: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := flex.Int32FromFrameworkLegacy(context.Background(), test.input)

			if diff := cmp.Diff(got, test.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}
