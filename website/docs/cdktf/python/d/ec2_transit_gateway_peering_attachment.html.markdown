---
subcategory: "Transit Gateway"
layout: "aws"
page_title: "AWS: aws_ec2_transit_gateway_peering_attachment"
description: |-
  Get information on an EC2 Transit Gateway Peering Attachment
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_ec2_transit_gateway_peering_attachment

Get information on an EC2 Transit Gateway Peering Attachment.

## Example Usage

### By Filter

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_ec2_transit_gateway_peering_attachment import DataAwsEc2TransitGatewayPeeringAttachment
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsEc2TransitGatewayPeeringAttachment(self, "example",
            filter=[DataAwsEc2TransitGatewayPeeringAttachmentFilter(
                name="transit-gateway-attachment-id",
                values=["tgw-attach-12345678"]
            )
            ]
        )
```

### By Identifier

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_ec2_transit_gateway_peering_attachment import DataAwsEc2TransitGatewayPeeringAttachment
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsEc2TransitGatewayPeeringAttachment(self, "attachment",
            id="tgw-attach-12345678"
        )
```

## Argument Reference

This data source supports the following arguments:

* `filter` - (Optional) One or more configuration blocks containing name-values filters. Detailed below.
* `id` - (Optional) Identifier of the EC2 Transit Gateway Peering Attachment.
* `tags` - (Optional) Mapping of tags, each pair of which must exactly match
  a pair on the specific EC2 Transit Gateway Peering Attachment to retrieve.

More complex filters can be expressed using one or more `filter` sub-blocks,
which take the following arguments:

* `name` - (Required) Name of the field to filter by, as defined by
  [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayPeeringAttachments.html).
* `values` - (Required) Set of values that are accepted for the given field.
  An EC2 Transit Gateway Peering Attachment be selected if any one of the given values matches.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN of the attachment.
* `peer_account_id` - Identifier of the peer AWS account.
* `peer_region` - Identifier of the peer AWS region.
* `peer_transit_gateway_id` - Identifier of the peer EC2 Transit Gateway.
* `transit_gateway_id` - Identifier of the local EC2 Transit Gateway.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `read` - (Default `20m`)

<!-- cache-key: cdktf-0.20.8 input-be0c89eb348f62d2a6a8eb5de8f29f1f725c750177016bb80a7bda1858be2ead -->