data "terraform_remote_state" "static_website" {
  count = var.terraform_state_bucket == null ? 0 : 1

  backend = "s3"
  config = {
    bucket = var.terraform_state_bucket
    region = var.terraform_state_bucket_region
    key    = var.terraform_state_bucket_static_website_key
  }
}

locals {
  # If var.url_to_proxy is specified, proxy the URL in it. Otherwise, fetch the static website remote state and read
  # the URL to proxy from that.
  url_to_proxy = var.terraform_state_bucket == null ? var.url_to_proxy : data.terraform_remote_state.static_website[0].outputs.website_url
}