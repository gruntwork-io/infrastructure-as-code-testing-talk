data "terraform_remote_state" "web_service" {
  count = var.terraform_state_bucket == null ? 0 : 1

  backend = "s3"
  config = {
    bucket = var.terraform_state_bucket
    region = var.terraform_state_bucket_region
    key    = var.terraform_state_bucket_web_service_key
  }
}

locals {
  # If var.url_to_proxy is specified, proxy the URL in it. Otherwise, fetch the web-service remote state and read
  # the URL to proxy from that.
  url_to_proxy = var.terraform_state_bucket == null ? var.url_to_proxy : data.terraform_remote_state.web_service[0].outputs.url
}