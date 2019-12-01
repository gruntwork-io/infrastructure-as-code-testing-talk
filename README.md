# Infrastructure as code testing talk

This repo contains sample code for the talk [How to test your infrastructure code: automated testing for Terraform,
Docker, Packer, Kubernetes, and more](https://qconsf.com/sf2019/presentation/infrastructure-0) by
[Yevgeniy Brikman](https://www.ybrikman.com/), with updates for AWS CDK by Ben Whaley.

**Note**: This repo is for demonstration and learning purposes only and should NOT be used to run anything important.
For production-ready versions of this code and many other types of infrastructure, check out
[Gruntwork](https://gruntwork.io/).

## Overview of the repo

This repo contains:

* [modules](/modules): a simple Terraform module used throughout the talk to demonstrate automated testing practices, and an [AWS CDK](https://github.com/aws/aws-cdk/) app used to demonstrate techniques for testing infrastructure created with the CDK.
* [examples](/examples): examples of (a) how to use the Terraform module in `/modules` and (b) how to deploy
  Dockerized apps to Kubernetes.
* [test](/test): automated tests for each of the examples in the `/examples` folder.

Dive into each of the folders above for more information!

## Running the examples manually

Check out the README in each of the examples in the [examples](/examples) folder for instructions on how to run them
manually.

## Running the automated tests

Check out the README in the [test](/test) folder for instructions on how to run the automated tests.

## License

See [LICENSE.txt](LICENSE.txt).
