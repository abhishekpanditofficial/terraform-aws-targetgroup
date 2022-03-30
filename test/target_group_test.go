package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformTargetGroupExample(t *testing.T) {
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/target-group",
		Vars: map[string]interface{}{
			"target_group_name":"test-target-group-name",
		},
	})

	terraform.InitAndApply(t, terraformOptions)

	defer terraform.Destroy(t, terraformOptions)

	// VPC test
	vpc_cidr_block := terraform.Output(t, terraformOptions, "vpc_cidr_block")
	vpc_id := terraform.Output(t, terraformOptions, "vpc_id")
	assert.Equal(t, "10.0.0.0/16", vpc_cidr_block)

	// Target Group test
	target_group_name := terraform.Output(t, terraformOptions, "target_group_name")
	assert.Equal(t, "test-target-group-name", target_group_name)

	// Target Group With VPC Test
	target_group_vpc := terraform.Output(t, terraformOptions, "target_group_vpc")
	assert.Equal(t, vpc_id, target_group_vpc)

}