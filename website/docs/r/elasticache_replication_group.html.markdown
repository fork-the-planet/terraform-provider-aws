---
subcategory: "ElastiCache"
layout: "aws"
page_title: "AWS: aws_elasticache_replication_group"
description: |-
  Provides an ElastiCache Replication Group resource.
---

# Resource: aws_elasticache_replication_group

Provides an ElastiCache Replication Group resource.

For working with a [Memcached cluster](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/WhatIs.html) or a
[single-node Redis instance (Cluster Mode Disabled)](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/WhatIs.html),
see the [`aws_elasticache_cluster` resource](/docs/providers/aws/r/elasticache_cluster.html).

~> **Note:** When you change an attribute, such as `engine_version`, by
default the ElastiCache API applies it in the next maintenance window. Because
of this, Terraform may report a difference in its planning phase because the
actual modification has not yet taken place. You can use the
`apply_immediately` flag to instruct the service to apply the change
immediately. Using `apply_immediately` can result in a brief downtime as
servers reboots.
See the AWS Documentation on
[Modifying an ElastiCache Cache Cluster](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.Modify.html)
for more information.

~> **Note:** Any attribute changes that re-create the resource will be applied immediately, regardless of the value of `apply_immediately`.

~> **Note:** Be aware of the terminology collision around "cluster" for `aws_elasticache_replication_group`. For example, it is possible to create a ["Cluster Mode Disabled [Redis] Cluster"](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.Create.CON.Redis.html). With "Cluster Mode Enabled", the data will be stored in shards (called "node groups"). See [Redis Cluster Configuration](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/cluster-create-determine-requirements.html#redis-cluster-configuration) for a diagram of the differences. To enable cluster mode, use a parameter group that has cluster mode enabled. The default parameter groups provided by AWS end with ".cluster.on", for example `default.redis6.x.cluster.on`.

## Example Usage

### Redis OSS/Valkey Cluster Mode Disabled

To create a single shard primary with single read replica:

```terraform
resource "aws_elasticache_replication_group" "example" {
  automatic_failover_enabled  = true
  preferred_cache_cluster_azs = ["us-west-2a", "us-west-2b"]
  replication_group_id        = "tf-rep-group-1"
  description                 = "example description"
  node_type                   = "cache.m4.large"
  num_cache_clusters          = 2
  parameter_group_name        = "default.redis3.2"
  port                        = 6379
}
```

You have two options for adjusting the number of replicas:

* Adjusting `num_cache_clusters` directly. This will attempt to automatically add or remove replicas, but provides no granular control (e.g., preferred availability zone, cache cluster ID) for the added or removed replicas. This also currently expects cache cluster IDs in the form of `replication_group_id-00#`.
* Otherwise for fine grained control of the underlying cache clusters, they can be added or removed with the [`aws_elasticache_cluster` resource](/docs/providers/aws/r/elasticache_cluster.html) and its `replication_group_id` attribute. In this situation, you will need to utilize the [lifecycle configuration block](https://www.terraform.io/docs/configuration/meta-arguments/lifecycle.html) with `ignore_changes` to prevent perpetual differences during Terraform plan with the `num_cache_cluster` attribute.

```terraform
resource "aws_elasticache_replication_group" "example" {
  automatic_failover_enabled  = true
  preferred_cache_cluster_azs = ["us-west-2a", "us-west-2b"]
  replication_group_id        = "tf-rep-group-1"
  description                 = "example description"
  node_type                   = "cache.m4.large"
  num_cache_clusters          = 2
  parameter_group_name        = "default.redis3.2"
  port                        = 6379

  lifecycle {
    ignore_changes = [num_cache_clusters]
  }
}

resource "aws_elasticache_cluster" "replica" {
  count = 1

  cluster_id           = "tf-rep-group-1-${count.index}"
  replication_group_id = aws_elasticache_replication_group.example.id
}
```

### Redis OSS/Valkey Cluster Mode Enabled

To create two shards with a primary and a single read replica each:

```terraform
resource "aws_elasticache_replication_group" "baz" {
  replication_group_id       = "tf-redis-cluster"
  description                = "example description"
  node_type                  = "cache.t2.small"
  port                       = 6379
  parameter_group_name       = "default.redis3.2.cluster.on"
  automatic_failover_enabled = true

  num_node_groups         = 2
  replicas_per_node_group = 1
}
```

### Redis Log Delivery configuration

```terraform
resource "aws_elasticache_replication_group" "test" {
  replication_group_id       = "myreplicaciongroup"
  description                = "test description"
  node_type                  = "cache.t3.small"
  port                       = 6379
  apply_immediately          = true
  auto_minor_version_upgrade = false
  maintenance_window         = "tue:06:30-tue:07:30"
  snapshot_window            = "01:00-02:00"
  log_delivery_configuration {
    destination      = aws_cloudwatch_log_group.example.name
    destination_type = "cloudwatch-logs"
    log_format       = "text"
    log_type         = "slow-log"
  }
  log_delivery_configuration {
    destination      = aws_kinesis_firehose_delivery_stream.example.name
    destination_type = "kinesis-firehose"
    log_format       = "json"
    log_type         = "engine-log"
  }
}
```

~> **Note:** We currently do not support passing a `primary_cluster_id` in order to create the Replication Group.

~> **Note:** Automatic Failover is unavailable for Redis versions earlier than 2.8.6,
and unavailable on T1 node types. For T2 node types, it is only available on Redis version 3.2.4 or later with cluster mode enabled. See the [High Availability Using Replication Groups](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Replication.html) guide
for full details on using Replication Groups.

### Creating a secondary replication group for a global replication group

A Global Replication Group can have up to two secondary Replication Groups in different regions. These are added to an existing Global Replication Group.

```terraform
resource "aws_elasticache_replication_group" "secondary" {
  replication_group_id        = "example-secondary"
  description                 = "secondary replication group"
  global_replication_group_id = aws_elasticache_global_replication_group.example.global_replication_group_id

  num_cache_clusters = 1
}

resource "aws_elasticache_global_replication_group" "example" {
  provider = aws.other_region

  global_replication_group_id_suffix = "example"
  primary_replication_group_id       = aws_elasticache_replication_group.primary.id
}

resource "aws_elasticache_replication_group" "primary" {
  provider = aws.other_region

  replication_group_id = "example-primary"
  description          = "primary replication group"

  engine         = "redis"
  engine_version = "5.0.6"
  node_type      = "cache.m5.large"

  num_cache_clusters = 1
}
```

### Redis AUTH and In-Transit Encryption Enabled

```terraform
resource "aws_elasticache_replication_group" "example" {
  replication_group_id = "example"
  description          = "example with authentication"
  node_type            = "cache.t2.micro"
  num_cache_clusters   = 1
  port                 = 6379
  subnet_group_name    = aws_elasticache_subnet_group.example.name
  security_group_ids   = [aws_security_group.example.id]
  parameter_group_name = "default.redis5.0"
  engine_version       = "5.0.6"

  transit_encryption_enabled = true
  auth_token                 = "abcdefgh1234567890"
  auth_token_update_strategy = "ROTATE"
}
```

~> When adding a new `auth_token` to a previously passwordless replication group, using the `ROTATE` update strategy will result in support for **both** the new token and passwordless authentication. To immediately require authorization when adding the initial token, use the `SET` strategy instead. See the [Authenticating with the Redis AUTH command](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/auth.html) guide for additional details.

## Argument Reference

The following arguments are required:

* `description` - (Required) User-created description for the replication group. Must not be empty.
* `replication_group_id` - (Required) Replication group identifier. This parameter is stored as a lowercase string.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `apply_immediately` - (Optional) Specifies whether any modifications are applied immediately, or during the next maintenance window. Default is `false`.
* `at_rest_encryption_enabled` - (Optional) Whether to enable encryption at rest.
  When `engine` is `redis`, default is `false`.
  When `engine` is `valkey`, default is `true`.
* `auth_token` - (Optional) Password used to access a password protected server. Can be specified only if `transit_encryption_enabled = true`.
* `auth_token_update_strategy` - (Optional) Strategy to use when updating the `auth_token`. Valid values are `SET`, `ROTATE`, and `DELETE`. Required if `auth_token` is set.
* `auto_minor_version_upgrade` - (Optional) Specifies whether minor version engine upgrades will be applied automatically to the underlying Cache Cluster instances during the maintenance window.
  Only supported for engine types `"redis"` and `"valkey"` and if the engine version is 6 or higher.
  Defaults to `true`.
* `automatic_failover_enabled` - (Optional) Specifies whether a read-only replica will be automatically promoted to read/write primary if the existing primary fails. If enabled, `num_cache_clusters` must be greater than 1. Must be enabled for Redis (cluster mode enabled) replication groups. Defaults to `false`.
* `cluster_mode` - (Optional) Specifies whether cluster mode is enabled or disabled. Valid values are `enabled` or `disabled` or `compatible`
* `data_tiering_enabled` - (Optional) Enables data tiering. Data tiering is only supported for replication groups using the r6gd node type. This parameter must be set to `true` when using r6gd nodes.
* `engine` - (Optional) Name of the cache engine to be used for the clusters in this replication group.
  Valid values are `redis` or `valkey`.
  Default is `redis`.
* `engine_version` - (Optional) Version number of the cache engine to be used for the cache clusters in this replication group.
  If the version is 7 or higher, the major and minor version should be set, e.g., `7.2`.
  If the version is 6, the major and minor version can be set, e.g., `6.2`,
  or the minor version can be unspecified which will use the latest version at creation time, e.g., `6.x`.
  Otherwise, specify the full version desired, e.g., `5.0.6`.
  The actual engine version used is returned in the attribute `engine_version_actual`, see [Attribute Reference](#attribute-reference) below.
* `final_snapshot_identifier` - (Optional) The name of your final node group (shard) snapshot. ElastiCache creates the snapshot from the primary node in the cluster. If omitted, no final snapshot will be made.
* `global_replication_group_id` - (Optional) The ID of the global replication group to which this replication group should belong. If this parameter is specified, the replication group is added to the specified global replication group as a secondary replication group; otherwise, the replication group is not part of any global replication group. If `global_replication_group_id` is set, the `num_node_groups` parameter cannot be set.
* `ip_discovery` - (Optional) The IP version to advertise in the discovery protocol. Valid values are `ipv4` or `ipv6`.
* `kms_key_id` - (Optional) The ARN of the key that you wish to use if encrypting at rest. If not supplied, uses service managed encryption. Can be specified only if `at_rest_encryption_enabled = true`.
* `log_delivery_configuration` - (Optional, Redis only) Specifies the destination and format of Redis OSS/Valkey [SLOWLOG](https://redis.io/commands/slowlog) or Redis OSS/Valkey [Engine Log](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Log_Delivery.html#Log_contents-engine-log). See the documentation on [Amazon ElastiCache](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Log_Delivery.html#Log_contents-engine-log). See [Log Delivery Configuration](#log-delivery-configuration) below for more details.
* `maintenance_window` - (Optional) Specifies the weekly time range for when maintenance on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
* `multi_az_enabled` - (Optional) Specifies whether to enable Multi-AZ Support for the replication group.
  If `true`, `automatic_failover_enabled` must also be enabled.
  Defaults to `false`.
* `network_type` - (Optional) The IP versions for cache cluster connections. Valid values are `ipv4`, `ipv6` or `dual_stack`.
* `node_type` - (Optional) Instance class to be used.
  See AWS documentation for information on [supported node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html) and [guidance on selecting node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/nodes-select-size.html).
  Required unless `global_replication_group_id` is set.
  Cannot be set if `global_replication_group_id` is set.
* `notification_topic_arn` - (Optional) ARN of an SNS topic to send ElastiCache notifications to. Example: `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
* `num_cache_clusters` - (Optional) Number of cache clusters (primary and replicas) this replication group will have.
  If `automatic_failover_enabled` or `multi_az_enabled` are `true`, must be at least 2.
  Updates will occur before other modifications.
  Conflicts with `num_node_groups` and `replicas_per_node_group`.
  Defaults to `1`.
* `num_node_groups` - (Optional) Number of node groups (shards) for this Redis replication group.
  Changing this number will trigger a resizing operation before other settings modifications.
  Conflicts with `num_cache_clusters`.
* `parameter_group_name` - (Optional) Name of the parameter group to associate with this replication group. If this argument is omitted, the default cache parameter group for the specified engine is used. To enable "cluster mode", i.e., data sharding, use a parameter group that has the parameter `cluster-enabled` set to true.
* `port` - (Optional) Port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379.
* `preferred_cache_cluster_azs` - (Optional) List of EC2 availability zones in which the replication group's cache clusters will be created. The order of the availability zones in the list is considered. The first item in the list will be the primary node. Ignored when updating.
* `replicas_per_node_group` - (Optional) Number of replica nodes in each node group.
  Changing this number will trigger a resizing operation before other settings modifications.
  Valid values are 0 to 5.
  Conflicts with `num_cache_clusters`.
  Can only be set if `num_node_groups` is set.
* `security_group_ids` - (Optional) IDs of one or more Amazon VPC security groups associated with this replication group. Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud.
* `security_group_names` - (Optional) Names of one or more Amazon VPC security groups associated with this replication group. Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud.
* `snapshot_arns` - (Optional) List of ARNs that identify Redis RDB snapshot files stored in Amazon S3. The names object names cannot contain any commas.
* `snapshot_name` - (Optional) Name of a snapshot from which to restore data into the new node group. Changing the `snapshot_name` forces a new resource.
* `snapshot_retention_limit` - (Optional, Redis only) Number of days for which ElastiCache will retain automatic cache cluster snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days before being deleted. If the value of `snapshot_retention_limit` is set to zero (0), backups are turned off. Please note that setting a `snapshot_retention_limit` is not supported on cache.t1.micro cache nodes
* `snapshot_window` - (Optional, Redis only) Daily time range (in UTC) during which ElastiCache will begin taking a daily snapshot of your cache cluster. The minimum snapshot window is a 60 minute period. Example: `05:00-09:00`
* `subnet_group_name` - (Optional) Name of the cache subnet group to be used for the replication group.
* `tags` - (Optional) Map of tags to assign to the resource. Adding tags to this resource will add or overwrite any existing tags on the clusters in the replication group and not to the group itself. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `transit_encryption_enabled` - (Optional) Whether to enable encryption in transit.
  Changing this argument with an `engine_version` < `7.0.5` will force a replacement.
  Engine versions prior to `7.0.5` only allow this transit encryption to be configured during creation of the replication group.
* `transit_encryption_mode` - (Optional) A setting that enables clients to migrate to in-transit encryption with no downtime.
  Valid values are `preferred` and `required`.
  When enabling encryption on an existing replication group, this must first be set to `preferred` before setting it to `required` in a subsequent apply.
  See the `TransitEncryptionMode` field in the [`CreateReplicationGroup` API documentation](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CreateReplicationGroup.html) for additional details.
* `user_group_ids` - (Optional) User Group ID to associate with the replication group. Only a maximum of one (1) user group ID is valid. **NOTE:** This argument _is_ a set because the AWS specification allows for multiple IDs. However, in practice, AWS only allows a maximum size of one.

### Log Delivery Configuration

The `log_delivery_configuration` block allows the streaming of Redis OSS/Valkey [SLOWLOG](https://redis.io/commands/slowlog) or Redis OSS/Valkey [Engine Log](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Log_Delivery.html#Log_contents-engine-log) to CloudWatch Logs or Kinesis Data Firehose. Max of 2 blocks.

* `destination` - Name of either the CloudWatch Logs LogGroup or Kinesis Data Firehose resource.
* `destination_type` - For CloudWatch Logs use `cloudwatch-logs` or for Kinesis Data Firehose use `kinesis-firehose`.
* `log_format` - Valid values are `json` or `text`
* `log_type` - Valid values are  `slow-log` or `engine-log`. Max 1 of each.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the created ElastiCache Replication Group.
* `engine_version_actual` - Because ElastiCache pulls the latest minor or patch for a version, this attribute returns the running version of the cache engine.
* `cluster_enabled` - Indicates if cluster mode is enabled.
* `configuration_endpoint_address` - Address of the replication group configuration endpoint when cluster mode is enabled.
* `id` - ID of the ElastiCache Replication Group.
* `member_clusters` - Identifiers of all the nodes that are part of this replication group.
* `primary_endpoint_address` - (Redis only) Address of the endpoint for the primary node in the replication group, if the cluster mode is disabled.
* `reader_endpoint_address` - (Redis only) Address of the endpoint for the reader node in the replication group, if the cluster mode is disabled.
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `60m`)
* `delete` - (Default `45m`)
* `update` - (Default `40m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import ElastiCache Replication Groups using the `replication_group_id`. For example:

```terraform
import {
  to = aws_elasticache_replication_group.my_replication_group
  id = "replication-group-1"
}
```

Using `terraform import`, import ElastiCache Replication Groups using the `replication_group_id`. For example:

```console
% terraform import aws_elasticache_replication_group.my_replication_group replication-group-1
```
