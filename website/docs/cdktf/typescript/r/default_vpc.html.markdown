---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "AWS: aws_default_vpc"
description: |-
  Manage a default VPC resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_default_vpc

Provides a resource to manage the [default AWS VPC](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/default-vpc.html)
in the current AWS Region.

If you created your AWS account after 2013-12-04 you have a default VPC in each AWS Region.

**This is an advanced resource** and has special caveats to be aware of when using it. Please read this document in its entirety before using this resource.

The `aws_default_vpc` resource behaves differently from normal resources in that if a default VPC exists, Terraform does not _create_ this resource, but instead "adopts" it into management.
If no default VPC exists, Terraform creates a new default VPC, which leads to the implicit creation of [other resources](https://docs.aws.amazon.com/vpc/latest/userguide/default-vpc.html#default-vpc-components).
By default, `terraform destroy` does not delete the default VPC but does remove the resource from Terraform state.
Set the `forceDestroy` argument to `true` to delete the default VPC.

## Example Usage

Basic usage with tags:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DefaultVpc } from "./.gen/providers/aws/default-vpc";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DefaultVpc(this, "default", {
      tags: {
        Name: "Default VPC",
      },
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

The arguments of an `aws_default_vpc` differ slightly from those of [`aws_vpc`](vpc.html):

* The `cidrBlock` and `instanceTenancy` arguments become computed attributes
* The default value for `enableDnsHostnames` is `true`

This resource supports the following additional arguments:

* `forceDestroy` - (Optional) Whether destroying the resource deletes the default VPC. Default: `false`

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `cidrBlock` - The primary IPv4 CIDR block for the VPC
* `instanceTenancy` - The allowed tenancy of instances launched into the VPC

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Default VPCs using the VPC `id`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DefaultVpc } from "./.gen/providers/aws/default-vpc";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    DefaultVpc.generateConfigForImport(this, "default", "vpc-a01106c2");
  }
}

```

Using `terraform import`, import Default VPCs using the VPC `id`. For example:

```console
% terraform import aws_default_vpc.default vpc-a01106c2
```

<!-- cache-key: cdktf-0.20.8 input-4835b9bc17ad186a1703d08750efb5d17642302c2b2608733af35573882eea29 -->