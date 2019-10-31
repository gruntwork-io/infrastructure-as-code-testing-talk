package test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"testing"
)

// An example of an integration test for the Terraform modules in examples/proxy-app and examples/static-website. This
// example breaks the test up into stages so you can skip any stage foo by setting the environment variable
// SKIP_foo=true.
func TestProxyAppIntegrationWithStages(t *testing.T) {
	t.Parallel()

	// Since we want to be able to run multiple tests in parallel on the same modules, we need to copy them into
	// temp folders so that the state files and .terraform folders don't clash
	staticWebsitePath := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/static-website")
	proxyAppPath := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/proxy-app")

	// Undeploy the static-website module at the end of the test
	defer test_structure.RunTestStage(t, "cleanup_static_website", func() {
		staticWebsiteOpts := test_structure.LoadTerraformOptions(t, staticWebsitePath)
		cleanupStaticWebsite(t, staticWebsiteOpts)
	})

	// Deploy the static-website module
	test_structure.RunTestStage(t, "deploy_static_website", func() {
		staticWebsiteOpts := configureStaticWebsiteOptions(t, staticWebsitePath)
		test_structure.SaveTerraformOptions(t, staticWebsitePath, staticWebsiteOpts)
		terraform.InitAndApply(t, staticWebsiteOpts)
	})

	// Undeploy the proxy-app module at the end of the test
	defer test_structure.RunTestStage(t, "cleanup_proxy_app", func() {
		proxyAppOpts := test_structure.LoadTerraformOptions(t, proxyAppPath)
		cleanupProxyApp(t, proxyAppOpts)
	})

	// Deploy the proxy-app module
	test_structure.RunTestStage(t, "deploy_proxy_app", func() {
		staticWebsiteOpts := test_structure.LoadTerraformOptions(t, staticWebsitePath)
		proxyAppOpts := configureProxyAppOptions(t, staticWebsiteOpts, proxyAppPath)
		test_structure.SaveTerraformOptions(t, proxyAppPath, proxyAppOpts)
		terraform.InitAndApply(t, proxyAppOpts)
	})

	// Validate the proxy-app module proxies the static-website correctly
	test_structure.RunTestStage(t, "validate_proxy_app", func() {
		proxyAppOpts := test_structure.LoadTerraformOptions(t, proxyAppPath)
		validateProxyApp(t, proxyAppOpts)
	})
}

