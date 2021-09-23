package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestExamplesComplete(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../../examples/complete",
		Upgrade:      true,
		VarFiles:     []string{},
		Vars:         map[string]interface{}{},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	foo := terraform.Output(t, terraformOptions, "foo")

	assert.Equal(t, "baz", foo)
}
