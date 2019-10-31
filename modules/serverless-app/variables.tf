# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ---------------------------------------------------------------------------------------------------------------------

variable "name" {
  description = "The name of the Lambda function and all other resources created by this module."
  type        = string
}

variable "source_dir" {
  description = "The path to the folder to zip up and upload as the Lambda deployment package."
  type        = string
}

variable "runtime" {
  description = "The runtime environment for the Lambda function (e.g. nodejs, python2.7, java8). See https://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html#SSS-CreateFunction-request-Runtime for all possible values."
  type        = string
}

variable "handler" {
  description = "The function entrypoint in your code. This is typically the name of a function or method in your code that AWS will execute when this Lambda function is triggered."
  type        = string
}

# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------

variable "timeout" {
  description = "The maximum amount of time, in seconds, your Lambda function will be allowed to run. Must be between 1 and 300 seconds."
  type        = number
  default     = 5
}

variable "memory_size" {
  description = "The maximum amount of memory, in MB, your Lambda function will be able to use at runtime. Can be set in 64MB increments from 128MB up to 1536MB. Note that the amount of CPU power given to a Lambda function is proportional to the amount of memory you request, so a Lambda function with 256MB of memory has twice as much CPU power as one with 128MB."
  type        = number
  default     = 128
}

variable "environment_variables" {
  description = "A map of environment variables to pass to the Lambda function."
  type        = map(string)

  # Lambda does not permit you to pass it an empty map of environment variables, so our default value has to contain
  # this totally useless placeholder.
  default = {
    EnvVarPlaceHolder = "Placeholder"
  }
}