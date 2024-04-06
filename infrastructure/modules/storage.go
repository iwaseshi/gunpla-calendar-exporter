package modules

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-google-go/google/v6/storagebucket"
	"github.com/cdktf/cdktf-provider-google-go/google/v6/storagebucketiampolicy"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewStorageBucket(stack cdktf.TerraformStack) {
	storage := storagebucket.NewStorageBucket(stack, jsii.String("gcs_bucket"), &storagebucket.StorageBucketConfig{
		Location: jsii.String("asia-northeast1"),
		Name:     jsii.String("gunpla-calendar-exporter"),
	})

	storagebucketiampolicy.NewStorageBucketIamPolicy(stack, jsii.String("bucket_iam"), &storagebucketiampolicy.StorageBucketIamPolicyConfig{
		Bucket: storage.Name(),
		PolicyData: jsii.String(`{
			"bindings": [
				{
					"role":"roles/storage.legacyObjectReader",
					"members": [
						"allUsers"
					]
				}
			]
		}`),
	})

}
