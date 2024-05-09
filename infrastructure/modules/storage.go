package modules

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/serviceaccount"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/storagebucket"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/storagebucketiampolicy"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewStorageBucket(stack cdktf.TerraformStack) {
	bucket := storagebucket.NewStorageBucket(stack, jsii.String("gcs_bucket"), &storagebucket.StorageBucketConfig{
		Location: jsii.String("asia-northeast1"),
		Name:     jsii.String("gunpla-calendar-exporter"),
	})

	account := serviceaccount.NewServiceAccount(stack, jsii.String("app_sa"), &serviceaccount.ServiceAccountConfig{
		AccountId:   jsii.String("app-account"),
		DisplayName: jsii.String("app account"),
	})

	storagebucketiampolicy.NewStorageBucketIamPolicy(stack, jsii.String("sa_iam"), &storagebucketiampolicy.StorageBucketIamPolicyConfig{
		Bucket: bucket.Name(),
		PolicyData: jsii.String(`{
			"bindings": [
				{
					"role": "roles/storage.objectAdmin",
					"members": [
						"serviceAccount:` + *account.Email() + `"
					]
				},
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
