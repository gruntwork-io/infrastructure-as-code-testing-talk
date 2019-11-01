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

// An example of a unit test for the Terraform module in examples/static-website
func TestStaticWebsiteUnit(t *testing.T) {
	t.Parallel()

	// A unique ID we can use to namespace all our resource names and ensure they don't clash across parallel tests
	uniqueId := random.UniqueId()

	// Since we want to be able to run multiple tests in parallel on the same modules, we need to copy them into
	// temp folders so that the state files and .terraform folders don't clash
	staticWebsitePath := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/static-website")

	// Configure the S3 backend where the static website module will store its state
	terraformBackend := configureBackendForStaticWebsite(t, uniqueId, staticWebsitePath)

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: staticWebsitePath,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name": strings.ToLower(fmt.Sprintf("static-website-test-%s", uniqueId)),
		},

		// Backend configuration that specifies where to store Terraform state for the module
		BackendConfig: terraformBackend,
	}

	// At the end of the test, clean up any resources that were created
	defer cleanupStaticWebsite(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	url := terraform.Output(t, terraformOptions, "website_url")

	// Verify the app returns a 200 OK with the text "Hello, World!"
	expectedStatus := 200
	expectedBody := "Hello, World!"
	maxRetries := 10
	timeBetweenRetries := 3 * time.Second
	http_helper.HttpGetWithRetry(t, url, nil, expectedStatus, expectedBody, maxRetries, timeBetweenRetries)
}