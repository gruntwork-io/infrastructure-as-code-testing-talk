provider "aws" {
  region = var.aws_region
}

terraform {
  # This is a partial configuration for the backend. All the other settings will be provided via command-line
  # parameters. https://www.terraform.io/docs/backends/config.html#partial-configuration
  backend "s3" {}
}

# ---------------------------------------------------------------------------------------------------------------------
# CREATE THE S3 STATIC WEBSITE
# ---------------------------------------------------------------------------------------------------------------------

module "static_website" {
  source = "../../modules/s3-website"

  name = var.name
}