package test

import (
	"fmt"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"strings"
	"testing"
	"time"
)

// Config the web-service module
func configWebService(t *testing.T, webServicePath string) *terraform.Options {
	// A unique ID we can use to namespace all our resource names and ensure they don't clash across parallel tests
	uniqueId := random.UniqueId()

	return &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: webServicePath,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name": strings.ToLower(fmt.Sprintf("test-proxy-app-%s", uniqueId)),
		},

		// An example of how to automatically retry on known errors
		RetryableTerraformErrors: map[string]string{
			"net/http: TLS handshake timeout": "https://github.com/hashicorp/terraform/issues/22456",
		},
		MaxRetries:         3,
		TimeBetweenRetries: 3 * time.Second,
	}
}

// Deploy the proxy app
func configProxyApp(t *testing.T, webServiceOpts *terraform.Options, proxyAppPath string) *terraform.Options {
	name := fmt.Sprintf("%s-proxy-app", readConfig(t, webServiceOpts.Vars, "name"))
	url := terraform.Output(t, webServiceOpts, "url")

	return &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: proxyAppPath,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name":         name,
			"url_to_proxy": url,
		},
	}
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
