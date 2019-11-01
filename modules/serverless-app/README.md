# serverless-app module

This folder contains a Terraform module that can be used to deploy a serverless web application on top of 
[AWS Lambda](https://aws.amazon.com/lambda/) and [API Gateway](https://aws.amazon.com/api-gateway/). It's used in the
talk [How to test your infrastructure code: automated testing for Terraform, 
Docker, Packer, Kubernetes, and more](https://qconsf.com/sf2019/presentation/infrastructure-0) by 
[Yevgeniy Brikman](https://www.ybrikman.com/) as a representation of typical infrastructure code that deploys a web
service for which you may wish to write automated tests. 

**Note**: This repo is for demonstration and learning purposes only and should NOT be used to run anything important. 
For production-ready versions of this code and many other types of infrastructure, check out 
[Gruntwork](https://gruntwork.io/).

## Features

* Create a Lambda function with configurable runtime, memory size, timeout, and environment variables. 
* Automatically zip up a specified folder and upload it as a Lambda deployment package.
* Create an IAM role for the Lambda function.
* Create an API Gateway that proxies ALL requests to the Lambda function.

## Usage

Check out [examples/hello-world-app](/examples/hello-world-app) and [examples/proxy-app](/examples/proxy-app) for 
example usage.

## Tests

Check out [test/hello_world_app_unit_test.go](/test/hello_world_app_unit_test.go), 
[test/proxy_app_unit_test.go](/test/proxy_app_unit_test.go), 
[test/proxy_app_integration_test.go](/test/proxy_app_integration_test.go), and
[test/proxy_app_integration_with_stages_test.go](/test/proxy_app_integration_with_stages_test.go) for examples of 
automated tests for this module.
