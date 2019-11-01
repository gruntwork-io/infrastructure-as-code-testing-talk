# Ruby App Examples

This folder contains several simple example Ruby apps:

* `hello_world_original.rb`: A "Hello, World" Ruby Web Server.
* `hello_world_unit_testable.rb`: A "Hello, World" Ruby Web Server. This version of the code has been refactored into
  standalone, smaller units so it is easier to unit test.
* `web_service_with_dependencies.rb`: A Ruby Web Server that calls an external web service (by default, `example.com`)
  and proxies its response.
* `web_service_with_dependencies_testable.rb`: A Ruby Web Server that calls an external web service (by default, 
  `example.com`) and proxies its response. This version of the code has been refactored with dependency injection so it 
  is easier to unit test.

This code is used in the talk 
[How to test your infrastructure code: automated testing for Terraform, Docker, Packer, Kubernetes, and more](https://qconsf.com/sf2019/presentation/infrastructure-0) 
by [Yevgeniy Brikman](https://www.ybrikman.com/) as a way to demonstrate common practices used for automated tests in
general purpose programming languages that you can apply to infrastructure code. 

**Note**: This repo is for demonstration and learning purposes only and should NOT be used to run anything important. 
For production-ready versions of this code and many other types of infrastructure, check out 
[Gruntwork](https://gruntwork.io/).

## Running this example manually

1. Install [Ruby](https://www.ruby-lang.org/en/).
1. Run `ruby <SCRIPT>`. E.g., `ruby hello_world_original.rb`.
1. Test the web servers by hitting http://localhost:8000 in your web browser or using `curl`.

## Running automated tests against this example

1. Install [Ruby](https://www.ruby-lang.org/en/).
1. `cd test`
1. To run the unit tests for `hello_world_unit_testable.rb`: `ruby hello_world_app_unit_test.go`
1. To run the unit tests for `web_service_with_dependencies_testable.rb`: `ruby web_service_with_dependencies_unit_test.go`
1. To run the integration tests for `web_service_with_dependencies_testable.rb`: `ruby web_service_with_dependencies_integration_test.go`
