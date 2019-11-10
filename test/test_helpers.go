package test

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/aws"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// Config the web-service module
func configWebService(t *testing.T, webServicePath string) *terraform.Options {
	// A unique ID we can use to namespace all our resource names and ensure they don't clash across parallel tests
	uniqueId := random.UniqueId()

	// Configure the S3 backend where the web-service module will store its state
	terraformBackend := configureBackendForWebService(t, uniqueId, webServicePath)

	return &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: webServicePath,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name": strings.ToLower(fmt.Sprintf("test-proxy-app-%s", uniqueId)),
		},

		// Backend configuration that specifies where to store Terraform state for the module
		BackendConfig: terraformBackend,

		RetryableTerraformErrors: map[string]string{
			"net/http: TLS handshake timeout": "https://github.com/hashicorp/terraform/issues/22456",
		},

		MaxRetries: 3,

		TimeBetweenRetries: 3 * time.Second,
	}
}

// Clean up the web-service module
func cleanupWebService(t *testing.T, webServiceOpts *terraform.Options) {
	// Run `terraform destroy` to clean up any resources that were created
	terraform.Destroy(t, webServiceOpts)

	// Run `terraform destroy` to clean up any resources that were created
	s3BucketRegion := readConfig(t, webServiceOpts.BackendConfig, "region")
	s3BucketName := readConfig(t, webServiceOpts.BackendConfig, "bucket")
	cleanupTerraformBackend(t, s3BucketRegion, s3BucketName)
}

// Deploy the proxy app
func configProxyApp(t *testing.T, webServiceOpts *terraform.Options, proxyAppPath string) *terraform.Options {
	name := fmt.Sprintf("%s-proxy-app", readConfig(t, webServiceOpts.Vars, "name"))
	s3BucketRegion := readConfig(t, webServiceOpts.BackendConfig, "region")
	s3BucketName := readConfig(t, webServiceOpts.BackendConfig, "bucket")
	s3BucketKey := readConfig(t, webServiceOpts.BackendConfig, "key")

	return &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: proxyAppPath,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name": name,

			// To make this an integration test, pass in the web-service Terraform state data so that the app
			// proxies the contents of this bucket
			"terraform_state_bucket": s3BucketName,
			"terraform_state_bucket_region": s3BucketRegion,
			"terraform_state_bucket_web_service_key": s3BucketKey,
		},
	}
}

// Clean up the proxy app
func cleanupProxyApp(t *testing.T, proxyAppOpts *terraform.Options) {
	// Run `terraform destroy` to clean up any resources that were created
	terraform.Destroy(t, proxyAppOpts)
}

// Validate the proxy app works
func validateProxyApp(t *testing.T, proxyAppOpts *terraform.Options) {
	// Run `terraform output` to get the values of output variables
	url := terraform.Output(t, proxyAppOpts, "url")

	// Verify the app returns a 200 OK with JSON data that contains the text "Hello, World!", which is what the
	// web-service module should be returning
	expectedStatus := 200
	expectedBody := `{"text":"Hello, World!"}`
	maxRetries := 10
	timeBetweenRetries := 3 * time.Second
	http_helper.HttpGetWithRetry(t, url, nil, expectedStatus, expectedBody, maxRetries, timeBetweenRetries)
}

// Create an S3 bucket to use as a Terraform backend and return the backend details in the format expected by Terratest
func configureBackendForWebService(t *testing.T, uniqueId string, webServicePath string) map[string]interface{} {
	s3BucketName := strings.ToLower(fmt.Sprintf("test-proxy-app-state-%s", uniqueId))
	s3BucketRegion := "us-east-2"
	s3BucketKey := "web-service/terraform.tfstate"

	// Create an S3 bucket to store Terraform state
	aws.CreateS3Bucket(t, s3BucketRegion, s3BucketName)

	// Clean up any previous terraform.tfstate that may be referencing an S3 bucket that was deleted in a previous
	// test run and no longer exists
	os.Remove(filepath.Join(webServicePath, ".terraform", "terraform.tfstate"))

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
