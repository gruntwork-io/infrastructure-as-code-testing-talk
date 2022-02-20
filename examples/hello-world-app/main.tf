terraform {
  # This module is now only being tested with Terraform 1.1.x. However, to make upgrading easier, we are setting 1.0.0 as the minimum version.
  required_version = ">= 1.0.0"
}

provider "aws" {
  region = var.aws_region
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE A SERVERLESS "HELLO, WORLD" APP
# ---------------------------------------------------------------------------------------------------------------------

module "hello_world_app" {
  source = "../../modules/serverless-app"

  name = var.name

  source_dir = "${path.module}/javascript"
  runtime    = "nodejs14.x"
  handler    = "index.handler"

  environment_variables = {
    NODE_ENV = "production"
  }
}