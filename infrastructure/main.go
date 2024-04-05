package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-google-go/google/v6/provider"
	"github.com/cdktf/cdktf-provider-google-go/google/v6/storagebucket"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	provider.NewGoogleProvider(stack, jsii.String("google"), &provider.GoogleProviderConfig{
		Project: jsii.String("gunpla-calendar-exporter"),
		Region:  jsii.String("asia-northeast1"),
		Zone:    jsii.String("asia-northeast1-a"),
	})

	cdktf.NewGcsBackend(stack, &cdktf.GcsBackendConfig{
		Bucket: jsii.String("gunpla-calendar-exporter-backend"),
		Prefix: jsii.String("terraform/state"),
		// 個人開発のためstate lockは行わない。
	})

	storagebucket.NewStorageBucket(stack, jsii.String("gcs_bucket"), &storagebucket.StorageBucketConfig{
		Location: jsii.String("asia-northeast1"),
		Name:     jsii.String("gunpla-calendar-exporter"),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)
	NewMyStack(app, "gunpla-calendar-exporter")
	app.Synth()
}
