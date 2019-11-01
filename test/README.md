# Automated tests

This folder contains the automated tests for all of the example code in the [/examples folder](/examples).

**Note**: This repo is for demonstration and learning purposes only and should NOT be used to run anything important. 
For production-ready versions of this code and many other types of infrastructure, check out 
[Gruntwork](https://gruntwork.io/).

## Running automated tests for the Terraform examples

**WARNING**: The Terraform modules and their automated tests deploy real resources into your AWS account which can cost 
you money. The resources are all part of the [AWS Free Tier](https://aws.amazon.com/free/), so if you haven't used that 
up, it should be free, but you are completely responsible for all AWS charges.

1. Sign up for [AWS](https://aws.amazon.com/).
1. Configure your AWS credentials using one of the [supported methods for AWS CLI
   tools](https://blog.gruntwork.io/a-comprehensive-guide-to-authenticating-to-aws-on-the-command-line-63656a686799), 
   such as setting the `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables. 
1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`.
1. Install [Golang](https://golang.org/), minimum version `1.13`.
1. `cd test`
1. To run all the tests: `go test -v -timeout 15m`
1. To run a single test: `go test -v -timeout 15m -run <TEST_NAME>`

## Running automated tests for the Ruby examples

1. Install [Ruby](https://www.ruby-lang.org/en/).
1. `cd test`
1. `ruby <FILE_NAME>`
