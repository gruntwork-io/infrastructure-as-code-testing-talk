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