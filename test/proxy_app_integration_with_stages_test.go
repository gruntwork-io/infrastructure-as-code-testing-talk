package test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"testing"
)

// An example of an integration test for the Terraform modules in examples/proxy-app and examples/web-service. This
// example breaks the test up into stages so you can skip any stage foo by setting the environment variable
// SKIP_foo=true.
func TestProxyAppIntegrationWithStages(t *testing.T) {
	t.Parallel()

	// Since we want to be able to run multiple tests in parallel on the same modules, we need to copy them into
	// temp folders so that the state files and .terraform folders don't clash
	webServicePath := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/web-service")
	proxyAppPath := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/proxy-app")

	// Undeploy the web-service module at the end of the test
	defer test_structure.RunTestStage(t, "cleanup_web_service", func() {
		webServiceOpts := test_structure.LoadTerraformOptions(t, webServicePath)
		terraform.Destroy(t, webServiceOpts)
	})

	// Deploy the web-service module
	test_structure.RunTestStage(t, "deploy_web_service", func() {
		webServiceOpts := configWebService(t, webServicePath)
		test_structure.SaveTerraformOptions(t, webServicePath, webServiceOpts)
		terraform.InitAndApply(t, webServiceOpts)
	})

	// Undeploy the proxy-app module at the end of the test
	defer test_structure.RunTestStage(t, "cleanup_proxy_app", func() {
		proxyAppOpts := test_structure.LoadTerraformOptions(t, proxyAppPath)
		terraform.Destroy(t, proxyAppOpts)
	})

	// Deploy the proxy-app module
	test_structure.RunTestStage(t, "deploy_proxy_app", func() {
		webServiceOpts := test_structure.LoadTerraformOptions(t, webServicePath)
		proxyAppOpts := configProxyApp(t, webServiceOpts, proxyAppPath)
		test_structure.SaveTerraformOptions(t, proxyAppPath, proxyAppOpts)
		terraform.InitAndApply(t, proxyAppOpts)
	})

	// Validate the proxy-app module proxies the web-service correctly
	test_structure.RunTestStage(t, "validate_proxy_app", func() {
		proxyAppOpts := test_structure.LoadTerraformOptions(t, proxyAppPath)
		validateProxyApp(t, proxyAppOpts)
	})
}
