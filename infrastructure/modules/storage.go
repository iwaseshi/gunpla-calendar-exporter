package modules

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/iamworkloadidentitypool"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/iamworkloadidentitypoolprovider"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/serviceaccount"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/serviceaccountiammember"
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

	pool := iamworkloadidentitypool.NewIamWorkloadIdentityPool(stack, jsii.String("wi_pool"), &iamworkloadidentitypool.IamWorkloadIdentityPoolConfig{
		DisplayName:            jsii.String("Github Action Pool"),
		WorkloadIdentityPoolId: jsii.String("github-actions-pool"),
	})

	iamworkloadidentitypoolprovider.NewIamWorkloadIdentityPoolProvider(stack, jsii.String("wi_provider"), &iamworkloadidentitypoolprovider.IamWorkloadIdentityPoolProviderConfig{
		DisplayName:                    jsii.String("Github Action Pool Provider"),
		WorkloadIdentityPoolId:         pool.WorkloadIdentityPoolId(),
		WorkloadIdentityPoolProviderId: jsii.String("github-actions-ip-provider"),
		AttributeMapping: &map[string]*string{
			"google.subject":       jsii.String("assertion.sub"),
			"attribute.repository": jsii.String("assertion.repository"),
		},
		Oidc: &iamworkloadidentitypoolprovider.IamWorkloadIdentityPoolProviderOidc{
			IssuerUri: jsii.String("https://token.actions.githubusercontent.com"),
		},
	})

	serviceaccountiammember.NewServiceAccountIamMember(stack, jsii.String("sa_member"), &serviceaccountiammember.ServiceAccountIamMemberConfig{
		Role:             jsii.String("roles/iam.workloadIdentityUser"),
		ServiceAccountId: account.Name(),
		Member:           jsii.String("principalSet://iam.googleapis.com/" + *pool.Name() + "/attribute.repository/iwaseshi/gunpla-calendar-exporter"),
	})

	storagebucketiampolicy.NewStorageBucketIamPolicy(stack, jsii.String("sa_iam"), &storagebucketiampolicy.StorageBucketIamPolicyConfig{
		Bucket: bucket.Name(),
		PolicyData: jsii.String(`{
		"bindings": [
			{
				"role": "roles/storage.admin",
				"members": [
					"serviceAccount:` + *account.Email() + `",
					"principalSet://iam.googleapis.com/` + *pool.Name() + `/attribute.repository/iwaseshi/gunpla-calendar-exporter"
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
