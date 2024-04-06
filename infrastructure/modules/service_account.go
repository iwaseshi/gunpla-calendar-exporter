package modules

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-google-go/google/v6/serviceaccount"
	"github.com/cdktf/cdktf-provider-google-go/google/v6/storagebucketiampolicy"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewAppServiceAccount(stack cdktf.TerraformStack) {
	account := serviceaccount.NewServiceAccount(stack, jsii.String("app_sa"), &serviceaccount.ServiceAccountConfig{
		AccountId:   jsii.String("app-account"),
		DisplayName: jsii.String("app account"),
	})

	storagebucketiampolicy.NewStorageBucketIamPolicy(stack, jsii.String("sa_iam"), &storagebucketiampolicy.StorageBucketIamPolicyConfig{
		Bucket: jsii.String("gunpla-calendar-exporter"),
		PolicyData: jsii.String(`{
			"bindings": [
				{
					"role": "roles/storage.objectAdmin",
					"members": [
						"serviceAccount:` + *account.Email() + `"
					]
				}
			]
		}`),
	})
}
