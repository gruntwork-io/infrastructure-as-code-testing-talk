terraform {
  # This module is now only being tested with Terraform 0.15.x. However, to make upgrading easier, we are setting
  # 0.12.26 as the minimum version, as that version added support for required_providers with source URLs, making it
  # forwards compatible with 0.15.x code.
  required_version = ">= 0.12.26"
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
  runtime    = "nodejs10.x"
  handler    = "index.handler"

  environment_variables = {
    NODE_ENV = "production"
  }
}