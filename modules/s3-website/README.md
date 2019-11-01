# s3-website module

This folder contains a Terraform module that can be used to deploy a static website on top of 
[Amazon S3](https://aws.amazon.com/s3/). It's used in the
talk [How to test your infrastructure code: automated testing for Terraform, 
Docker, Packer, Kubernetes, and more](https://qconsf.com/sf2019/presentation/infrastructure-0) by 
[Yevgeniy Brikman](https://www.ybrikman.com/) as a representation of typical infrastructure code that deploys an 
external dependency, such as a data store, that you may wish to mock when writing automated tests.

**Note**: This repo is for demonstration and learning purposes only and should NOT be used to run anything important. 
For production-ready versions of this code and many other types of infrastructure, check out 
[Gruntwork](https://gruntwork.io/).

## Features

* Create an S3 bucket configured for website serving.
* Configure an ACL and IAM policy to make the bucket publicly accessible.
* Upload dummy index.html and error.html pages.

## Usage

Check out [examples/static-website](/examples/static-website) for example usage.

## Tests

Check out [test/proxy_app_integration_test.go](/test/proxy_app_integration_test.go) and
[test/proxy_app_integration_with_stages_test.go](/test/proxy_app_integration_with_stages_test.go) for examples of 
automated tests for this module.
