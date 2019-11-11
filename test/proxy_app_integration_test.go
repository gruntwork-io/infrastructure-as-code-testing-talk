package test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"testing"
)

// An example of an integration test for the Terraform modules in examples/proxy-app and examples/web-service.
func TestProxyAppIntegration(t *testing.T) {
	t.Parallel()

	// Since we want to be able to run multiple tests in parallel on the same modules, we need to copy them into
	// temp folders so that the state files and .terraform folders don't clash
	webServicePath := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/web-service")
	proxyAppPath := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/proxy-app")

	// Deploy the web-service module
	webServiceOpts := configWebService(t, webServicePath)
	defer terraform.Destroy(t, webServiceOpts)
	terraform.InitAndApply(t, webServiceOpts)

	// Deploy the proxy-app module
	proxyAppOpts := configProxyApp(t, webServiceOpts, proxyAppPath)
	defer terraform.Destroy(t, proxyAppOpts)
	terraform.InitAndApply(t, proxyAppOpts)

	// Validate the proxy-app module proxies the web-service correctly
	validateProxyApp(t, proxyAppOpts)
}
