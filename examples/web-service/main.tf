provider "aws" {
  region = var.aws_region
}

terraform {
  # This is a partial configuration for the backend. All the other settings will be provided via command-line
  # parameters. https://www.terraform.io/docs/backends/config.html#partial-configuration
  backend "s3" {}
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE A SERVERLESS "HELLO, WORLD" WEB SERVICE
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