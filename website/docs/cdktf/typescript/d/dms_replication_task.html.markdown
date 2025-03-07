---
subcategory: "DMS (Database Migration)"
layout: "aws"
page_title: "AWS: aws_dms_replication_task"
description: |-
  Terraform data source for managing an AWS DMS (Database Migration) Replication Task.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_dms_replication_task

Terraform data source for managing an AWS DMS (Database Migration) Replication Task.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsDmsReplicationTask } from "./.gen/providers/aws/data-aws-dms-replication-task";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsDmsReplicationTask(this, "test", {
      replicationTaskId: Token.asString(
        awsDmsReplicationTaskTest.replicationTaskId
      ),
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `replicationTaskId` - (Required) The replication task identifier.

    - Must contain from 1 to 255 alphanumeric characters or hyphens.
    - First character must be a letter.
    - Cannot end with a hyphen.
    - Cannot contain two consecutive hyphens.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `cdcStartPosition` - (Conflicts with `cdcStartTime`) Indicates when you want a change data capture (CDC) operation to start. The value can be in date, checkpoint, or LSN/SCN format depending on the source engine. For more information, see [Determining a CDC native start point](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Task.CDC.html#CHAP_Task.CDC.StartPoint.Native).
* `cdcStartTime` - (Conflicts with `cdcStartPosition`) The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
* `migrationType` - The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
* `replicationInstanceArn` - The Amazon Resource Name (ARN) of the replication instance.
* `replicationTaskSettings` - An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
* `sourceEndpointArn` - The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
* `startReplicationTask` -  Whether to run or stop the replication task.
* `status` - Replication Task status.
* `tableMappings` - An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
* `targetEndpointArn` - The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
* `replicationTaskArn` - The Amazon Resource Name (ARN) for the replication task.

<!-- cache-key: cdktf-0.20.8 input-fa1a1915aa39f86d68f0c890bf2884e184e4111a8bb8bbe6b38fd7c6c992b063 -->