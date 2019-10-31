# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------

variable "aws_region" {
  description = "The AWS region to deploy into"
  type        = string
  default     = "us-east-2"
}

variable "name" {
  description = "The name of the Lambda function, S3 bucket, and all other resources created by this module."
  type        = string
  default     = "proxy-app-example"
}

variable "terraform_state_bucket" {
  description = "The name of the S3 bucket that stores your Terraform state. Either the terraform_state_bucket_xxx variables or the url_to_proxy variable should be set to tell the app which URL to proxy."
  type        = string
  default     = null
}

variable "terraform_state_bucket_region" {
  description = "The region where the S3 bucket that stores your Terraform state lives. Either the terraform_state_bucket_xxx variables or the url_to_proxy variable should be set to tell the app which URL to proxy."
  type        = string
  default     = null
}

variable "terraform_state_bucket_static_website_key" {
  description = "The path in the S3 bucket where the Terraform state for the static-website module lives. Either the terraform_state_bucket_xxx variables or the url_to_proxy variable should be set to tell the app which URL to proxy."
  type        = string
  default     = null
}

variable "url_to_proxy" {
  description = "The URL to proxy. Either this variable or the terraform_state_bucket_xxx variables should be set to tell the app which URL to proxy."
  type        = string
  default     = "https://www.example.com"
}
