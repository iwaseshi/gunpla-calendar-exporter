package main

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-google-go/google/v6/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func SetupGoogleProvider(stack cdktf.TerraformStack) {
	provider.NewGoogleProvider(stack, jsii.String("google"), &provider.GoogleProviderConfig{
		Project: jsii.String("gunpla-calendar-exporter"),
		Region:  jsii.String("asia-northeast1"),
		Zone:    jsii.String("asia-northeast1-a"),
	})
}
