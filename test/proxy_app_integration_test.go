package test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"testing"
)

// An example of an integration test for the Terraform modules in examples/proxy-app and examples/static-website.
func TestProxyAppIntegration(t *testing.T) {
	t.Parallel()

	// Deploy the static-website module
	staticWebsiteOpts := configureStaticWebsiteOptions(t)
	defer cleanupStaticWebsite(t, staticWebsiteOpts)
	terraform.InitAndApply(t, staticWebsiteOpts)

	// Deploy the proxy-app module
	proxyAppOpts := configureProxyAppOptions(t, staticWebsiteOpts)
	defer cleanupProxyApp(t, proxyAppOpts)
	terraform.InitAndApply(t, proxyAppOpts)

	// Validate the proxy-app module proxies the static-website correctly
	validateProxyApp(t, proxyAppOpts)
}