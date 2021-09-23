package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestExamplesComplete(t *testing.T) {
	t.Parallel()

	//	rand.Seed(time.Now().UnixNano())
	//	randID := strconv.Itoa(rand.Intn(100000))

	terraformOptions := &terraform.Options{
		TerraformDir: "../../examples/complete",
		Upgrade:      true,
		VarFiles:     []string{},
		Vars:         map[string]interface{}{},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	bucketId := terraform.Output(t, terraformOptions, "bucket_id")

	assert.Equal(t, "foo", bucketId)
}
