package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestExamplesComplete(t *testing.T) {
	t.Parallel()

	rand.Seed(time.Now().UnixNano())
	randID := rand.Intn(100000)

	expectedBucketId := fmt.Sprintf("thirstydev-%d", randID)

	terraformOptions := &terraform.Options{
		TerraformBinary: "terragrunt",
		TerraformDir:    "../../examples/complete",
		Upgrade:         true,
		VarFiles:        []string{},
		Vars: map[string]interface{}{
			"id": expectedBucketId,
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.TgApplyAll(t, terraformOptions)

	actualBucketId := terraform.Output(t, terraformOptions, "bucket_id")

	assert.Equal(t, expectedBucketId, actualBucketId)
}
