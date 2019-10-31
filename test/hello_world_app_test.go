package test

import (
	"fmt"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/random"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// An example of how to test the Terraform module in examples/hello-world-app
func TestHelloWorldApp(t *testing.T) {
	t.Parallel()

	uniqueId := random.UniqueId()

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/hello-world-app",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name": fmt.Sprintf("hello-world-app-%s", uniqueId),
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	url := terraform.Output(t, terraformOptions, "url")

	// Verify the app returns a 200 OK with the text "Hello, World!"
	expectedStatus := 200
	expectedBody := "Hello, World!"
	maxRetries := 10
	timeBetweenRetries := 3 * time.Second
	http_helper.HttpGetWithRetry(t, url, nil, expectedStatus, expectedBody, maxRetries, timeBetweenRetries)
}