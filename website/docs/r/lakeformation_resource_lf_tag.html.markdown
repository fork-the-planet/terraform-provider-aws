---
subcategory: "Lake Formation"
layout: "aws"
page_title: "AWS: aws_lakeformation_resource_lf_tag"
description: |-
  Terraform resource for managing an AWS Lake Formation Resource LF Tag.
---
# Resource: aws_lakeformation_resource_lf_tag

Terraform resource for managing an AWS Lake Formation Resource LF Tag.

## Example Usage

### Basic Usage

```terraform
resource "aws_lakeformation_resource_lf_tag" "example" {
  database {
    name = aws_glue_catalog_database.example.name
  }

  lf_tag {
    key   = aws_lakeformation_lf_tag.example.key
    value = "stowe"
  }
}
```

## Argument Reference

The following arguments are required:

* `lf_tag` - (Required) Set of LF-tags to attach to the resource. See [LF Tag](#lf-tag) for more details.

Exactly one of the following is required:

* `database` - (Optional) Configuration block for a database resource. See [Database](#database) for more details.
* `table` - (Optional) Configuration block for a table resource. See [Table](#table) for more details.
* `table_with_columns` - (Optional) Configuration block for a table with columns resource. See [Table With Columns](#table-with-columns) for more details.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `catalog_id` - (Optional) Identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.

### LF Tag

The following arguments are required:

* `key` - (Required) Key name for an existing LF-tag.
* `value` - (Required) Value from the possible values for the LF-tag.

The following argument is optional:

* `catalog_id` - (Optional) Identifier for the Data Catalog. By default, it is the account ID of the caller.

### Database

The following argument is required:

* `name` - (Required) Name of the database resource. Unique to the Data Catalog.

The following argument is optional:

* `catalog_id` - (Optional) Identifier for the Data Catalog. By default, it is the account ID of the caller.

### Table

The following argument is required:

* `database_name` - (Required) Name of the database for the table. Unique to a Data Catalog.
* `name` - (Required, at least one of `name` or `wildcard`) Name of the table.
* `wildcard` - (Required, at least one of `name` or `wildcard`) Whether to use a wildcard representing every table under a database. Defaults to `false`.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `catalog_id` - (Optional) Identifier for the Data Catalog. By default, it is the account ID of the caller.

### Table With Columns

The following arguments are required:

* `column_names` - (Required, at least one of `column_names` or `wildcard`) Set of column names for the table.
* `database_name` - (Required) Name of the database for the table with columns resource. Unique to the Data Catalog.
* `name` - (Required) Name of the table resource.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `catalog_id` - (Optional) Identifier for the Data Catalog. By default, it is the account ID of the caller.
* `column_wildcard` - (Optional) Option to add column wildcard. See [Column Wildcard](#column-wildcard) for more details.

### Column Wildcard

* `excluded_column_names` - (Optional) Set of column names for the table to exclude. If `excluded_column_names` is included, `wildcard` must be set to `true` to avoid Terraform reporting a difference.

## Attribute Reference

This resource exports no additional attributes.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `20m`)
* `delete` - (Default `20m`)

## Import

You cannot import this resource.
