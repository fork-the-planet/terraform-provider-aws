---
subcategory: "Lightsail"
layout: "aws"
page_title: "AWS: aws_lightsail_static_ip"
description: |-
  Manages a Lightsail Static IP.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_lightsail_static_ip

Manages a static IP address.

Use this resource to allocate a static IP address that can be attached to Lightsail instances to provide a consistent public IP address that persists across instance restarts.

~> **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.lightsail_static_ip import LightsailStaticIp
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        LightsailStaticIp(self, "example",
            name="example"
        )
```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name for the allocated static IP.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the Lightsail static IP.
* `ip_address` - Allocated static IP address.
* `support_code` - Support code for the static IP. Include this code in your email to support when you have questions about a static IP in Lightsail. This code enables our support team to look up your Lightsail information more easily.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_lightsail_static_ip` using the name attribute. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.lightsail_static_ip import LightsailStaticIp
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        LightsailStaticIp.generate_config_for_import(self, "example", "example")
```

Using `terraform import`, import `aws_lightsail_static_ip` using the name attribute. For example:

```console
% terraform import aws_lightsail_static_ip.example example
```

<!-- cache-key: cdktf-0.20.8 input-dd772c37153df1e64a96b6b30283b7d2a377dcd590087c123ee4380c2126dc2c -->