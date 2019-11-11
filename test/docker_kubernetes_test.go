package test

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/docker"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/random"
	"strings"
	"testing"
	"time"
)

// An example of a unit test for the Docker app deployed to Kubernetes in examples/docker-kubernetes
func TestDockerKubernetesUnit(t *testing.T) {
	t.Parallel()

	// Build the example Docker image
	buildDockerImage(t)

	// Path to the Kubernetes resource config we will test
	kubeResourcePath := "../examples/docker-kubernetes/deployment.yml"

	// To ensure we can reuse the resource config on the same cluster to test different scenarios, we setup a unique
	// namespace for the resources for this test. Note that namespaces must be lowercase.
	namespaceName := strings.ToLower(random.UniqueId())

	// Setup the kubectl config and context. Here we choose to use the defaults, which is:
	// - HOME/.kube/config for the kubectl config file
	// - Current context of the kubectl config file
	options := k8s.NewKubectlOptions("", "", namespaceName)

	// At the end of the test, make sure to delete the namespace
	defer k8s.DeleteNamespace(t, options, namespaceName)

	// Create the namespace
	k8s.CreateNamespace(t, options, namespaceName)

	// At the end of the test, run `kubectl delete -f RESOURCE_CONFIG` to clean up any resources that were created.
	defer k8s.KubectlDelete(t, options, kubeResourcePath)

	// This will run `kubectl apply -f RESOURCE_CONFIG` and fail the test if there are any errors
	k8s.KubectlApply(t, options, kubeResourcePath)

	// Validate the app is working
	validateK8SApp(t, options)
}

// Build the example Docker image
func buildDockerImage(t *testing.T) {
	options := &docker.BuildOptions{
		Tags: []string{"gruntwork-io/hello-world-app:v1"},
	}
	docker.Build(t, "../examples/docker-kubernetes", options)
}

// Validate the app is working
func validateK8SApp(t *testing.T, options *k8s.KubectlOptions) {
	// This will wait up to 10 seconds for the service to become available, to ensure that we can access it.
	k8s.WaitUntilServiceAvailable(t, options, "hello-world-app-service", 10, 1*time.Second)

	// Now we verify that the service will successfully boot and start serving requests
	url := serviceUrl(t, options)
	expectedStatus := 200
	expectedBody := "Hello, World!"
	maxRetries := 10
	timeBetweenRetries := 3 * time.Second
	http_helper.HttpGetWithRetry(t, url, nil, expectedStatus, expectedBody, maxRetries, timeBetweenRetries)
}

// Get the service URL from Kubernetes
func serviceUrl(t *testing.T, options *k8s.KubectlOptions) string {
	service := k8s.GetService(t, options, "hello-world-app-service")
	endpoint := k8s.GetServiceEndpoint(t, options, service, 8080)
	return fmt.Sprintf("http://%s", endpoint)
}