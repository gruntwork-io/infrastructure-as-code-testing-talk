# Docker and Kubernetes Hello World App Example

This folder contains example code that can be used to build a Docker image with a simple "Hello, World!" web app, and 
deploy the app to a [Kubernetes](https://kubernetes.io/) cluster. 

This code is used in the talk 
[How to test your infrastructure code: automated testing for Terraform, Docker, Packer, Kubernetes, and more](https://qconsf.com/sf2019/presentation/infrastructure-0) 
by [Yevgeniy Brikman](https://www.ybrikman.com/) as a representation of typical infrastructure code that deploys a web
service for which you may wish to write automated tests. 

**Note**: This repo is for demonstration and learning purposes only and should NOT be used to run anything important. 
For production-ready versions of this code and many other types of infrastructure, check out 
[Gruntwork](https://gruntwork.io/).

## Running this example manually

1. You'll need access to a [Kubernetes](https://kubernetes.io/) cluster to run deploy this example. **Recommended option**: 
   If you're using the [Docker Desktop app](https://www.docker.com/products/docker-desktop), then you already have a local 
   [Kubernetes cluster installed](https://www.docker.com/blog/kubernetes-is-now-available-in-docker-desktop-stable-channel/)!
   Alternatively, you can run Kubernetes locally using 
   [MiniKube](https://kubernetes.io/docs/setup/learning-environment/minikube/) or run these tests against a Kubernetes
   cluster in the cloud, such as [Amazon EKS](https://aws.amazon.com/eks/), 
   [Google Container Engine](https://cloud.google.com/kubernetes-engine/), or 
   [Azure Kubernetes Service](https://azure.microsoft.com/en-us/services/kubernetes-service/). 
1. Whichever option you choose to run your Kubernetes cluster, you'll need to 
   [authentiate to it](https://kubernetes.io/docs/reference/access-authn-authz/authentication/). If you're running 
   Kubernetes locally (e.g., via Docker for Desktop), you're probably already authenticated to it, so there's nothing
   to do.
1. Install [Docker](https://www.docker.com/).
1. Build the Docker image: `docker build -t gruntwork-io/hello-world-app:v1 .` 
1. Deploy the Docker image to Kubernetes: `kubectl apply -f deployment.yml`
1. See if the pods got created: `kubectl get pods`. Look for `hello-world-app-deployment`.
1. See if the service is exposed: `kubectl get services`. Look for `hello-world-app-service`.
1. From the previous step, look at the `EXTERNAL-IP` and `PORT(S)` for `hello-world-app-service` to get the IP and 
   port you can use to test the service. If you're running with a local Kubernetes cluster (e.g., via Docker for 
   Desktop), this will most likely be `localhost` and `8080`, so you can test with: `curl localhost:8080`.
1. When you're done testing, clean up all deployed resourcces: `kubectl delete -f deployment.yml`     

## Running automated tests against this example

1. Install [Docker](https://www.docker.com/).
1. Install [Golang](https://golang.org/), minimum version `1.13`.
1. You'll need access to a [Kubernetes](https://kubernetes.io/) cluster to run these tests. **Recommended option**: 
   If you're using the [Docker Desktop app](https://www.docker.com/products/docker-desktop), then you already have a local 
   [Kubernetes cluster installed](https://www.docker.com/blog/kubernetes-is-now-available-in-docker-desktop-stable-channel/)!
   Alternatively, you can run Kubernetes locally using 
   [MiniKube](https://kubernetes.io/docs/setup/learning-environment/minikube/) or run these tests against a Kubernetes
   cluster in the cloud, such as [Amazon EKS](https://aws.amazon.com/eks/), 
   [Google Container Engine](https://cloud.google.com/kubernetes-engine/), or 
   [Azure Kubernetes Service](https://azure.microsoft.com/en-us/services/kubernetes-service/). 
1. Whichever option you choose to run your Kubernetes cluster, you'll need to 
   [authentiate to it](https://kubernetes.io/docs/reference/access-authn-authz/authentication/). If you're running 
   Kubernetes locally (e.g., via Docker for Desktop), you're probably already authenticated to it, so there's nothing
   to do.
1. `cd test`
1. Run the automated test for this example: `go test -v -timeout 15m -run '^TestDockerKubernetesUnit$'`
