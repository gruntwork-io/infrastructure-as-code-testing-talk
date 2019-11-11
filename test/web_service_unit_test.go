package test

import (
	"fmt"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/random"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// An example of a unit test for the Terraform module in examples/web-service
func TestWebServiceUnit(t *testing.T) {
	t.Parallel()

	// A unique ID we can use to namespace all our resource names and ensure they don't clash across parallel tests
	uniqueId := random.UniqueId()

	// Since we want to be able to run multiple tests in parallel on the same modules, we need to copy them into
	// temp folders so that the state files and .terraform folders don't clash
	webServicePath := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/web-service")

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: webServicePath,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name": strings.ToLower(fmt.Sprintf("web-service-test-%s", uniqueId)),
		},
	}

	// At the end of the test, clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	url := terraform.Output(t, terraformOptions, "url")

	// Verify the app returns a 200 OK with JSON data that contains the text "Hello, World!"
	expectedStatus := 200
	expectedBody := `{"text":"Hello, World!"}`
	maxRetries := 10
	timeBetweenRetries := 3 * time.Second
	http_helper.HttpGetWithRetry(t, url, nil, expectedStatus, expectedBody, maxRetries, timeBetweenRetries)
}
