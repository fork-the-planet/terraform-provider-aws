resource "aws_appflow_flow" "test" {
{{- template "region" }}
  name = var.rName

  source_flow_config {
    connector_type = "S3"
    source_connector_properties {
      s3 {
        bucket_name   = aws_s3_bucket_policy.test_source.bucket
        bucket_prefix = "flow"
      }
    }
  }

  destination_flow_config {
    connector_type = "S3"
    destination_connector_properties {
      s3 {
        bucket_name = aws_s3_bucket_policy.test_destination.bucket

        s3_output_format_config {
          prefix_config {
            prefix_type = "PATH"
          }
        }
      }
    }
  }

  task {
    source_fields     = ["testField"]
    destination_field = "testField"
    task_type         = "Map"

    connector_operator {
      s3 = "NO_OP"
    }
  }

  trigger_config {
    trigger_type = "OnDemand"
  }

{{- template "tags" . }}
}

data "aws_partition" "current" {}

resource "aws_s3_bucket" "test_source" {
{{- template "region" }}
  bucket = "${var.rName}-source"
}

resource "aws_s3_bucket_policy" "test_source" {
{{- template "region" }}
  bucket = aws_s3_bucket.test_source.bucket
  policy = <<EOF
{
    "Statement": [
        {
            "Effect": "Allow",
            "Sid": "AllowAppFlowSourceActions",
            "Principal": {
                "Service": "appflow.amazonaws.com"
            },
            "Action": [
                "s3:ListBucket",
                "s3:GetObject"
            ],
            "Resource": [
                "arn:${data.aws_partition.current.partition}:s3:::${aws_s3_bucket.test_source.bucket}",
                "arn:${data.aws_partition.current.partition}:s3:::${aws_s3_bucket.test_source.bucket}/*"
            ]
        }
    ],
	"Version": "2012-10-17"
}
EOF
}

resource "aws_s3_object" "test" {
{{- template "region" }}
  bucket = aws_s3_bucket.test_source.bucket
  key    = "flow_source.csv"
  source = "test-fixtures/flow_source.csv"
}

resource "aws_s3_bucket" "test_destination" {
{{- template "region" }}
  bucket = "${var.rName}-destination"
}

resource "aws_s3_bucket_policy" "test_destination" {
{{- template "region" }}
  bucket = aws_s3_bucket.test_destination.bucket
  policy = <<EOF

{
    "Statement": [
        {
            "Effect": "Allow",
            "Sid": "AllowAppFlowDestinationActions",
            "Principal": {
                "Service": "appflow.amazonaws.com"
            },
            "Action": [
                "s3:PutObject",
                "s3:AbortMultipartUpload",
                "s3:ListMultipartUploadParts",
                "s3:ListBucketMultipartUploads",
                "s3:GetBucketAcl",
                "s3:PutObjectAcl"
            ],
            "Resource": [
                "arn:${data.aws_partition.current.partition}:s3:::${aws_s3_bucket.test_destination.bucket}",
                "arn:${data.aws_partition.current.partition}:s3:::${aws_s3_bucket.test_destination.bucket}/*"
            ]
        }
    ],
	"Version": "2012-10-17"
}
EOF
}
