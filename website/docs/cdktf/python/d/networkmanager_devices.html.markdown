---
subcategory: "Network Manager"
layout: "aws"
page_title: "AWS: aws_networkmanager_devices"
description: |-
  Retrieve information about devices.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_networkmanager_devices

Retrieve information about devices.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_networkmanager_devices import DataAwsNetworkmanagerDevices
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsNetworkmanagerDevices(self, "example",
            global_network_id=global_network_id.string_value,
            tags={
                "Env": "test"
            }
        )
```

## Argument Reference

This data source supports the following arguments:

* `global_network_id` - (Required) ID of the Global Network of the devices to retrieve.
* `site_id` - (Optional) ID of the site of the devices to retrieve.
* `tags` - (Optional) Restricts the list to the devices with these tags.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `ids` - IDs of the devices.

<!-- cache-key: cdktf-0.20.8 input-644fce3d16db9a6450adf13f869f38e3cb841b67d89b67b5220257321dba7e4a -->