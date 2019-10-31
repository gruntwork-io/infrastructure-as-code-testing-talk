# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ---------------------------------------------------------------------------------------------------------------------

variable "name" {
  description = "The name of the S3 bucket and all other resources created by this module."
  type        = string
}