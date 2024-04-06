package main

import (
	"cdk.tf/go/stack/modules"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	SetupGoogleProvider(stack)
	SetupGcsBackend(stack)

	modules.NewStorageBucket(stack)

	return stack
}

func main() {
	app := cdktf.NewApp(nil)
	NewMyStack(app, "gunpla-calendar-exporter")
	app.Synth()
}
