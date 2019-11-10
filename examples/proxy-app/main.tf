provider "aws" {
  region = var.aws_region
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE A SERVERLESS APP THAT PROXIES THE WEB-SERVICE MODULE
# ---------------------------------------------------------------------------------------------------------------------

module "proxy_app" {
  source = "../../modules/serverless-app"

  name = var.name

  source_dir = "${path.module}/javascript"
  runtime    = "nodejs10.x"
  handler    = "index.handler"

  environment_variables = {
    NODE_ENV     = "production"
    URL_TO_PROXY = local.url_to_proxy
  }
}
