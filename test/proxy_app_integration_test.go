package test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"testing"
)

// An example of an integration test for the Terraform modules in examples/proxy-app and examples/static-website.
func TestProxyAppIntegration(t *testing.T) {
	t.Parallel()

	// Since we want to be able to run multiple tests in parallel on the same modules, we need to copy them into
	// temp folders so that the state files and .terraform folders don't clash
	staticWebsitePath := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/static-website")
	proxyAppPath := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/proxy-app")

	// Deploy the static-website module
	staticWebsiteOpts := configureStaticWebsiteOptions(t, staticWebsitePath)
	defer cleanupStaticWebsite(t, staticWebsiteOpts)
	terraform.InitAndApply(t, staticWebsiteOpts)

	// Deploy the proxy-app module
	proxyAppOpts := configureProxyAppOptions(t, staticWebsiteOpts, proxyAppPath)
	defer cleanupProxyApp(t, proxyAppOpts)
	terraform.InitAndApply(t, proxyAppOpts)

	// Validate the proxy-app module proxies the static-website correctly
	validateProxyApp(t, proxyAppOpts)
}