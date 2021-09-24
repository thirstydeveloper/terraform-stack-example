package test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestExamplesComplete(t *testing.T) {
	t.Parallel()

	rand.Seed(time.Now().UnixNano())

	namespace := "thirstydev"
	environment := strconv.Itoa(rand.Intn(100000))

	terraformOptions := &terraform.Options{
		TerraformBinary: "terragrunt",
		// Use the verify stack as the target so we can access its outputs TerraformDir: "../../examples/complete/verify",
		TerraformDir: "../../examples/complete/verify",
		Upgrade:      true,
		// These vars are used by the fixture only. We cannot pass them to
		// vars because then our stack under test will fail complaining
		// about -var CLI arguments for undeclared variables. We can't use
		// VarFiles because we need dynamic values (environment) and our
		// test needs easy access to them.
		//
		// See:
		//
		// https://www.terraform.io/docs/language/values/variables.html#values-for-undeclared-variables
		EnvVars: map[string]string{
			"TF_VAR_fixture_namespace":   namespace,
			"TF_VAR_fixture_environment": environment,

			// Since we are specifying the verify stack as the terragrunt dir,
			// we need to include its dependencies so the fixture and stack
			// code is run.
			"TERRAGRUNT_INCLUDE_EXTERNAL_DEPENDENCIES": "true",
		},
		VarFiles: []string{},
		Vars:     map[string]interface{}{},
	}

	defer terraform.TgDestroyAll(t, terraformOptions)

	terraform.TgApplyAll(t, terraformOptions)

	var bucket map[string]interface{}

	terraform.OutputStruct(t, terraformOptions, "bucket", &bucket)

	assert.Equal(t, fmt.Sprintf("%s-%s", namespace, environment), bucket["id"])
}
