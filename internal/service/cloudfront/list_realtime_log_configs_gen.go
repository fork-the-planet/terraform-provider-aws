// Code generated by "internal/generate/listpages/main.go -ListOps=ListRealtimeLogConfigs -InputPaginator=Marker -OutputPaginator=RealtimeLogConfigs.NextMarker -- list_realtime_log_configs_gen.go"; DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

func listRealtimeLogConfigsPages(ctx context.Context, conn *cloudfront.Client, input *cloudfront.ListRealtimeLogConfigsInput, fn func(*cloudfront.ListRealtimeLogConfigsOutput, bool) bool, optFns ...func(*cloudfront.Options)) error {
	for {
		output, err := conn.ListRealtimeLogConfigs(ctx, input, optFns...)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.RealtimeLogConfigs.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.Marker = output.RealtimeLogConfigs.NextMarker
	}
	return nil
}
