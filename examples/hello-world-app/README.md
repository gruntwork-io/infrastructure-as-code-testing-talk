# Hello World App Example

This folder contains example code that shows how to use the [serverless-app module](/modules/serverless-app) to deploy
a "Hello, World" Node.js app on top of [AWS Lambda](https://aws.amazon.com/lambda/) and 
[API Gateway](https://aws.amazon.com/api-gateway/). 

This code is used in the talk 
[How to test your infrastructure code: automated testing for Terraform, Docker, Packer, Kubernetes, and more](https://www.infoq.com/presentations/automated-testing-terraform-docker-packer/) 
by [Yevgeniy Brikman](https://www.ybrikman.com/) as a representation of typical infrastructure code that deploys a web
service for which you may wish to write automated tests. 

**Note**: This repo is for demonstration and learning purposes only and should NOT be used to run anything important. 
For production-ready versions of this code and many other types of infrastructure, check out 
[Gruntwork](https://gruntwork.io/).

**WARNING**: This module and the automated tests for it deploy real resources into your AWS account which can cost you
money. The resources are all part of the [AWS Free Tier](https://aws.amazon.com/free/), so if you haven't used that up,
it should be free, but you are completely responsible for all AWS charges.

## Running this example manually

1. Sign up for [AWS](https://aws.amazon.com/).
1. Configure your AWS credentials using one of the [supported methods for AWS CLI
   tools](https://blog.gruntwork.io/a-comprehensive-guide-to-authenticating-to-aws-on-the-command-line-63656a686799), 
   such as setting the `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables. 
1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`.
1. Run `terraform init`.
1. Run `terraform apply`.
1. This module will output the URL of the "Hello, World" app at the end of `apply`. Try this URL out in your browser or
   via `curl` to see if it's working!
1. When you're done, run `terraform destroy`.

## Running automated tests against this example

1. Sign up for [AWS](https://aws.amazon.com/).
1. Configure your AWS credentials using one of the [supported methods for AWS CLI
   tools](https://blog.gruntwork.io/a-comprehensive-guide-to-authenticating-to-aws-on-the-command-line-63656a686799), 
   such as setting the `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables. 
1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`.
1. Install [Golang](https://golang.org/), minimum version `1.13`.
1. `cd test`
1. To run the unit test for this example: `go test -v -timeout 15m -run '^TestHelloWorldAppUnit$'`
