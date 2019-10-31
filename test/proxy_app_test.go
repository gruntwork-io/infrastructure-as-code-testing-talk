package test

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/aws"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/random"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

const proxyAppPath = "../examples/proxy-app"
const staticWebsitePath = "../examples/static-website"

// An example of a unit test for the Terraform module in examples/proxy-app
func TestProxyAppUnit(t *testing.T) {
	t.Parallel()

	uniqueId := random.UniqueId()

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: proxyAppPath,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name": fmt.Sprintf("proxy-app-%s", uniqueId),

			// To make this a unit test, we proxy a known, "mock" endpoint
			"url_to_proxy": "https://www.example.com",
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	url := terraform.Output(t, terraformOptions, "url")

	// Verify the app returns a 200 OK with the text from example.com
	maxRetries := 10
	timeBetweenRetries := 3 * time.Second
	http_helper.HttpGetWithRetryWithCustomValidation(t, url, nil, maxRetries, timeBetweenRetries, func(statusCode int, body string) bool {
		return statusCode == 200 && strings.Contains(body, "<h1>Example Domain</h1>")
	})
}

// An example of an integration test for the Terraform modules in examples/proxy-app and examples/static-website. This
// example breaks the test up into stages so you can skip any stage foo by setting the environment variable
// SKIP_foo=true.
func TestProxyAppIntegrationWithStages(t *testing.T) {
	t.Parallel()

	defer test_structure.RunTestStage(t, "cleanup_static_website", func() {
		cleanupStaticWebsite(t)
	})

	test_structure.RunTestStage(t, "deploy_static_website", func() {
		deployStaticWebsite(t)
	})

	defer test_structure.RunTestStage(t, "cleanup_proxy_app", func() {
		cleanupProxyApp(t)
	})

	test_structure.RunTestStage(t, "deploy_proxy_app", func() {
		deployProxyApp(t)
	})

	test_structure.RunTestStage(t, "validate_proxy_app", func() {
		validateProxyApp(t)
	})
}

// Deploy the static website
func deployStaticWebsite(t *testing.T) {
	uniqueId := random.UniqueId()
	terraformBackend := configureBackendForStaticWebsite(t, uniqueId)

	staticWebsiteOpts := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: staticWebsitePath,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name": strings.ToLower(fmt.Sprintf("test-proxy-app-%s", uniqueId)),
		},

		// Backend configuration that specifies where to store Terraform state for the module
		BackendConfig: terraformBackend,
	}

	// Save the Terraform options to disk so other test stages can read them
	test_structure.SaveTerraformOptions(t, staticWebsitePath, staticWebsiteOpts)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, staticWebsiteOpts)
}

// Clean up the static website
func cleanupStaticWebsite(t *testing.T) {
	// Read the Terraform options saved by an earlier test stage
	staticWebsiteOpts := test_structure.LoadTerraformOptions(t, staticWebsitePath)

	// Run `terraform destroy` to clean up any resources that were created
	terraform.Destroy(t, staticWebsiteOpts)

	// Run `terraform destroy` to clean up any resources that were created
	s3BucketRegion := readConfig(t, staticWebsiteOpts.BackendConfig, "region")
	s3BucketName := readConfig(t, staticWebsiteOpts.BackendConfig, "bucket")
	cleanupTerraformBackend(t, s3BucketRegion, s3BucketName)
}

// Deploy the proxy app
func deployProxyApp(t *testing.T) {
	// Read the Terraform options saved by an earlier test stage
	staticWebsiteOpts := test_structure.LoadTerraformOptions(t, staticWebsitePath)

	name := readConfig(t, staticWebsiteOpts.Vars, "name")
	s3BucketRegion := readConfig(t, staticWebsiteOpts.BackendConfig, "region")
	s3BucketName := readConfig(t, staticWebsiteOpts.BackendConfig, "bucket")
	s3BucketKey := readConfig(t, staticWebsiteOpts.BackendConfig, "key")

	proxyAppOpts := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: proxyAppPath,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name": name,

			// To make this an integration test, pass in the static website Terraform state data so that the app
			// proxies the contents of this bucket
			"terraform_state_bucket": s3BucketName,
			"terraform_state_bucket_region": s3BucketRegion,
			"terraform_state_bucket_static_website_key": s3BucketKey,
		},
	}

	// Save the Terraform options to disk so other test stages can read them
	test_structure.SaveTerraformOptions(t, proxyAppPath, proxyAppOpts)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, proxyAppOpts)
}

// Clean up the proxy app
func cleanupProxyApp(t *testing.T) {
	// Read the Terraform options saved by an earlier test stage
	proxyAppOpts := test_structure.LoadTerraformOptions(t, proxyAppPath)

	// Run `terraform destroy` to clean up any resources that were created
	terraform.Destroy(t, proxyAppOpts)
}

// Validate the proxy app works
func validateProxyApp(t *testing.T) {
	// Read the Terraform options saved by an earlier test stage
	proxyAppOpts := test_structure.LoadTerraformOptions(t, proxyAppPath)

	// Run `terraform output` to get the values of output variables
	url := terraform.Output(t, proxyAppOpts, "url")

	// Verify the app returns a 200 OK with the text "Hello, World!", which is what the index.html of the static
	// website contains
	expectedStatus := 200
	expectedBody := "Hello, World!"
	maxRetries := 10
	timeBetweenRetries := 3 * time.Second
	http_helper.HttpGetWithRetry(t, url, nil, expectedStatus, expectedBody, maxRetries, timeBetweenRetries)
}

// Create an S3 bucket to use as a Terraform backend and return the backend details in the format expected by Terratest
func configureBackendForStaticWebsite(t *testing.T, uniqueId string) map[string]interface{} {
	s3BucketName := strings.ToLower(fmt.Sprintf("test-proxy-app-state-%s", uniqueId))
	s3BucketRegion := "us-east-2"
	s3BucketKey := "static-website/terraform.tfstate"

	// Create an S3 bucket to store Terraform state
	aws.CreateS3Bucket(t, s3BucketRegion, s3BucketName)

	// Clean up any previous terraform.tfstate that may be referencing an S3 bucket that was deleted in a previous
	// test run and no longer exists
	os.Remove(filepath.Join(staticWebsitePath, ".terraform", "terraform.tfstate"))

	return map[string]interface{}{
		"bucket": s3BucketName,
		"region": s3BucketRegion,
		"key": s3BucketKey,
	}
}

// Clean up the S3 bucket used as a Terraform backend
func cleanupTerraformBackend(t *testing.T, s3BucketRegion string, s3BucketName string) {
	aws.EmptyS3Bucket(t, s3BucketRegion, s3BucketName)
	aws.DeleteS3Bucket(t, s3BucketRegion, s3BucketName)
}

// Read a config from a backend or vars map of terraform.Options and return its value as a string. If the key isn't
// present or isn't a string, fail the test.
func readConfig(t *testing.T, values map[string]interface{}, key string) string {
	value, present := values[key]
	if !present {
		t.Fatalf("Required key %s not found", key)
	}

	valueAsString, isString := value.(string)
	if !isString {
		t.Fatalf("Key %s was not a string", key)
	}

	return valueAsString
}