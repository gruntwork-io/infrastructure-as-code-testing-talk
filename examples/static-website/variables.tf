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
  description = "The name of the S3 bucket and all other resources created by this module."
  type        = string
  default     = "static-website-example"
}